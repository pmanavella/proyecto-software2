package main

import (
    "courses-api/clients"
    "courses-api/handlers"
    "courses-api/repositories"
    "courses-api/queues"
    "log"
    "net/http"
)

func main() {
    // Configuración de MongoDB
    mongoURI := "mongodb://localhost:27017"
    client := clients.ConnectMongoDB(mongoURI)
    repo := repositories.NewCourseRepository(client, "coursesDB", "courses")

    // Configuración de RabbitMQ
    rabbitConfig := clients.RabbitConfig{
        Host:      "localhost",
        Port:      "5672",
        Username:  "guest",
        Password:  "guest",
        QueueName: "course_updates",
    }
    rabbit := clients.NewRabbit(rabbitConfig)

    // Configuración de Handlers
    handler := handlers.NewHandler(repo, rabbit)

    http.HandleFunc("/courses", handler.CreateCourse)
    http.HandleFunc("/courses/availability", handler.CalculateAvailability)
    http.HandleFunc("/courses/enrollment", handler.CreateEnrollment)

    log.Fatal(http.ListenAndServe(":8080", nil))
}