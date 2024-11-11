package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"search-api/clients/queues"
	controllers "search-api/controllers"
	repositories "search-api/repositories"
	services "search-api/services"
)

func main() {
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
		Username:  "guest",
		Password:  "guest",
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
	router := gin.Default()
	router.GET("/search", controller.Search)
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
