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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	var currentUser models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1", email)
	result := database.DB.Raw(query).Scan(&currentUser)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, currentUser, nil, "")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var currentUser models.User

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s LIMIT 1", id)
	result := database.DB.Raw(query).Scan(&currentUser)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendResponse(w, http.StatusOK, &currentUser, nil, "")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
		UPDATE users
		SET username = '%s', firstname = '%s', lastname = '%s', email = '%s',
			profile_picture = '%s', background_picture = '%s', followers = %d,
			following = %d, locale = '%s', google_id = '%s'
		WHERE id = %s
	`,
		updatedUser.Username, updatedUser.Firstname, updatedUser.Lastname, updatedUser.Email,
		updatedUser.ProfilePicture, updatedUser.BackgroundPicture,
		updatedUser.Followers, updatedUser.Following, updatedUser.Locale, updatedUser.GoogleID, id,
	)

	result := database.DB.Exec(query)
	if result.Error != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, &updatedUser, nil, "")
}

func GetSuggestedFriends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var suggestedUsers []models.User

	query := fmt.Sprintf(`
		SELECT *
		FROM users
		WHERE id != %s
		AND id NOT IN (
			SELECT following_id
			FROM followers
			WHERE follower_id = %s
		)
	`, userID, userID)

	result := database.DB.Raw(query).Scan(&suggestedUsers)
	if result.Error != nil {
		http.Error(w, "Failed to fetch suggested friends", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, http.StatusOK, suggestedUsers, nil, "")
}
