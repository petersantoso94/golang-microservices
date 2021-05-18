module github.com/petersantoso94/golang-microservices/web-api

go 1.16

replace github.com/petersantoso94/golang-microservices/user-service => ../user-service

replace github.com/petersantoso94/golang-microservices/movie-service => ../movie-service

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/petersantoso94/golang-microservices/movie-service v0.0.0-00010101000000-000000000000
	github.com/petersantoso94/golang-microservices/user-service v0.0.0
)
