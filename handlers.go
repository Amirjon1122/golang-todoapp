package main

import (
 "context"
 "encoding/json"
 "fmt"
 "net/http"

 "github.com/jackc/pgx/v5/pgxpool"
)


// Получить все задачи
func GetTasks(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {

  rows, err := db.Query(
   context.Background(),
   "SELECT id, title, description, completed FROM tasks",
  )

  if err != nil {
   http.Error(w, err.Error(), 500)
   return
  }

  defer rows.Close()


  var tasks []Task


  for rows.Next() {

   var task Task


   err := rows.Scan(
    &task.ID,
    &task.Title,
    &task.Description,
    &task.Completed,
   )


   if err != nil {
    http.Error(w, err.Error(), 500)
    return
   }


   tasks = append(tasks, task)
  }


  w.Header().Set(
   "Content-Type",
   "application/json",
  )


  json.NewEncoder(w).Encode(tasks)
 }
}



// Создать задачу через JSON
func CreateTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  var task Task


  err := json.NewDecoder(r.Body).Decode(&task)


  if err != nil {

   http.Error(w, err.Error(), 400)

   return
  }



  _, err = db.Exec(
   context.Background(),
   "INSERT INTO tasks(title, description) VALUES($1, $2)",
   task.Title,
   task.Description,
  )



  if err != nil {

   http.Error(w, err.Error(), 500)

   return
  }



  w.Header().Set(
   "Content-Type",
   "application/json",
  )


  json.NewEncoder(w).Encode(task)

 }
}



// Изменить задачу
func UpdateTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  id := r.FormValue("id")


  var task Task


  err := json.NewDecoder(r.Body).Decode(&task)


  if err != nil {

   http.Error(w, err.Error(), 400)

   return
  }



  _, err = db.Exec(
   context.Background(),
   "UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4",
   task.Title,
   task.Description,
   task.Completed,
   id,
  )



  if err != nil {

   http.Error(w, err.Error(), 500)

   return
  }



  w.Header().Set(
   "Content-Type",
   "application/json",
  )


  json.NewEncoder(w).Encode(task)

 }
}



// Удалить задачу
func DeleteTask(db *pgxpool.Pool) http.HandlerFunc {

 return func(w http.ResponseWriter, r *http.Request) {


  id := r.FormValue("id")


  _, err := db.Exec(
   context.Background(),
   "DELETE FROM tasks WHERE id=$1",
   id,
  )



  if err != nil {

   http.Error(w, err.Error(), 500)

   return
  }



  fmt.Fprintln(w, "Задача удалена")

 }
}