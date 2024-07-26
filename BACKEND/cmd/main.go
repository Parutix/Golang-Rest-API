package main

import (
	"log"

	"gitub.com/Parutix/Golang-Rest-API/cmd/api"
)

func main() {
	server := api.CreateNewAPIServer(":8081", nil)
	log.Fatal(server.Run())
}