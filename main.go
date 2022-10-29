package main

import (
	"affan/database"
	"affan/router"
)

func main() {
	database.DatabaseConnection()
	r := router.StartApp()
	r.Run(":8080")
}