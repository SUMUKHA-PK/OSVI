package main

import (
	"log"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/OSVI/backend/routing"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r = routing.SetupRouting(r)

	err := server.Server(r, "55555", false)
	if err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
