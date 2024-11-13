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

func AllowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}

func main() {

	// Configuración del router de Gin
	router := gin.Default()
	router.Use(AllowCORS) // Aplica la función AllowCORS a todas las rutas

	// MySQL
	mySQLRepo := repositories.NewMySQL(
		repositories.MySQLConfig{
			Host:     "mysql",
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
	controller := controllers.NewController(&service)

	// URL mappings
	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetByID)
	router.POST("/users", controller.Create)
	router.PUT("/users/:id", controller.Update)
	router.POST("/login", controller.Login)
	router.POST("/users/:userID/courses/:courseID/enroll", controller.InscriptionCourses)

	// Run application
	if err := router.Run(":8080"); err != nil {
		log.Panicf("Error running application: %v", err)
	}
}
