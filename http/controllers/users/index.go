package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    email := vars["email"]
    var currentUser models.User
		
    result := database.DB.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&currentUser)

    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(currentUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando usuário por ID...")
	vars := mux.Vars(r)
	id := vars["id"]
	var currentUser models.User
	database.DB.First(&currentUser, id)

	json.NewEncoder(w).Encode(currentUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var newUser models.User

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	database.DB.Create(&newUser)
	json.NewEncoder(w).Encode(newUser)
}
