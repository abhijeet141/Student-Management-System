package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtsecret := os.Getenv("JWT_SECRET")
	if jwtsecret == "" {
		log.Fatal("Connection String is not set in the environment")
	}

	jwtSecret = []byte(jwtsecret)
}

func GenerateJWT(userName string) (string, error) {
	claims := jwt.MapClaims{
		"username": userName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}
	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	o := orm.NewOrm()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	userInfo := models.User{UserName: user.UserName}
	err = o.Read(&userInfo, "UserName")
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	if user.Password != userInfo.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(userInfo.UserName)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
