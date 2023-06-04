package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"UserServer/controllers/v1"
	"UserServer/constant"
	"UserServer/postgresql"
	"UserServer/proto/UserServer"
)

func init() {
	constant.ReadConfig(".env")
	postgresql.Initialize()
}

func main() {
	lis, err := net.Listen("tcp", ":1236")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	UserServer.RegisterUserServerServer(s, &v1.UserServe{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
