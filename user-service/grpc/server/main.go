package main
import (
  "net"
  "os"
  controller "github.com/petersantoso94/golang-microservices/user-service/grpc/service"
  coreSvc "github.com/petersantoso94/golang-microservices/user-service/core"
  grpcSvc "github.com/petersantoso94/golang-microservices/user-service/grpc"
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
  con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
  if err != nil {
    panic(err)
  }
  err = server.Serve(con)
  if err != nil {
    panic(err)
  }
}