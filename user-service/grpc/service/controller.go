package controller

import (
	"context"
	"log"

	userSvc "github.com/petersantoso94/golang-microservices/user-service"
	pb "github.com/petersantoso94/golang-microservices/user-service/grpc"
)

// UserServiceController implements the gRPC UserServiceServer interface.
type UserServiceController struct {
	userService userSvc.Service
	pb.UnimplementedUserServiceServer
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService userSvc.Service) *UserServiceController {
	return &UserServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (s *UserServiceController) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := s.userService.GetUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, userSvc.ErrNotFound
	}
	var grpcUser []*pb.User
	log.Printf("users:%v", users)
	for _, user := range users {
		u := marshalUser(user)
		grpcUser = append(grpcUser, u)
	}
	return &pb.GetUsersResponse{Users: grpcUser}, nil
}
func (s *UserServiceController) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.GetUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: marshalUser(user)}, nil
}
func (s *UserServiceController) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userGrpc := in.User
	_, err := s.userService.CreateUser(unmarshalUser(userGrpc))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: int64(userGrpc.Id)}, nil
}
func unmarshalUser(grpcUser *pb.User) userSvc.User {
	return userSvc.User{ID: grpcUser.Id, Name: grpcUser.Name}
}
func marshalUser(user *userSvc.User) *pb.User {
	return &pb.User{Id: int64(user.ID), Name: user.Name}
}
