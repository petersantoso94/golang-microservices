package client

import (
	"context"
	"time"

	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	pb "github.com/petersantoso94/golang-microservices/movie-service/grpc"
	"google.golang.org/grpc"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient pb.MovieServiceClient
}

// NewGRPCService creates a new gRPC movie service connection using the specified connection string.
func NewGRPCService(connString string) (movieSvc.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: pb.NewMovieServiceClient(conn)}, nil
}
func (s *grpcService) GetMovies() (result []*movieSvc.Movie, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetMovies(ctx, &pb.EmptyRequest{})
	if err != nil {
		return
	}
	for _, grpcMovie := range resp.GetMovies() {
		u := unmarshalMovie(grpcMovie)
		result = append(result, u)
	}
	return
}
func (s *grpcService) GetUserMovie(OwnerID int64) (result []*movieSvc.Movie, err error) {
	req := &pb.GetUserMovieRequest{
		UserId: OwnerID,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUserMovie(ctx, req)
	if err != nil {
		return
	}
	for _, m := range resp.GetMovies() {
		u := unmarshalMovie(m)
		result = append(result, u)
	}
	return
}
func (s *grpcService) CreateMovie(movie movieSvc.Movie) error {
	req := &pb.CreateMovieRequest{
		Movie: marshalMovie(&movie),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	_, err := s.grpcClient.CreateMovie(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
func marshalMovie(movie *movieSvc.Movie) *pb.Movie {
	return &pb.Movie{Id: int64(movie.ID), Name: movie.Name, OwnerID: movie.OwnerID}
}
func unmarshalMovie(grpcMovie *pb.Movie) *movieSvc.Movie {
	return &movieSvc.Movie{ID: grpcMovie.Id, Name: grpcMovie.Name, OwnerID: grpcMovie.OwnerID}
}
