package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexagonal-architecture/adapter/repository"
	"github.com/rinonkia/go-hexagonal-architecture/handler"
)

func main() {
	r := gin.Default()

	// repository
	userRepository := repository.NewInMemoryUserRepository()

	// handler
	healthCheckHandler := handler.HealthCheck()
	signupHandler := handler.Signup(userRepository)
	getUsersHandler := handler.GetUsers(userRepository)
	loginHandler := handler.Login(userRepository)
	logoutHandler := handler.Logout()

	r.GET("/health", healthCheckHandler)
	r.POST("/signup", signupHandler)
	r.GET("/users", getUsersHandler)
	r.POST("/login", loginHandler)
	r.GET("/logout", logoutHandler)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
