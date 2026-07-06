package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type Task struct {
 ID   int
 Text string
 Done bool
}

var tasks []Task
var nextID = 1

func main() {
 reader := bufio.NewReader(os.Stdin)

 for {
  fmt.Println("\n===== TODO LIST =====")
  fmt.Println("1. Показать задачи")
  fmt.Println("2. Добавить задачу")
  fmt.Println("3. Выполнить задачу")
  fmt.Println("4. Удалить задачу")
  fmt.Println("5. Выход")
  fmt.Print("Выберите пункт: ")

  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  switch input {
  case "1":
   showTasks()
  case "2":
   addTask(reader)
  case "3":
   completeTask(reader)
  case "4":
   deleteTask(reader)
  case "5":
   fmt.Println("До свидания!")
   return
  default:
   fmt.Println("Неверный выбор.")
  }
 }
}

func showTasks() {
 if len(tasks) == 0 {
  fmt.Println("Список задач пуст.")
  return
 }

 fmt.Println("\nВаши задачи:")

 for _, task := range tasks {
  status := " "
  if task.Done {
   status = "✓"
  }

  fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
 }
}

func addTask(reader *bufio.Reader) {
 fmt.Print("Введите задачу: ")

 text, _ := reader.ReadString('\n')
 text = strings.TrimSpace(text)

 tasks = append(tasks, Task{
  ID:   nextID,
  Text: text,
  Done: false,
 })

 nextID++

 fmt.Println("Задача добавлена.")
}

func completeTask(reader *bufio.Reader) {
 fmt.Print("Введите ID задачи: ")

 input, _ := reader.ReadString('\n')
 input = strings.TrimSpace(input)

 id, err := strconv.Atoi(input)
 if err != nil {
  fmt.Println("Некорректный ID.")
  return
 }

 for i := range tasks {
  if tasks[i].ID == id {
   tasks[i].Done = true
   fmt.Println("Задача выполнена.")
   return
  }
 }

 fmt.Println("Задача не найдена.")
}

func deleteTask(reader *bufio.Reader) {
 fmt.Print("Введите ID задачи: ")

 input, _ := reader.ReadString('\n')
 input = strings.TrimSpace(input)

 id, err := strconv.Atoi(input)
 if err != nil {
  fmt.Println("Некорректный ID.")
  return
 }

 for i := range tasks {
  if tasks[i].ID == id {
   tasks = append(tasks[:i], tasks[i+1:]...)
   fmt.Println("Задача удалена.")
   return
  }
 }

 fmt.Println("Задача не найдена.")
}