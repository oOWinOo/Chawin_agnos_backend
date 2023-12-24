package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oOWinOo/Chawin_agnos_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
const(
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)


func ConnectToDataBase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,			//Time manage
			LogLevel: logger.Info,				//4 types Silent Error Warn Info  [notification]
			Colorful: true,						//color
		},
	)
	var err error
	Db,err = gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: newLogger,
	})
	if err != nil{
		panic("Failed to connect to database")
	}
	fmt.Println("Connect to database Complete")
}

func UpdateDatabase(){
	Db.AutoMigrate(&models.LogInOut{})
}