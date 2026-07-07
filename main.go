package main

import (
 "fmt"
)

func main() {

 db, err := ConnectDB()

 if err != nil {
  fmt.Println("Ошибка подключения:", err)
  return
 }

 defer db.Close()

 fmt.Println("Программа работает")
}