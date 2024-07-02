package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/go-hexarch/adapter/repository"
	"github.com/rinonkia/go-hexarch/adapter/service"
	"github.com/rinonkia/go-hexarch/config"
	"github.com/rinonkia/go-hexarch/handler"
	"github.com/rinonkia/go-hexarch/handler/middleware"
)

func main() {
	c := config.GetEnvConfig()

	// repository
	userRepository := repository.NewInMemoryUserRepository()

	// service
	tokenGenerator := service.NewTokenGenerator(c.SecretKey)

	// middleware
	checkAuthorizationMiddleware := middleware.CheckAuthorization(c.SecretKey)

	// handler
	healthCheckHandler := handler.HealthCheck()
	signupHandler := handler.Signup(tokenGenerator, userRepository)
	getUsersHandler := handler.GetUsers(userRepository)
	loginHandler := handler.Login(tokenGenerator, userRepository)

	r := gin.Default()
	r.Use(middleware.SecureHeader(c))
	r.GET("/health", healthCheckHandler)
	r.POST("/signup", signupHandler)
	r.GET("/users", checkAuthorizationMiddleware, getUsersHandler)
	r.POST("/login", loginHandler)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
