package main

import (
	"learnWednesday/handlers"
	"learnWednesday/service"
)

const (
	ErrDatabase = "Error connecting to MySQL database"
)

func main() {
	// Setting up MySQL database
	mySQL, err := connectMySQL()

	// If there is an error connecting to mySQL database - panic and stop the program
	if err != nil {
		panic(ErrDatabase)
	}

	// Setting up mysql store to implement interface functions
	db := service.NewMySQLStore(mySQL)

	// Setting up learnWednesday Client
	learnWednesday := service.NewLearnWednesdayService(db)

	// Setting up a Router to call API endpoints
	handlers.SetupRouter(learnWednesday)
}
