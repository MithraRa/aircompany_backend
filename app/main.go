package main

import (
	"airline"
	"airline/database"
	"airline/routes"
	_ "github.com/lib/pq"
	"log"
)

const serverPort = "8080"

func main() {
	// Creating a connection to the DB
	database.Connect()
	// Creating a set of routes
	root := routes.InitRoutes()
	// Creating a server
	srv := new(server.Server)
	// Run it till it stops
	if err := srv.Run(serverPort, root); err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
}
