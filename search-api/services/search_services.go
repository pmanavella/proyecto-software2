package search


import (
	"context"
	"fmt"
	"log"
	dao "search-api/dao"
	"search-api/dto/search"
	http "net/http"
	bytes "bytes"
	json "encoding/json"
	
)

// Repository define la interfaz para la búsqueda de cursos
type Repository interface {
	Index(ctx context.Context, course dao.Search) (string, error)
	Update(ctx context.Context, course dao.Search) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, offset int, limit int) ([]dao.Search, error)
}

// RabbitMQ2 define la interfaz para consumir mensajes de RabbitMQ
type RabbitMQ2 interface {
	ConsumeQueue()
}

type ExternalRepository interface {
	GetCourseByID(ctx context.Context, id string) (dao.Search, error)
}

// Service representa el servicio principal de búsqueda de cursos
type Service struct {
	repository Repository
	coursesAPI ExternalRepository
}

// NewService crea una nueva instancia del servicio de búsqueda
func NewService(repository Repository, coursesAPI ExternalRepository) Service {
	return Service{
		repository: repository,
		coursesAPI: coursesAPI,
	}
}

// Search realiza una búsqueda de cursos basados en un término de consulta, offset y limit
func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]search_dto.SearchDto, error) {
	// Prueba de conexión para verificar si llegan mensajes
	log.Println("Conexión correcta")

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


func (service Service) HandleCourseNew(courseNew dao.CourseNew) {
	switch courseNew.Operation {
	case "CREATE", "UPDATE":
		// Fetch course details from the local service
		course, err := service.coursesAPI.GetCourseByID(context.Background(), courseNew.CourseID)
		if err != nil {
			fmt.Printf("Error getting course (%s) from API: %v\n", courseNew.CourseID, err)
			return
		}

		// Map to the Search DAO structure
		courseDAO := dao.Search{
			ID_Course:    course.ID_Course,
			Description:  course.Description,
			Category:     course.Category,
			ImageURL:     course.ImageURL,
			Duration:     course.Duration,
			Instructor:   course.Instructor,
			Requirements: course.Requirements,
			Capacity:     course.Capacity,
			Points:       course.Points,
		}

		// Handle Index operation
		if courseNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), courseDAO); err != nil {
				fmt.Printf("Error indexing course (%s): %v\n", courseNew.CourseID, err)
			} else {
				fmt.Println("Course indexed successfully:", courseNew.CourseID)
			}
		} else { // Handle Update operation
			if err := service.repository.Update(context.Background(), courseDAO); err != nil {
				fmt.Printf("Error updating course (%s): %v\n", courseNew.CourseID, err)
			} else {
				fmt.Println("Course updated successfully:", courseNew.CourseID)
			}
		}

	case "DELETE":
		// Call Delete method directly since no course details are needed
		if err := service.repository.Delete(context.Background(), courseNew.CourseID); err != nil {
			fmt.Printf("Error deleting course (%s): %v\n", courseNew.CourseID, err)
		} else {
			fmt.Println("Course deleted successfully:", courseNew.CourseID)
		}

	default:
		fmt.Printf("Unknown operation: %s\n", courseNew.Operation)
	}
}

type SolrRepository struct {
	solrURL string
}

// NewSolrRepository crea una nueva instancia de SolrRepository
func NewSolrRepository(solrURL string) *SolrRepository {
	return &SolrRepository{solrURL: solrURL}
}

// Index envía un documento de curso a Solr para indexación
func (repo *SolrRepository) Index(ctx context.Context, course dao.Search) (string, error) {
	doc := map[string]interface{}{
		"id_course":    course.ID_Course,
		"description":  course.Description,
		"category":     course.Category,
		"image_url":    course.ImageURL,
		"duration":     course.Duration,
		"instructor":   course.Instructor,
		"points":       course.Points,
		"requirements": course.Requirements,
		"capacity":     course.Capacity,
	}

	indexRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	jsonData, err := json.Marshal(indexRequest)
	if err != nil {
		return "", fmt.Errorf("error marshaling course data to JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", repo.solrURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Solr: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error indexing course in Solr, status code: %d", resp.StatusCode)
	}

	return course.ID_Course, nil
}

// Update modifica un documento de curso existente en Solr
func (repo *SolrRepository) Update(ctx context.Context, course dao.Search) error {
	doc := map[string]interface{}{
		"id_course":    course.ID_Course,
		"description":  course.Description,
		"category":     course.Category,
		"image_url":    course.ImageURL,
		"duration":     course.Duration,
		"instructor":   course.Instructor,
		"points":       course.Points,
		"requirements": course.Requirements,
		"capacity":     course.Capacity,
	}

	updateRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	body, err := json.Marshal(updateRequest)
	if err != nil {
		return fmt.Errorf("error marshaling course document: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", repo.solrURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update course in Solr, status code: %d", resp.StatusCode)
	}

	return nil
}
