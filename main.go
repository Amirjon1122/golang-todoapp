package main

import (
	"log"
	"net/http"
)

func main() {

	// подключаем базу данных
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// маршрут /tasks
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		// получить все задачи
		case http.MethodGet:
			GetTasks(db)(w, r)

		// создать задачу
		case http.MethodPost:
			CreateTask(db)(w, r)

		// обновить задачу
		case http.MethodPut:
			UpdateTask(db)(w, r)

		// удалить задачу
		case http.MethodDelete:
			DeleteTask(db)(w, r)

		default:
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})

	log.Println("Server started on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
