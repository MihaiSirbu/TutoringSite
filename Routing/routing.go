package Routing

import (

	"github.com/gorilla/mux"

	"log"
	"net/http"
	"html/template"
	"encoding/json"

	_ "github.com/lib/pq"

	//"gorm.io/driver/postgres"
  	//"gorm.io/gorm"
	
	"github.com/MihaiSirbu/TutoringSite/initializers"
	"github.com/MihaiSirbu/TutoringSite/models"

)

type LessonRequest struct {
    Title         string `json:"title"`
    Description   string `json:"description"`
    Creator       string `json:"creator"`
    Student       string `json:"student"`
    LessonDate    int    `json:"lessonDate"`
    LessonNumber  int    `json:"lessonNumber"`
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    // Parse the HTML template file
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        // Handle the error, such as by logging it and sending a generic
        // "Internal Server Error" message to the client
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Execute the template, writing the generated HTML to the `http.ResponseWriter`
    // The `data` parameter is used to pass dynamic data to the template
    err = t.Execute(w, data)
    if err != nil {
        // Handle the error as above
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}


func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"templates/homePage.html",nil)
}

func lessonsPage(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"templates/lessonsPage.html",nil)

}

func loginPage(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"templates/loginPage.html",nil)
}

// CRUD operations for lessons

// creates a lesson and adds it to the DB as a response to a POST request to /lessons
func CreateLesson(w http.ResponseWriter, r *http.Request){
	var req LessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
	lesson := models.Lesson{Title:req.Title, Description:req.Description, Creator:req.Creator, Student:req.Student, LessonDate:req.LessonDate, LessonNumber:req.LessonNumber}
	log.Println("Creating lesson !")
	log.Println("Title: ",req.Title)
	log.Println("Descrip", req.Description)
	result := initializers.DB.Create(&lesson)
	if result.Error != nil{
		log.Fatal("StatusInternalServerError")
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(lesson)
}


/*
func UpdateLesson(w http.ResponseWriter, r *http.Request){
	var req LessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
	lesson := models.Lesson{Title:req.Title, Description:req.Description, Creator:req.Creator, Student:req.Student, LessonDate:req.LessonDate, LessonNumber:req.LessonNumber}
	log.Println("Creating lesson !")
	log.Println("Title: ",req.Title)
	log.Println("Descrip", req.Description)
	result := initializers.DB.Create(&lesson)
	if result.Error != nil{
		log.Fatal("StatusInternalServerError")
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(lesson)
}

func DeleteLesson(db *gorm.DB){
	initializers.DB.Delete(&alesson)
}
*/

// MUST UPDATE TO RECEIVE FROM SPECIFIC TOKEN
func GetAllLessons(w http.ResponseWriter, r *http.Request){
	var lessons []models.Lesson

	// Assuming 'X' is the value for the student you are searching for.
	initializers.DB.Where(&models.Lesson{Student: "Nick"}).Find(&lessons)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(lessons)
	
}










func RunServer(port int) {
	router := mux.NewRouter()
	// get methods
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/lessons", CreateLesson).Methods("POST")
	router.HandleFunc("/lessons", GetAllLessons).Methods("GET")
	router.HandleFunc("/login", loginPage).Methods("GET")
	//router.HandleFunc("/login", loginAuth).Methods("POST")

	




	log.Println("Server starting on port:",port)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
			log.Fatal("There's an error with the server,", err)
	}

}