package twitchapi

import (
    "context"
    "fmt"
    "twirp-oauth-backend/rpc/twirpAPI"
)

type Server struct {}

func (s *Server) Login(ctx context.Context, req *twirpAPI.LoginRequest) (*twirpAPI.LoginResponse, error) {

    fmt.Println("Does it hit here")
    return &twirpAPI.LoginResponse{
        UserId: "1234",
    }, nil

}
