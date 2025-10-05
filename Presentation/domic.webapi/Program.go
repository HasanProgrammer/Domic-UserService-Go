package main

import (
	"Domic.Infrastructure/concretes"
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

	//dependencies

	unitOfWork, err := concretes.NewUnitOfWork("")

	server := grpc.NewServer()

	userRpc.RegisterUserServiceServer(server, &endpoints.UserServer{
		UnitOfWork:  unitOfWork,
		IdGenerator: concretes.NewIdentityGenerator(),
	})

	log.Println("gRPC server listening on :1996")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
