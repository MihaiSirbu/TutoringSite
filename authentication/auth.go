package auth
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"html/template"
	"github.com/dgrijalva/jwt-go"
    "time"
    
    
)


type UserCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func loginAuth(w http.ResponseWriter, r *http.Request) {
    var creds UserCredentials
    // Decode the JSON body into the `creds` struct.
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        // If there is something wrong with the request body, return a 400 status
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // TODO: Get the user's hashed password from the database.
    expectedPasswordHash, err := GetUserPasswordHash(creds.Username)

    // TODO: Compare the stored hashed password, with the hashed version of the password that was received
    if err != nil || !CheckPasswordHash(creds.Password, expectedPasswordHash) {
        // If the passwords don't match or there's an error, return a 401 status
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // If the login is successful, generate a new JWT token for the user
    tokenString, err := GenerateJWT(creds.Username)
    if err != nil {
        // If there is an error in creating the JWT return an internal server error
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Finally, set the client JWT token as a cookie and send a response
    http.SetCookie(w, &http.Cookie{
        Name: "token",
        Value: tokenString,
        Expires: time.Now().Add(24 * time.Hour), // or whatever expiry you want
        HttpOnly: true, // ensures the cookie is sent only over HTTP(S), not accessible by JavaScript
    })

    // You could also send the token in the response body, depending on your frontend's requirements
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func GenerateJWT(username string)(string){
	return 
}

func CheckPasswordHash(creds.Password, expectedPasswordHash string)(bool){
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GetUserPasswordHash(username string)(string){

}