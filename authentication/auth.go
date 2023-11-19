package auth

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
    "time"
    "github.com/MihaiSirbu/TutoringSite/initializers"
    "github.com/MihaiSirbu/TutoringSite/models"
    "fmt"
    "gorm.io/gorm"
    "errors"
    


)

var jwtKey = []byte("Mihai04")


type UserCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
    fmt.Println("We entered login POST rsq")
    
    var creds UserCredentials
    

	
    // Decode the JSON body into the `creds` struct.
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        // If there is something wrong with the request body, return a 400 status
        w.WriteHeader(http.StatusBadRequest)
        return
    }


    // TODO: Get the user's hashed password from the database.
    expectedPasswordHash,err := GetUserPasswordHash(creds.Username)

    // TODO: Compare the stored hashed password, with the hashed version of the password that was received
    if err != nil || !CheckPasswordHash(expectedPasswordHash, creds.Password) {
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
        Expires: time.Now().Add(24 * time.Hour), 
        //HttpOnly: true, // ensures the cookie is sent only over HTTP(S), not accessible by JavaScript
    })

 
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// Custom claims structure
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
func GenerateJWT(username string) (string, error) {
    // Set token claims
    expirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 24hrs
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            Issuer:    "TutoringSite", 
        },
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign token and return
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err // Handle any error that occurred in signing the token
    }

    return tokenString, nil
}

func CheckPasswordHash(password string, expectedPasswordHash string)(bool){
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(expectedPasswordHash))

    return err == nil
}

func GetUserPasswordHash(username string) (string, error) {
    var user models.User


    result := initializers.DB.Where("username = ?", username).First(&user)

    if result.Error != nil {
        
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return "", fmt.Errorf("username not found")
        }
        
        return "", result.Error
    }

    return user.PasswordHash, nil
}


func TokenVerifyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extracting the token from the cookie
        cookie, err := r.Cookie("token")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenString := cookie.Value
        claims := &Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        
        var checkUser models.User
        initializers.DB.Where("username = ?", claims.Username).First(&checkUser)

        // Check if the user exists and is valid
        if checkUser.Username != claims.Username {
            http.Error(w, "Unauthorized: User not found", http.StatusUnauthorized)
            return
        }

        // If the token is valid, call the next handler
        next.ServeHTTP(w, r)
    })
}

