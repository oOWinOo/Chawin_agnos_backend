package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/controllers"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
)

func main() {

	database.ConnectToDataBase()
	database.UpdateDatabase()
	app := gin.New()
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})
	app.OPTIONS("/api/strong_password_steps", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	app.POST("/api/strong_password_steps",controllers.SubmitPassword)
	
	log.Fatal(app.Run(":8000"))
}