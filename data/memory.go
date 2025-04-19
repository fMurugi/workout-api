package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"workout-api/models"
)

var DB *gorm.DB
func InitDB(){
	var err error
	DB, err = gorm.Open(sqlite.Open("workouts.db"),&gorm.Config{})
	if err != nil{
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&models.Workout{})
	if err != nil{
		log.Fatal("failed to migrate database:", err)
	}
}

var Workouts = make(map[string]models.Workout)
