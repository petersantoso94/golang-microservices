package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	userCtr "github.com/petersantoso94/golang-microservices/web-api/user"
	movieCtr "github.com/petersantoso94/golang-microservices/web-api/movie"
)

var (
	servicePort = os.Getenv("WEB_API_PORT")
	userSvcUrl  = os.Getenv("USER_GRPC_ADDR")
	movieSvcUrl = os.Getenv("MOVIE_GRPC_ADDR")
)

func main() {
	userApi := userCtr.NewController(&userCtr.UserApiConn{MovieServiceUrl: movieSvcUrl, UserServiceUrl: userSvcUrl})
	movieApi := movieCtr.NewController(&movieCtr.MovieApiConn{MovieServiceUrl: movieSvcUrl})

	engine := gin.New()
	engine.Use(gin.Recovery())

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	engine.Use(cors.New(cfg))

	baseGroup := engine.Group("/api/v1")
	{
		usersGroup := baseGroup.Group("/users")
		{
			usersGroup.GET("", userApi.GetUsers)
			usersGroup.GET("/:id/movies", userApi.GetMoviesByUserId)
		}
		moviesGroup := baseGroup.Group("/movies")
		{
			moviesGroup.GET("", movieApi.GetMovies)
		}
	}

	if err := engine.Run(":" + servicePort); err != nil {
		log.Fatalln(err)
	}

}
