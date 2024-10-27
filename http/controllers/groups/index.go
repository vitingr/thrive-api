package controllers

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Thrive default API route")
	fmt.Println("testando")
}

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando todos os grupos...")
	var groups []models.Group

	database.DB.Find(&groups)
	json.NewEncoder(w).Encode(groups)
}

func GetGroupById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando grupo por ID...")
	vars := mux.Vars(r)
	id := vars["id"]
	var currentGroup models.Group
	database.DB.First(&currentGroup, id)

	json.NewEncoder(w).Encode(currentGroup)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando novo grupo...")
	var newGroup models.Group
	json.NewDecoder(r.Body).Decode(&newGroup)
	database.DB.Create(&newGroup)
	json.NewEncoder(w).Encode(newGroup)
}
