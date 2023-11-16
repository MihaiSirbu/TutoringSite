package main

import (

	"github.com/MihaiSirbu/TutoringSite/initializers"	
	"github.com/MihaiSirbu/TutoringSite/models"	
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Lesson{})
	initializers.DB.AutoMigrate(&models.User{})
	
}