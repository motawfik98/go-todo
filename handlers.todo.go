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
