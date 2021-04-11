package client

import (
	"context"
	"time"

	user "github.com/petersantoso94/golang-microservices/user-service"
	usergrpc "github.com/petersantoso94/golang-microservices/user-service/grpc"
	"google.golang.org/grpc"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient usergrpc.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (user.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: usergrpc.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers() (result []user.User, err error) {
	req := &usergrpc.GetUsersRequest{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		u := unmarshalUser(grpcUser)
		result = append(result, u)
	}
	return
}
func (s *grpcService) GetUser(id int64) (result user.User, err error) {
	req := &usergrpc.GetUserRequest{
		Id: int64(id),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUser(ctx, req)
	if err != nil {
		return
	}
	result = unmarshalUser(resp.GetUser())
	return result, nil
}
func (s *grpcService) CreateUser(user user.User) (id int64, err error) {
	req := &usergrpc.CreateUserRequest{
		User: marshalUser(&user),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.CreateUser(ctx, req)
	if err != nil {
		return
	}
	id = resp.Id
	return
}
func unmarshalUser(grpcUser *usergrpc.User) (result user.User) {
	result.ID = grpcUser.Id
	result.Name = grpcUser.Name
	return
}
func marshalUser(user *user.User) (result *usergrpc.User) {
	result.Id = int64(user.ID)
	result.Name = user.Name
	return
}
