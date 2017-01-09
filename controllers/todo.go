package controllers

import (
    "../models"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

func AddTodo(c *gin.Context) {

    db := c.MustGet("DB").(*gorm.DB)
    
    var todo models.Todo
    if c.BindJSON(&todo) != nil {
        c.JSON(406, gin.H{"message": "Invalid data", "data": todo})
        c.Abort()
        return
    }

    if !db.NewRecord(todo) {
        c.JSON(406, gin.H{"message": "Todo could not be created"})
        c.Abort()
        return
        
    }
    db.Create(&todo)
    c.JSON(200, gin.H{"message": "Todo created"})
}

func DeleteTodo(c *gin.Context) {
    db := c.MustGet("DB").(*gorm.DB)
    id := c.Param("id")
    var todo models.Todo
    db.First(&todo, id)
    db.Delete(&todo)
    c.JSON(200, gin.H{"message": "ok"})
}

func ListTodo(c *gin.Context) {
    db := c.MustGet("DB").(*gorm.DB)
    todos := []models.Todo{}
    db.Find(&todos)
    c.JSON(200, gin.H{"todos": &todos})
}