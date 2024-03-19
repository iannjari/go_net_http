package main

import (
	"go_net_http/api"
	"go_net_http/database"
	"net/http"
)

func main() {

	dbClient := database.GetDB()

	http.ListenAndServe("localhost:8080", api.NewServer(dbClient))
}
