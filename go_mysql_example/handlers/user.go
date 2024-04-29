package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"

    "githhub.com/developerwritings/go_development/go_mysql_example/config"
    "githhub.com/developerwritings/go_development/go_mysql_example/models"
)

var jwtKey = []byte(config.SecretKey)

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func generateToken(username string) (string, error) {
    expirationTime := time.Now().Add(15 * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func SignUp(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword

    err = models.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    token, err := generateToken(user.Username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    savedUser, err := models.GetUserByUsername(user.Username)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    err = comparePasswords(savedUser.Password, user.Password)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    token, err := generateToken(savedUser.Username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract the username from the token
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    username := claims.Username

    user, err := models.GetUserByUsername(username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Extract the username from the token
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    username := claims.Username

    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    // Make sure the user in the token matches the user being updated
    if username != user.Username {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    err = models.UpdateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Extract the username from the token
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    username := claims.Username

    err = models.DeleteUser(username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User deleted successfully")
}
