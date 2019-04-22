package main

import (
  "strconv"
  "github.com/gin-gonic/gin"
  "net/http"
)

func createTodo(c *gin.Context) {
  completed, _ := strconv.Atoi(c.PostForm("completed"))
  todo := todoModel{Title: c.PostForm("title"), Completed: completed}
  db.Save(&todo)
  c.JSON(http.StatusCreated,
    gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID},
  )
}

func fetchAllTodo(c *gin.Context) {
  var todos []todoModel
  var _todos []todoModel

  db.Find(&todos)

  if len(todos) <= 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }
  for _, item := range todos {
    _todos = append(_todos, todoModel{Title: item.Title, Completed: item.Completed})
 }
 c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func fetchSingleTodo(c *gin.Context) {
  var todo todoModel
  todoID := c.Param("id")
  db.First(&todo, todoID)
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func updateTodo(c *gin.Context) {
  var todo todoModel
  todoID := c.Param("id")

  db.First(&todo, todoID)
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }

  db.Model(&todo).Update("title", c.PostForm("title"))
  completed, _ := strconv.Atoi(c.PostForm("completed"))
  db.Model(&todo).Update("completed", completed)
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}
