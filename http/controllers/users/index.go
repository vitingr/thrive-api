package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	var users []models.User

	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	vars := mux.Vars(r)
	email := vars["email"]
	var currentUser models.User

	result := database.DB.Where("email = ?", email).First(&currentUser)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, currentUser, nil, "")
}

func GetUserByGoogleID(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	vars := mux.Vars(r)
	googleId := vars["google_id"]
	var currentUser models.User

	result := database.DB.Where("email = ?", googleId).First(&currentUser)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, currentUser, nil, "")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	vars := mux.Vars(r)
	id := vars["id"]
	var currentUser models.User
	database.DB.First(&currentUser, id)

	utils.SendResponse(w, http.StatusOK, &currentUser, nil, "")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
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

	utils.SendResponse(w, http.StatusOK, &newUser, nil, "")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	result := database.DB.First(&existingUser, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	result = database.DB.Model(&existingUser).Updates(updatedUser)
	if result.Error != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &existingUser, nil, "")
}

func GetSuggestedFriends(w http.ResponseWriter, r *http.Request) {
	database.DB.Exec("DEALLOCATE ALL")
	vars := mux.Vars(r)
	userID := vars["id"]

	var suggestedUsers []models.User

	query := `
    SELECT * 
    FROM users 
    WHERE id != $1 
    AND id NOT IN (
        SELECT following_id 
        FROM followers 
        WHERE follower_id = $2
    )
`

	result := database.DB.Raw(query, userID, userID).Scan(&suggestedUsers)
	if result.Error != nil {
		http.Error(w, "Failed to fetch suggested friends", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, suggestedUsers, nil, "")
}
