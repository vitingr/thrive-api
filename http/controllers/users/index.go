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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando todos os usuários...")
	var users []models.User

	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando usuário por email...")
	vars := mux.Vars(r)
	email := vars["email"]
	var currentUser models.User
	database.DB.First(&currentUser, email)

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
	fmt.Println("Criando novo usuário...")
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	database.DB.Create(&newUser)
	json.NewEncoder(w).Encode(newUser)
}
