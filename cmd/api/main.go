/*
export DB_DRIVER=postgres
export DB_NAME=postgres
export DB_PASS=password
export DB_PORT=5432
export DB_USER=postgres
export DB_HOST=0.0.0.0
go run main.go
*/
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

	srv.Mux.HandleFunc("/healthz", srv.Healthz)
	srv.Mux.HandleFunc("/ready", srv.Ready)
	//srv.Mux.Handle("/randoms", app.RateLimiter(srv.Randoms))
	srv.Mux.HandleFunc("/randoms", srv.Randoms)

	addr := fmt.Sprintf("%s:%d", srv.Api.Host, srv.Api.Port)
	log.Printf("running api at %v\n", addr)
	err := http.ListenAndServe(addr, srv.Mux)
	if err != nil {
		return
	}

	defer srv.DB.Close()
}
