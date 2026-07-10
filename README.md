# Todo List API

REST API для управления задачами (Todo List), написанный на Go с использованием PostgreSQL.

## Технологии

- Go
- PostgreSQL
- pgx
- net/http

---

# Установка и запуск

## 1. Клонировать проект
git clone https://github.com/Amirjon1122/golang-todoapp.git

Перейти в папку проекта:
cd golang-todoapp

---

## 2. Создать базу данных PostgreSQL

Создать базу:
CREATE DATABASE todo;

Подключиться к базе и выполнить файл:
\i schema.sql

Файл schema.sql создаст таблицу:
tasks

---

## 3. Настроить подключение к базе

Открыть файл:
database.go

и указать свои данные PostgreSQL:
postgres://username:password@localhost:5432/todo

---

## 4. Запустить проект

В терминале выполнить:
go run .

Сервер запустится:
http://localhost:8080

---

# API запросы

## Получить все задачи

GET:
http://localhost:8080/tasks

---

## Получить задачу по ID

GET:
http://localhost:8080/tasks?id=1

---

## Создать задачу

POST:
http://localhost:8080/tasks

Body:
{
    "title": "Изучить Go",
    "description": "Создать Todo List",
    "completed": false
}

---

## Обновить задачу

PUT:
http://localhost:8080/tasks?id=1

Body:
{
    "title": "Изучить PostgreSQL",
    "description": "Добавить изменения",
    "completed": true
}

---

## Удалить задачу

DELETE:
http://localhost:8080/tasks?id=1

---

# Структура проекта
golang-todoapp
│
├── main.go
├── database.go
├── handlers.go
├── task.go
├── schema.sql
├── go.mod
└── go.sum