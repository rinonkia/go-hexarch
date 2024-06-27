package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexagonal-architecture/adapter/repository"
	"github.com/rinonkia/go-hexagonal-architecture/handler"
)

func main() {
	r := gin.Default()

	userRepository := repository.NewInMemoryUserRepository()

	healthCheckHandler := handler.HealthCheck()
	signupHandler := handler.Signup(userRepository)
	getUsersHandler := handler.GetUsers(userRepository)

	r.GET("/health", healthCheckHandler)
	r.POST("/signup", signupHandler)
	r.GET("/users", getUsersHandler)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
