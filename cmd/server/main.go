package main

import (
	"fmt"
	"net/http"
	twirpServer "twirp-oauth-backend/internal/twirpAPI"
	"twirp-oauth-backend/rpc/twirpAPI"
)

func main() {
	fmt.Println("This is a test")

    server := &twirpServer.Server{}
    twirpHandler := twirpAPI.NewAuthServiceServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
