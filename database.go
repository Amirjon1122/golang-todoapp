package main

import (
 "context"
 "fmt"

 "github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {

 db, err := pgxpool.New(
  context.Background(),
  "postgres://postgres:postrest@localhost:5432/todo",
 )

 if err != nil {
  return nil, err
 }

 err = db.Ping(context.Background())

 if err != nil {
  return nil, err
 }

 fmt.Println("База данных подключена!")

 return db, nil
}