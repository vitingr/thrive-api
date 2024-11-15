package controllers

import (
	"bytes"
	"encoding/json"
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

func GetUserByGoogleID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["google_id"]
	var currentUser models.User
	
	result := database.DB.Raw("SELECT * FROM users WHERE google_id = ? LIMIT 1", email).Scan(&currentUser)

	if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(currentUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
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

	query := `
		INSERT INTO users (username, firstname, lastname, email, profile_picture, background_picture, followers, following, locale, google_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result := database.DB.Exec(query, newUser.Username, newUser.Firstname, newUser.Lastname, newUser.Email, newUser.ProfilePicture, newUser.BackgroundPicture, newUser.Followers, newUser.Following, newUser.Locale, newUser.GoogleID)

	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := struct {
		Data  *models.User `json:"data,omitempty"`
		Meta  interface{}   `json:"meta,omitempty"`
		Error string        `json:"error,omitempty"`
	}{
		Data:  &newUser,
		Meta:  nil, 
		Error: "",  
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
