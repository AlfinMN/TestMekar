package main

import (
	"BackendTestMekar/config"
	"BackendTestMekar/domain"
)

func main() {
	db, _, dbHost, portServer := config.Connection()
	router := config.CreateRouter()

	domain.Init(router, db)
	config.RunServer(router, dbHost, portServer)
}
