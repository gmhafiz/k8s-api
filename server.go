package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Cfg

	Mux *http.ServeMux
	DB  *pgxpool.Pool
}

func New(ctx context.Context) *Server {
	cfg := Config()

	dsn := fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.Database.User,
		cfg.Database.Pass,
	)

	pool, err := pgxpool.New(ctx, dsn)

	//conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = pool.Exec(ctx, "SELECT true")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	return &Server{
		Cfg: cfg,
		DB:  pool,
		Mux: mux,
	}
}
