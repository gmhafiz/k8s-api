package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	app "github.com/gmhafiz/k8s-api"
)

func main() {
	ctx := context.Background()
	srv := app.New(ctx)

	srv.Mux.Handle("/healthz", cors(http.HandlerFunc(srv.Healthz)))
	srv.Mux.Handle("/ready", cors(http.HandlerFunc(srv.Ready)))
	srv.Mux.Handle("/randoms", cors(http.HandlerFunc(srv.Randoms)))

	addr := fmt.Sprintf("%s:%d", srv.Api.Host, srv.Api.Port)
	log.Printf("running api at %v\n", addr)
	err := http.ListenAndServe(addr, srv.Mux)
	if err != nil {
		return
	}

	defer srv.DB.Close()
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
