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
func (s *grpcService) GetUsers(ids []int64) (result map[int64]user.User, err error) {
  result = map[int64]user.User{}
  req := &usergrpc.GetUsersRequest{
    Ids: ids,
  }
  ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
  defer cancelFunc()
  resp, err := s.grpcClient.GetUsers(ctx, req)
  if err != nil {
    return
  }
  for _, grpcUser := range resp.GetUsers() {
    u := unmarshalUser(grpcUser)
    result[u.ID] = u
  }
  return
}
func (s *grpcService) GetUser(id int64) (result user.User, err error) {
  req := &usergrpc.GetUsersRequest{
    Ids: []int64{id},
  }
  ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
  defer cancelFunc()
  resp, err := s.grpcClient.GetUsers(ctx, req)
  if err != nil {
    return
  }
  for _, grpcUser := range resp.GetUsers() {
    // sanity check: only the requested user should be present in results
    if grpcUser.GetId() == id {
      return unmarshalUser(grpcUser), nil
    }
  }
  return result, user.ErrNotFound
}
func unmarshalUser(grpcUser *usergrpc.User) (result user.User) {
  result.ID = grpcUser.Id
  result.Name = grpcUser.Name
  return
}