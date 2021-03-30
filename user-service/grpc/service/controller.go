package controller
import (
  "context"
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
func (ctlr *UserServiceController) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (resp *pb.GetUsersResponse, err error) {
  resultMap, err := ctlr.userService.GetUsers(req.GetIds())
  if err != nil {
    return
  }
  resp = &pb.GetUsersResponse{}
  for _, u := range resultMap {
    resp.Users = append(resp.Users, marshalUser(&u))
  }
  return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *userSvc.User) *pb.User {
  return &pb.User{Id: u.ID, Name: u.Name}
}