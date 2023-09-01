package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Task struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
    Done    bool   `json:"done"`
}

var tasks []Task
var currentID = 1

func GetTasks(c echo.Context) error {
    return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
    task := new(Task)
    if err := c.Bind(task); err != nil {
        return err
    }

    task.ID = currentID
    currentID++
    tasks = append(tasks, *task)

    return c.JSON(http.StatusCreated, task)
}

func main() {
    e := echo.New()

    // Dummy data tugas
    tasks = append(tasks, Task{ID: 1, Content: "Belajar Go", Done: false})
    tasks = append(tasks, Task{ID: 2, Content: "Buat API", Done: true})

    // Route untuk mendapatkan daftar tugas
    e.GET("/tasks", GetTasks)

    // Route untuk membuat tugas baru
    e.POST("/tasks", CreateTask)

    e.Start(":8080")
}
