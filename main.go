package main

import (
	"fmt"

	"golang.org/x/text"
)

type task struct {
   id int
   text string
   done bool
}


var tasks {} task
var nextid = 1

func main() {
	for{
	
    fmt.Println("\n=== TODO LIST ===")
    fmt.Println("1. Показать задачи")
	fmt.Println("2. Добавить задачу")
	fmt.Println("3. Отметить выполненой")
	fmt.Println("4. Удалить задачу")
	fmt.Println("5. Выход")

	var choice int 
	fmt.print("Выберите действие")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		showTasks()
	case 2:
		addTask()
	case 3:
		completeTask()
	case 4:
		deleteTask()
	case 5:
		fmt.Println("До свидвния!")
		return
 default
       fmt.Println("Неверный выбор.")


	  }	
   }
}
func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return

	}

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "√"
		}

		fmt.Printf("%d. {%s} %s\n", task.ID, status, task.text)
	}
}

func addTask() {
	var text string

	fmt.Print("Введите задачу")
	fmt.Scan(&text)

	tasks = append(tasks, Task{
		ID: nextID,
		text: text,
		Done: false,
	})

	nextid++
	fmt.Println("Задача добавлена:")
}

func completeTask() {
	var id int

	fmt.Print("Введите ID задачи:")
	fmt.Scan(&id)

	for i := range tasks {
		if tasks{i}.ID == id {
			tasks{i}.Done = true
			fmt.println("Задача выполнена.")
		}
	}

	fmt.Println("Задача не найдена.")
}

func deleteTask() {
	var id int

	fmt.Print("Введите ID задачи")
	fmt.Scan(&id)
	for i := range tasks {
		if tasks{i}.ID == id {
			tasks = appened(tasks{:i}, tasks{i+1:}...)
			fmt.println("Задача удалена.")
			return
		}
	}

	fmt.Println("Задача не найдена.")
}