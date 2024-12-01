package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"main/database"
	"main/models"
	"main/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserByGoogleID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	googleID := vars["google_id"]
	var currentUser models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE google_id = '%s' LIMIT 1", googleID)
	result := database.DB.Raw(query).Scan(&currentUser)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, currentUser, nil, "")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var newUser models.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
		INSERT INTO users (username, firstname, lastname, email, profile_picture, background_picture, followers, following, locale, google_id)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', '%s')
	`,
		newUser.Username, newUser.Firstname, newUser.Lastname, newUser.Email,
		newUser.ProfilePicture, newUser.BackgroundPicture,
		newUser.Followers, newUser.Following, newUser.Locale, newUser.GoogleID,
	)

	result := database.DB.Exec(query)
	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &newUser, nil, "")
}
