package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/gmhafiz/k8s-api"
)

func main() {
	ctx := context.Background()
	srv := app.New(ctx)

	srv.Mux.Handle("/healthz", recovery(cors(http.HandlerFunc(srv.Healthz))))
	srv.Mux.Handle("/ready", recovery(cors(http.HandlerFunc(srv.Ready))))
	srv.Mux.Handle("/randoms", recovery(cors(http.HandlerFunc(srv.Randoms))))

	addr := fmt.Sprintf("%s:%d", srv.Api.Host, srv.Api.Port)
	log.Printf("running api at %v\n", addr)

	server := &http.Server{Addr: addr, Handler: srv.Mux}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error while shutting down: %v\n", err)
	}

	srv.DB.Close()
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if (*r).Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func recovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				defer r.Body.Close()
				log.Printf("PANIC: %v", rvr)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
