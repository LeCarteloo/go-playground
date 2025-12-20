package main

import (
	"context"
	"log/slog"
	"os"

	"go_playground/internal/env"

	"github.com/jackc/pgx/v5"
)

func setupLogger() {
	var handler slog.Handler

	if env.GetString("ENV", "development") == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	slog.SetDefault(slog.New(handler))
}

func main() {
	setupLogger()

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecommerce sslmode=disable"),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		os.Exit(1)
		// panic(err)
	}

	defer conn.Close(ctx)

	// TODO: Add db name + port
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
