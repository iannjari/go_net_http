package main

import (
	"go_net_http/api"
	"net/http"
)

func main() {

	http.ListenAndServe("localhost:8080", api.NewServer())

}
