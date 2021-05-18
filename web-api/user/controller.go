package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	movieClient "github.com/petersantoso94/golang-microservices/movie-service/grpc/client"
	userSvc "github.com/petersantoso94/golang-microservices/user-service"
	userClient "github.com/petersantoso94/golang-microservices/user-service/grpc/client"
)

type UserController struct {
	movieService movieSvc.Service
	userService  userSvc.Service
}

type UserApiConn struct {
	UserServiceUrl  string
	MovieServiceUrl string
}

func NewController(c *UserApiConn) *UserController {
	userClt, err := userClient.NewGRPCService(c.UserServiceUrl)
	if err != nil {
		log.Fatalln(err)
	}
	movieClt, err := movieClient.NewGRPCService(c.MovieServiceUrl)
	if err != nil {
		log.Fatalln(err)
	}

	return &UserController{movieService: movieClt, userService: userClt}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	res, err := u.userService.GetUsers()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}
