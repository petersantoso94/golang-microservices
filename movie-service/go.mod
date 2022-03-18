module github.com/petersantoso94/golang-microservices/movie-service

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/petersantoso94/golang-microservices/user-service v0.0.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.7
)

replace github.com/petersantoso94/golang-microservices/user-service => ../user-service
