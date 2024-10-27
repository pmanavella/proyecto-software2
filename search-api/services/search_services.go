package search


import (
	"context"
	"fmt"
	"log"
	"search-api/search_dao"
	"search-api/search_dto"
)

// Repository define la interfaz para la búsqueda de cursos
type Repository interface {
	Search(ctx context.Context, query string, offset int, limit int) ([]search_dao.Search, error)
}

// RabbitMQ2 define la interfaz para consumir mensajes de RabbitMQ
type RabbitMQ2 interface {
	ConsumeQueue()
}

// Service representa el servicio principal de búsqueda de cursos
type Service struct {
	repository Repository
	rabbitRepo RabbitMQ2
}

// NewService crea una nueva instancia del servicio de búsqueda
func NewService(repository Repository, rabbitRepo RabbitMQ2) Service {
	return Service{
		repository: repository,
		rabbitRepo: rabbitRepo,
	}
}

// Search realiza una búsqueda de cursos basados en un término de consulta, offset y limit
func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]search_dto.SearchDto, error) {
	// Prueba de conexión para verificar si llegan mensajes
	log.Println("Conexión correcta")
	service.rabbitRepo.ConsumeQueue()

	// Realiza la búsqueda de cursos en el repositorio
	courses, err := service.repository.Search(ctx, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("error buscando cursosDAO: %s", err.Error())
	}

	// Convertir los resultados obtenidos en una lista de SearchDto
	result := make([]search_dto.SearchDto, 0)
	for _, course := range courses {
		result = append(result, search_dto.SearchDto{
			ID_Course:    course.ID_Course,
			Title:        course.Description, // Cambiar según corresponda
			Description:  course.Description,
			Category:     course.Category,
			ImageURL:     course.ImageURL,
			Duration:     course.Duration,
			Requirements: course.Requirements,
			Instructor:   course.Instructor,
		})
	}

	return result, nil
}
