package main

import (

	"github.com/MihaiSirbu/TutoringSite/Routing"
	"github.com/MihaiSirbu/TutoringSite/initializers"


    
	

)



func init() {
	initializers.ConnectToDB()
}

func main() {
	Routing.RunServer(8080)
}
