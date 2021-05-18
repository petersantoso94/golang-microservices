package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	userCtr "github.com/petersantoso94/golang-microservices/web-api/user"
)

var (
	servicePort = os.Getenv("WEB_API_PORT")
	userSvcUrl  = os.Getenv("USER_GRPC_ADDR")
	movieSvcUrl = os.Getenv("MOVIE_GRPC_ADDR")
)

func main() {
	userApi := userCtr.NewController(&userCtr.UserApiConn{MovieServiceUrl: movieSvcUrl, UserServiceUrl: userSvcUrl})

	engine := gin.New()
	engine.Use(gin.Recovery())

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	engine.Use(cors.New(cfg))

	baseGroup := engine.Group("/api/v1")
	{
		userGroup := baseGroup.Group("/users")
		{
			userGroup.GET("", userApi.GetUsers)
		}
	}

	if err := engine.Run(":" + servicePort); err != nil {
		log.Fatalln(err)
	}

}
