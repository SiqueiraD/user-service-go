package main

import (
	"github.com/siqueirad/user-service-go/database"
	"github.com/siqueirad/user-service-go/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
