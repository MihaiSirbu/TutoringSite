package main

import (

	"github.com/MihaiSirbu/TutoringSite/initializers"	
	"github.com/MihaiSirbu/TutoringSite/models"	
)

func init() {
	initializers.ConnectToDB()
}

func dropTable(model interface{}){
	if initializers.DB.Migrator().HasTable(&model) {
		initializers.DB.Migrator().DropTable(&model)
	}
}

func main() {
	
	//dropTable(models.Lesson{})
	//dropTable(models.User{})
	
	
	initializers.DB.AutoMigrate(&models.Lesson{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Exercise{})
	
}