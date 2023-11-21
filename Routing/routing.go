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
	"bytes"

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
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Use a buffer to execute the template
    buf := new(bytes.Buffer)
    err = t.Execute(buf, data)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Write the buffer to the response writer
    w.Write(buf.Bytes())
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
		fmt.Println("Error in getting claims in Lessons page")
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
	

	fmt.Println("sending data back to user")
	renderTemplate(w,"templates/lessonsPage.html",data)
	fmt.Println("data has been sent successfully")
	


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

	newExercise := models.Exercise{ExerciseNumber:req.ExerciseNumber,LessonID:req.LessonID, ExerciseContent:req.ExerciseContent, Answer:req.Answer}
	result := initializers.DB.Create(&newExercise)
	if result.Error != nil{
		log.Fatal("StatusInternalServerError")
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newExercise)

    
}

func UpdateExercise(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]

	var exercise models.Exercise
	
	// fetching single lesson based on id
	initializers.DB.First(&exercise, id)

	var req models.Exercise
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
	initializers.DB.Model(&exercise).Where("status != ?", "Completed").Updates(models.Exercise{ExerciseNumber:req.ExerciseNumber, LessonID:req.LessonID, ExerciseContent:req.ExerciseContent, Answer:req.Answer, Status:req.Status})



	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(exercise)
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

	// return all the lessons where the student name is passed in
	initializers.DB.Where("Student = ?", studentName).Find(&lessons)

	// need to add errNotFound404

	return lessons
	
}

// get specific lesson w/ exercises
func GetLesson(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
    id := vars["id"]

	var lesson models.Lesson
	
	result := initializers.DB.Where("id = ?", id).Preload("Exercises").Find(&lesson)

	fmt.Println("our specific lesson title: ", lesson.Title)
	fmt.Println("our specific lesson id: ", lesson.ID)

	if result.Error != nil {
        fmt.Println("Error fetching lessons with exercises:", result.Error)
        return
    }

	renderTemplate(w,"templates/singleLessonPage.html",lesson)
	
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
	router.HandleFunc("/lessons/{id:[0-9]+}", GetLesson).Methods("GET")
	router.HandleFunc("/lessons/{id}", DeleteLesson).Methods("DELETE")

	router.HandleFunc("/exercises",CreateExercise).Methods("POST")
	router.HandleFunc("/exercises/{id:[0-9]+}",UpdateExercise).Methods("PUT")


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