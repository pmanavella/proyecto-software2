package main

import (
	"log"
	"search-api/clients/queues"
	controllers "search-api/controllers"
	repositories "search-api/repositories"
	services "search-api/services"

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

	// Solr
	solrRepo := repositories.NewSolr(repositories.SolrConfig{
		Host:       "solr",    // Solr host
		Port:       "8983",    // Solr port
		Collection: "courses", // Collection name
	})

	// Rabbit
	eventsQueue, err := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "root",
		Password:  "root",
		QueueName: "courses-news",
	})

	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
		return
	}

	log.Println(eventsQueue)

	// Courses API
	coursesAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "courses-api",
		Port: "8080",
	})

	// Services
	service := services.NewService(solrRepo, coursesAPI)

	// Controllers
	controller := controllers.NewCourseController(service)

	// Launch rabbit consumer
	if err := eventsQueue.StartConsumer(service.HandleCourseNew); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	// Create router
	router.GET("/search", controller.Search)
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
