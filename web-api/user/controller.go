package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	movieClient "github.com/petersantoso94/golang-microservices/movie-service/grpc/client"
	userSvc "github.com/petersantoso94/golang-microservices/user-service"
	userClient "github.com/petersantoso94/golang-microservices/user-service/grpc/client"
	errHandler "github.com/petersantoso94/golang-microservices/web-api/common"
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
		errHandler.ErrorHandler(ctx,http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *UserController) GetMoviesByUserId(ctx *gin.Context) {
	userId,errC := strconv.ParseInt(ctx.Param("id"),10,64)
	if  errC != nil {
		errHandler.ErrorHandler(ctx,http.StatusBadRequest,errC)
		return
	}
	res, err := u.movieService.GetUserMovie(userId)
	if err != nil {
		errHandler.ErrorHandler(ctx,http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
