package main

import (
	"log"
	"time"
	controllers "users-api/controllers"
	tokenizers "users-api/internal"
	repositories "users-api/repositories"
	services "users-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// MySQL
	mySQLRepo := repositories.NewMySQL(
		repositories.MySQLConfig{
			Host:     "localhost",
			Port:     "3306",
			Database: "users_api",
			Username: "root",
			Password: "root",
		},
	)

	// Cache
	cacheRepo := repositories.NewCache(repositories.CacheConfig{
		TTL: 1 * time.Minute,
	})

	// Memcached
	memcachedRepo := repositories.NewMemcached(repositories.MemcachedConfig{
		Host: "memcached",
		Port: "11211",
	})

	// Tokenizer
	jwtTokenizer := tokenizers.NewTokenizer(
		tokenizers.JWTConfig{
			Key:      "ThisIsAnExampleJWTKey!",
			Duration: 1 * time.Hour,
		},
	)

	// Services
	service := services.NewService(mySQLRepo, cacheRepo, memcachedRepo, jwtTokenizer)

	// Handlers
	controller := controllers.NewController(service)

	// Create router
	router := gin.Default()

	// URL mappings
	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetByID)
	router.POST("/users", controller.Create)
	router.PUT("/users/:id", controller.Update)
	router.POST("/login", controller.Login)

	// Run application
	if err := router.Run(":8080"); err != nil {
		log.Panicf("Error running application: %v", err)
	}
}
