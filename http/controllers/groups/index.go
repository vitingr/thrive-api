package controllers

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando todos os grupos...")
	var groups []models.Group

	query := "SELECT * FROM groups"
	result := database.DB.Raw(query).Scan(&groups)
	if result.Error != nil {
		http.Error(w, "Failed to fetch groups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func GetGroupById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando grupo por ID...")
	vars := mux.Vars(r)
	id := vars["id"]

	var currentGroup models.Group
	query := fmt.Sprintf("SELECT * FROM groups WHERE id = %s LIMIT 1", id)
	result := database.DB.Raw(query).Scan(&currentGroup)
	if result.Error != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentGroup)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando novo grupo...")
	var newGroup models.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
		INSERT INTO groups (name, description, activities, group_picture, background_picture, is_private, followers, locale, members)
		VALUES ('%s', '%s', '%s', '%s', '%s', %t, %d, '%s', %d)
	`,
		newGroup.Name, newGroup.Description, newGroup.Activities,
		newGroup.GroupPicture, newGroup.BackgroundPicture,
		newGroup.IsPrivate, newGroup.Followers, newGroup.Locale, newGroup.Members,
	)

	result := database.DB.Exec(query)
	if result.Error != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newGroup)
}
