package controller

import (
	"context"
	"log"

	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	pb "github.com/petersantoso94/golang-microservices/movie-service/grpc"
	userSvc "github.com/petersantoso94/golang-microservices/user-service"
	userClient "github.com/petersantoso94/golang-microservices/user-service/grpc/client"
)

type ConnConfig struct {
	GrpcUserHost string
}

// MovieServiceController implements the gRPC MovieServiceServer interface.
type MovieServiceController struct {
	movieService movieSvc.Service
	userService  userSvc.Service
	pb.UnimplementedMovieServiceServer
}

// NewMovieServiceController instantiates a new MovieServiceServer.
func NewMovieServiceController(movieService movieSvc.Service, conn *ConnConfig) *MovieServiceController {
	userService, err := userClient.NewGRPCService(conn.GrpcUserHost)
	if err != nil {
		log.Fatalln(err)
	}
	return &MovieServiceController{
		movieService: movieService,
		userService:  userService,
	}
}

func (s *MovieServiceController) GetMovies(ctx context.Context, in *pb.EmptyRequest) (*pb.GetMoviesResponse, error) {
	movies, err := s.movieService.GetMovies()
	if err != nil {
		return nil, err
	}
	if len(movies) == 0 {
		return nil, movieSvc.ErrNotFound
	}
	var grpcMovie []*pb.Movie
	log.Printf("movies:%v", movies)
	for _, movie := range movies {
		u := marshalMovie(movie)
		grpcMovie = append(grpcMovie, u)
	}
	return &pb.GetMoviesResponse{Movies: grpcMovie}, nil
}

func (s *MovieServiceController) GetUserMovie(ctx context.Context, in *pb.GetUserMovieRequest) (*pb.GetUserMovieResponse, error) {
	movies, err := s.movieService.GetUserMovie(in.UserId)
	if err != nil {
		return nil, err
	}
	if len(movies) == 0 {
		return nil, movieSvc.ErrNotFound
	}
	var grpcMovie []*pb.Movie
	log.Printf("movies:%v", movies)
	for _, movie := range movies {
		u := marshalMovie(movie)
		grpcMovie = append(grpcMovie, u)
	}
	return &pb.GetUserMovieResponse{Movies: grpcMovie}, nil
}

func (s *MovieServiceController) CreateMovie(ctx context.Context, in *pb.CreateMovieRequest) (*pb.EmptyResponse, error) {
	_, errUser := s.userService.GetUser(in.Movie.OwnerID)
	if errUser != nil {
		return nil, errUser
	}
	err := s.movieService.CreateMovie(unmarshalMovie(in.Movie))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func marshalMovie(movie *movieSvc.Movie) *pb.Movie {
	return &pb.Movie{Id: int64(movie.ID), Name: movie.Name, OwnerID: movie.OwnerID}
}
func unmarshalMovie(grpcMovie *pb.Movie) movieSvc.Movie {
	return movieSvc.Movie{ID: grpcMovie.Id, Name: grpcMovie.Name, OwnerID: grpcMovie.OwnerID}
}
