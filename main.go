package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"./controllers"
	"./db"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	redisUri := os.Getenv("REDIS_URI")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	connectionString := os.Getenv("CONNECTION_STRING")
	port := os.Getenv("PORT")

	store, _ := sessions.NewRedisStore(10, "tcp", redisUri, "", []byte(redisPassword))
	r.Use(sessions.Sessions("app-session", store))
	r.Use(CORSMiddleware())
	r.Use(db.Database(connectionString))

	v1 := r.Group("/v1")
	{
		v1.POST("/todo", controllers.AddTodo)
		v1.GET("/todos", controllers.ListTodo)
		v1.DELETE("/todo/:id", controllers.DeleteTodo)
	}
	r.Run(":" + string(port))
}
