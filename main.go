package main

import (
	"MyGram/database"
	"MyGram/router"
)

func main() {
	database.StartDB()

	r := router.StartDB()
	r.Run(":8080")
}
