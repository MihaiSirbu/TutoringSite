package Routing

import (

	"github.com/gorilla/mux"
	"net/http"
	//"encoding/json"

	_ "github.com/lib/pq"

	//"gorm.io/driver/postgres"
  	//"gorm.io/gorm"
	
	"github.com/MihaiSirbu/TutoringSite/initializers"
	"github.com/MihaiSirbu/TutoringSite/models"

)


// Go handler function that responds to a GET request for a specific lesson's details
func GetLesson(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
    id := vars["id"]

	var lesson models.Lesson
	
	// fetching single lesson based on id
	initializers.DB.First(&lesson, id)

	// need to add error finding
	initializers.DB.Where("LessonNumber = ?", lesson.LessonNumber).Find(&lesson.Exercises)

	

	renderTemplate(w,"templates/loginPage.html",lesson)
	
}


