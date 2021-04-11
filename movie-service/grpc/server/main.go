package main

import (
	"log"
	"net"
	"os"

	coreSvc "github.com/petersantoso94/golang-microservices/movie-service/core"
	grpcSvc "github.com/petersantoso94/golang-microservices/movie-service/grpc"
	controller "github.com/petersantoso94/golang-microservices/movie-service/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	movieService := coreSvc.NewService()
	// configure our gRPC service controller
	userGrpcAddr := os.Getenv("GRPC_ADDR")
	movieServiceController := controller.NewMovieServiceController(movieService, &controller.ConnConfig{GrpcUserHost: userGrpcAddr})
	// start a gRPC server
	server := grpc.NewServer()
	grpcSvc.RegisterMovieServiceServer(server, movieServiceController)
	reflection.Register(server)
	grpcAddr := os.Getenv("MOVIE_GRPC_ADDR")
	log.Printf("running movie-grpc service at: %s\n", grpcAddr)
	con, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
