package main

import (
	"fmt"
	"main/database"
	"main/http/routes"
)

func main() {
	database.ConnectionWithDatabase()
	fmt.Println("Starting Thrive rest server :D")
	routes.HandleRequest()
}