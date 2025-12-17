package main

import (
	"context"
	"log/slog"
	"os"

	"go_playground/internal/env"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecommerce sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		os.Exit(1)
		// panic(err)
	}

	defer conn.Close(ctx)

	slog.Info("connected to database")

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("error running server", "error", err)
		os.Exit(1)
	}
}
