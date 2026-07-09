package main

import (
 "context"
 "encoding/json"
 "net/http"
 "strconv"

 "github.com/jackc/pgx/v5/pgxpool"
)


// GET /tasks
func GetTasks(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {

  rows, err := db.Query(
   context.Background(),
   "SELECT id, title, description, completed FROM tasks ORDER BY id",
  )

  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }

  defer rows.Close()

  tasks := []Task{}

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

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(tasks)
 }
}


// POST /tasks
func CreateTask(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {

  var task Task

  err := json.NewDecoder(r.Body).Decode(&task)

  if err != nil {
   http.Error(w, err.Error(), http.StatusBadRequest)
   return
  }


  err = db.QueryRow(
   context.Background(),
   `INSERT INTO tasks(title, description, completed)
    VALUES($1,$2,$3)
    RETURNING id`,
   task.Title,
   task.Description,
   task.Completed,
  ).Scan(&task.ID)


  if err != nil {
   http.Error(w, err.Error(), 500)
   return
  }


  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(task)
 }
}


// PUT /tasks?id=1
func UpdateTask(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {

  id := r.URL.Query().Get("id")

  taskID, err := strconv.Atoi(id)

  if err != nil {
   http.Error(w, "invalid id", http.StatusBadRequest)
   return
  }


  var task Task

  err = json.NewDecoder(r.Body).Decode(&task)

  if err != nil {
   http.Error(w, err.Error(), http.StatusBadRequest)
   return
  }


  result, err := db.Exec(
   context.Background(),
   `UPDATE tasks
    SET title=$1,
        description=$2,
        completed=$3
    WHERE id=$4`,
   task.Title,
   task.Description,
   task.Completed,
   taskID,
  )


  if err != nil {
   http.Error(w, err.Error(), 500)
   return
  }


  rows := result.RowsAffected()

  if rows == 0 {
   http.Error(w, "task not found", http.StatusNotFound)
   return
  }


  w.WriteHeader(http.StatusNoContent)
 }
}


// DELETE /tasks?id=1
func DeleteTask(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {

  id := r.URL.Query().Get("id")

  taskID, err := strconv.Atoi(id)

  if err != nil {
   http.Error(w, "invalid id", http.StatusBadRequest)
   return
  }


  result, err := db.Exec(
   context.Background(),
   "DELETE FROM tasks WHERE id=$1",
   taskID,
  )


  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }


  rows := result.RowsAffected()

  if rows == 0 {
   http.Error(w, "task not found", http.StatusNotFound)
   return
  }


  w.WriteHeader(http.StatusNoContent)
 }
}