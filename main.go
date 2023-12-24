package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/controllers"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
)

func main() {

	database.ConnectToDataBase()
	database.UpdateDatabase()
	app := gin.New()

	app.POST("/",controllers.SubmitPassword)
	
	log.Fatal(app.Run(":1234"))
}