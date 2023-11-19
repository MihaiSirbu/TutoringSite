package user

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/MihaiSirbu/TutoringSite/initializers"
    "github.com/MihaiSirbu/TutoringSite/models"
	"net/mail"
	"unicode"
    "net/http"
    "encoding/json"
    "fmt"
    "time"
    "github.com/MihaiSirbu/TutoringSite/authentication"
	

)
// ...

type User struct {
    Username         string `json:"username"`
    Email           string `json:"email"`
    Password        string `json:"password"`
}



func VerificationsUserInputRegistration(username string,email string,password string)(bool,string){
	if !verifyUsernameStrength(username) {
        return false, "Please choose another username(Minimum 5 letters)"
    }
    if !verifyEmailDetails(email) {
        return false, "invalid email format"
    }
    if !verifyPasswordStrength(password) {
        return false, "Password is not strong enough(must contain at least 5 characters and 1 number and 1 special character)"
    }
    return true, ""
}

func verifyUsernameStrength(username string)(bool){
	if len(username) < 5{
		return false
	}

	var exists_User models.User
	initializers.DB.Where(&models.User{Username: username}).Find(&exists_User)
    if exists_User.Username != ""{
        return false
    }
	
	return true

}

func verifyEmailDetails(email string)(bool){

	_, err := mail.ParseAddress(email)
	return err == nil
}

func verifyPasswordStrength(password string)(bool){
	// password requirements : 
	//	longer than 6 chars
	//	at least 1 number
	//	at least 1 special character
	if len(password) < 7{
		return false
	}

	numberFlag,specialChar:= false,false

	for _, char := range password {
        switch {
        case unicode.IsNumber(char):
            numberFlag = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            specialChar = true
        }
        if numberFlag && specialChar {
            return true
        }
    }
	return false
		
}



func RegisterUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("We are trying to register a user")
    var req User
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    isVerified, verifErr := VerificationsUserInputRegistration(req.Username, req.Email, req.Password)
    if !isVerified {
        http.Error(w, verifErr, http.StatusBadRequest) // send back error
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error generating password hash", http.StatusInternalServerError)
        return
    }

    // Passed all the checks and inserting into our database a new user
    newUser := models.User{
        Username: req.Username,
        Email: req.Email,
        PasswordHash: string(hashedPassword),
    }
    initializers.DB.Create(&newUser)


    tokenString, err := auth.GenerateJWT(newUser.Username)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    http.SetCookie(w, &http.Cookie{
        Name: "token",
        Value: tokenString,
        Expires: time.Now().Add(24 * time.Hour), 
        
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

