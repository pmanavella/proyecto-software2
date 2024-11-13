package main

import (
	clients "courses-api/clients/queues"
	"courses-api/handlers"
	"courses-api/repositories"
	"courses-api/services"
	"log"

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
	// Configuración de MongoDB
	repository := repositories.NewMongo(repositories.MongoConfig{
		Host:       "mongodb",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "courses",
		Collection: "courses",
	})

	// Configuración de RabbitMQ
	rabbitConfig := clients.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "guest",
		Password:  "guest",
		QueueName: "course_updates",
	}
	rabbit := clients.NewRabbit(rabbitConfig)

	// Configuración de Servicios y Handlers
	service := services.NewCourseService(repository, rabbit)
	handler := handlers.NewHandler(service)

	// Configuración del router de Gin
	router := gin.Default()
	router.Use(AllowCORS) // Aplica la función AllowCORS a todas las rutas

	// Configuración del Router
	
	// router.GET("/courses/:id", handler.GetCourseByID)
	router.POST("/createCourse", handler.CreateCourse)
	router.PUT("/edit/:course_id", handler.UpdateCourse)
	router.DELETE("/delete/:course_id", handler.DeleteCourse)
	router.GET("/search", handler.SearchByTitle)

	/////////
	router.GET("/course", AllowCORS, handler.GetAll)
	router.GET("/course/:id", AllowCORS,handler.GetCourseByID)
	router.GET("/course/title=:title", handler.SearchByTitle)
	router.GET("/course/category=:category", handler.SearchByCategory)
	router.GET("/course/description=:description", handler.SearchByDescription)

	//router.PUT("/course/:id_course", courseController.PutCourse)
	router.PUT("/course/:id_course", handler.Update) // no debería solo dejar este put? Igual dentro de Put se llama a update

	router.DELETE("/course", handler.Delete)

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("error running application: %v", err)
	}
}

/////// alternative

// package main

// import (
//     "courses-api/clients"
//     "courses-api/handlers"
//     "courses-api/repositories"
//     "courses-api/services"
//     "github.com/gin-gonic/gin"
//     "log"
// )

// func main() {
//     // Configuración de MongoDB
//     mongoConfig := repositories.MongoConfig{
//         Host:       "localhost",
//         Port:       "27017",
//         Username:   "root",
//         Password:   "root",
//         Database:   "courses",
//         Collection: "courses",
//     }
//     mainRepository := repositories.NewCourseRepository(mongoConfig)

//     // Configuración de RabbitMQ
//     rabbitConfig := clients.RabbitConfig{
//         Host:      "localhost",
//         Port:      "5672",
//         Username:  "guest",
//         Password:  "guest",
//         QueueName: "course_updates",
//     }
//     rabbit := clients.NewRabbit(rabbitConfig)

//     // Configuración de Servicios y Handlers
//     service := services.NewCourseService(mainRepository, rabbit)
//     handler := handlers.NewHandler(service)

//     // Configuración del Router
//     router := gin.Default()
//     router.GET("/courses/:id", handler.GetCourseByID)
//     router.POST("/createCourse", handler.CreateCourse)
//     router.PUT("/edit/:course_id", handler.UpdateCourse)
//     router.DELETE("/delete/:course_id", handler.DeleteCourse)
//     router.GET("/search", handler.SearchCoursesByTitle)
//     router.POST("/enroll", handler.CreateEnrollment)
//     router.GET("/availability", handler.CalculateAvailability)

//     if err := router.Run(":8081"); err != nil {
//         log.Fatalf("error running application: %v", err)
//     }
// }
