package main

import (
	"github.com/meajudaaqui/user-service/database"
	"github.com/meajudaaqui/user-service/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
