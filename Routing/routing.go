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
	"fmt"
	
	"github.com/MihaiSirbu/TutoringSite/initializers"
	"github.com/MihaiSirbu/TutoringSite/models"
	"github.com/MihaiSirbu/TutoringSite/authentication"
	"github.com/MihaiSirbu/TutoringSite/user"

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

    
    err = t.Execute(w, data)
    if err != nil {
        // Handle the error as above
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}


func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"templates/homePage.html",nil)
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"templates/registerPage.html",nil)
}

func lessonsPage(w http.ResponseWriter, r *http.Request){
	claims, ok := r.Context().Value("claims").(*auth.Claims)
    if !ok {
        // Handle the error appropriately
        http.Error(w, "Error getting claims", http.StatusInternalServerError)
        return
    }
	


	
	allLessons := GetAllLessons(claims.Username)
	data := struct {
        Lessons []models.Lesson
    }{
        Lessons: allLessons,
    }

	renderTemplate(w,"templates/lessonsPage.html",data)


}




func loginPage(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"templates/loginPage.html",nil)
}

// CRUD operations for lessons

// creates a lesson and adds it to the DB as a response to a POST request to /lessons
func CreateLesson(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating lesson")
	var req LessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
	lesson := models.Lesson{Title:req.Title, Description:req.Description, Creator:req.Creator, Student:req.Student, LessonDate:req.LessonDate, LessonNumber:req.LessonNumber}
	result := initializers.DB.Create(&lesson)
	if result.Error != nil{
		log.Fatal("StatusInternalServerError")
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(lesson)
}



func UpdateLesson(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]

	var lesson models.Lesson
	
	// fetching single lesson based on id
	initializers.DB.First(&lesson, id)

	var req LessonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
	initializers.DB.Model(&lesson).Updates(models.Lesson{Title:req.Title, Description:req.Description, Creator:req.Creator, Student:req.Student, LessonDate:req.LessonDate, LessonNumber:req.LessonNumber})



	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(lesson)
}

func CreateExercise(w http.ResponseWriter, r *http.Request){
	var req models.Exercise
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

	newExercise := models.Exercise{ExerciseNumber:req.ExerciseNumber, LessonNumber:req.LessonNumber, ExerciseContent:req.ExerciseContent, Answer:req.Answer}
	result := initializers.DB.Create(&newExercise)
	if result.Error != nil{
		log.Fatal("StatusInternalServerError")
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newExercise)


    
}

func DeleteLesson(w http.ResponseWriter, r *http.Request){
	// probably needs some checks as well as error responses
	vars := mux.Vars(r)
    id := vars["id"]
	initializers.DB.Delete(&models.Lesson{}, id)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
}





// MUST UPDATE TO RECEIVE FROM SPECIFIC TOKEN
func GetAllLessons(studentName string)([]models.Lesson){
	var lessons []models.Lesson

	// Assuming 'X' is the value for the student you are searching for.
	initializers.DB.Where(&models.Lesson{Student: studentName}).Find(&lessons)

	// need to add errNotFound404

	return lessons
	
}










func RunServer(port int) {

	
 


	// routing
	router := mux.NewRouter()
	//homepage
	router.HandleFunc("/", homePage).Methods("GET")

	// lessons


	router.Handle("/lessons", auth.TokenVerifyMiddleware(http.HandlerFunc(lessonsPage))).Methods("GET")


	router.HandleFunc("/lessons", CreateLesson).Methods("POST")

	router.HandleFunc("/lessons/{id}", UpdateLesson).Methods("PUT")
	//router.HandleFunc("/lessons/{id}", DisplayLesson).Methods("GET")
	//router.HandleFunc("/lessons/{id}", GetLesson).Methods("GET")
	router.HandleFunc("/lessons/{id}", DeleteLesson).Methods("DELETE")

	router.HandleFunc("/exercises",CreateExercise).Methods("POST")


	//logins
	router.HandleFunc("/login", loginPage).Methods("GET")
	router.HandleFunc("/login", auth.LoginAuth).Methods("POST")

	//registerUser
	router.HandleFunc("/register", registrationPage).Methods("GET")
	router.HandleFunc("/register", user.RegisterUser).Methods("POST")

	//serving static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	

	



	log.Println("Server starting on port:",port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
			log.Fatal("There's an error with the server,", err)
	}

}