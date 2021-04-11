package main

import (
	"log"
	"net"
	"os"

	coreSvc "github.com/petersantoso94/golang-microservices/user-service/core"
	grpcSvc "github.com/petersantoso94/golang-microservices/user-service/grpc"
	controller "github.com/petersantoso94/golang-microservices/user-service/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	userService := coreSvc.NewService()
	// configure our gRPC service controller
	userServiceController := controller.NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	grpcSvc.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)
	grpcAddr := os.Getenv("GRPC_ADDR")
	log.Printf("running user-grpc service at: %s\n", grpcAddr)
	con, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
