package main

import (
	"context"
	"fmt"
	"gRPC/pkg/api"
	"log"
	"net"

	"google.golang.org/grpc"

)

type GRPCServer struct {
    api.AuthServer
    users map[string]struct{}
}

func (s GRPCServer) CheckAuth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
    username := req.GetUsername()
    fmt.Println(username)
    _, ok := s.users[username]
    if !ok {
        return &api.AuthResponse{Exists: false}, nil
    }

    return &api.AuthResponse{Exists: true}, nil
}

func (s GRPCServer) mustEmbedUnimplementedAuthServer(ctx context.Context, req *api.AuthRequest) {}

func main() {
    s := grpc.NewServer()

    users := make(map[string]struct{})
    users["Echo"] = struct{}{}
    users["Mercy"] = struct{}{}
    users["D.Va"] = struct{}{}

    srv := &GRPCServer{users: users}

    api.RegisterAuthServer(s, srv)

    l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сервер запущен")
    fmt.Println(srv.users)

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}