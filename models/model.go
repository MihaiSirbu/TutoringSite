package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username 	 string
	Email		 string
	PasswordHash string `gorm:"size:255"` //need at least 60 to store the bcrypt hashed password

}


type Lesson struct {
	gorm.Model
	Title			string
	Description		string
	Creator         string
	Student         string
	LessonDate      int
	LessonNumber	int
	Exercises		[]Exercise `gorm:"foreignKey:LessonID"`
	
}

type Exercise struct {
	gorm.Model
	LessonID    	uint	
	ExerciseContent string
	Answer          string
}
