package main

import (
	"domic.webapi/endpoints"
	userRpc "github.com/HasanProgrammer/Domic-GrpcService-Go/UserService/compile/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":1996")

	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	server := grpc.NewServer()

	userRpc.RegisterUserServiceServer(server, &endpoints.UserServer{})

	log.Println("gRPC server listening on :1996")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
