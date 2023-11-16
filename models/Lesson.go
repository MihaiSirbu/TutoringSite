package models

import (
	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Title			string
	Description		string
	Creator         string
	Student         string
	LessonDate      int
	LessonNumber	int
	
}

type User struct {
	gorm.Model
	Username string
	PasswordHash string

}