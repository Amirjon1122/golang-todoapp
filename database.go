package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(
		context.Background(),
		dbURL,
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	return db, nil
}
