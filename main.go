package main

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"html/template"
)

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



func RunServer(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/lessons", lessonsPage).Methods("GET")
	router.HandleFunc("/login", loginPage).Methods("GET")


	log.Println("Server starting on port:",port)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
			log.Fatal("There's an error with the server,", err)
	}

}

func main() {
	RunServer(8080)
}
