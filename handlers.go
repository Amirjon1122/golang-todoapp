package main

import (
 "context"
 "fmt"
 "net/http"

 "github.com/jackc/pgx/v5/pgxpool"
)


// Получение всех задач
func GetTasks(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {

  rows, err := db.Query(
   context.Background(),
   "SELECT id, title, description, completed FROM tasks",
  )

  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }

  defer rows.Close()


  for rows.Next() {

   var id int
   var title string
   var description string
   var completed bool


   err := rows.Scan(
    &id,
    &title,
    &description,
    &completed,
   )

   if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
   }


   fmt.Fprintf(
    w,
    "%d | %s | %s | %t<br>",
    id,
    title,
    description,
    completed,
   )
  }
 }
}



// Создание задачи
func CreateTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  title := r.FormValue("title")
  description := r.FormValue("description")


  _, err := db.Exec(
   context.Background(),
   "INSERT INTO tasks(title, description) VALUES($1, $2)",
   title,
   description,
  )


  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }


  fmt.Fprintln(w, "Задача создана")
 }
}



// Изменение задачи
func UpdateTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  id := r.FormValue("id")
  title := r.FormValue("title")
  description := r.FormValue("description")
  completed := r.FormValue("completed")


  _, err := db.Exec(
   context.Background(),
   "UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4",
   title,
   description,
   completed,
   id,
  )


  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }


  fmt.Fprintln(w, "Задача обновлена")
 }
}



// Удаление задачи
func DeleteTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  id := r.FormValue("id")


  _, err := db.Exec(
   context.Background(),
   "DELETE FROM tasks WHERE id=$1",
   id,
  )


  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }


  fmt.Fprintln(w, "Задача удалена")
 }
}