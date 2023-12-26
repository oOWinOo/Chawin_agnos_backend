package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/controllers"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
)

func main() {
	database.ConnectToDataBase("postgres")
	database.UpdateDatabase()
	app := gin.New()
	app.Use(CORSMiddleware())
	app.POST("/api/strong_password_steps", controllers.SubmitPassword)

	log.Fatal(app.Run(":8000"))
}

func CORSMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
		}
		return
	}
}
