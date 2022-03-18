package controller

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	movieClient "github.com/petersantoso94/golang-microservices/movie-service/grpc/client"
	errHandler "github.com/petersantoso94/golang-microservices/web-api/common"
)

type MovieController struct {
	movieService movieSvc.Service
}

type MovieApiConn struct {
	MovieServiceUrl string
}

func NewController(c *MovieApiConn) *MovieController {
	movieClt, err := movieClient.NewGRPCService(c.MovieServiceUrl)
	if err != nil {
		log.Fatalln(err)
	}

	return &MovieController{movieService: movieClt}
}

func (u *MovieController) GetMovies(ctx *gin.Context) {
	res, err := u.movieService.GetMovies()
	if err != nil {
		errHandler.ErrorHandler(ctx,http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}