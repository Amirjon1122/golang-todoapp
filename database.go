package main

import (
 "context"

 "github.com/jackc/pgx/v5/pgxpool"
)


func ConnectDB() (*pgxpool.Pool, error) {

 db, err := pgxpool.New(
  context.Background(),
  "postgres://postgres:postgres12345@localhost:5432/todo?sslmode=disable",
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