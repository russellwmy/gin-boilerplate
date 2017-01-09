package db

import (
    "../models"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database(connString string) gin.HandlerFunc {
    db, err := gorm.Open("mysql", connString)
    db.LogMode(true)
    // Error
    if err != nil {
        panic(err)
    }
    
    db.AutoMigrate(&models.Todo{})

    return func(c *gin.Context) {
        c.Set("DB", db)
        c.Next()
    }
}