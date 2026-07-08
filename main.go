package main

import (
 "fmt"
 "net/http"
)

func main() {

 db, err := ConnectDB()

 if err != nil {
  fmt.Println("Ошибка подключения:", err)
  return
 }

 defer db.Close()


 http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

  if r.Method == "GET" {
   GetTasks(db)(w, r)
   return
  }


  if r.Method == "POST" {
   CreateTask(db)(w, r)
   return
  }


  http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)

 })


 fmt.Println("Сервер запущен: http://localhost:8080")


 err = http.ListenAndServe(":8080", nil)

 if err != nil {
  fmt.Println("Ошибка сервера:", err)
 }
}