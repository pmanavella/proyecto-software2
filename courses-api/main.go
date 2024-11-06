package main

import (
    "courses-api/clients"
    "courses-api/handlers"
    "courses-api/repositories"
    "courses-api/services"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Configuración de MongoDB
    repository := repositories.NewMongo(repositories.MongoConfig{
        Host:       "localhost",
        Port:       "27017",
        Username:   "root",
        Password:   "root",
        Database:   "courses",
        Collection: "courses",
    })

    // Configuración de RabbitMQ
    rabbitConfig := clients.RabbitConfig{
        Host:      "localhost",
        Port:      "5672",
        Username:  "guest",
        Password:  "guest",
        QueueName: "course_updates",
    }
    rabbit := clients.NewRabbit(rabbitConfig)

    // Configuración de Servicios y Handlers
    service := services.NewCourseService(repository, rabbit)
    handler := handlers.NewHandler(service)

    // Configuración del Router
    router := gin.Default()
    router.GET("/courses/:id", handler.GetCourseByID)
    router.POST("/createCourse", handler.CreateCourse)
    router.PUT("/edit/:course_id", handler.UpdateCourse)
    router.DELETE("/delete/:course_id", handler.DeleteCourse)
    router.GET("/search", handler.SearchCoursesByTitle)
    router.POST("/enroll", handler.CreateEnrollment)
    router.GET("/availability", handler.CalculateAvailability)

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