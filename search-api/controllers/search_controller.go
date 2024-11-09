package controllers

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "search-api/dto/search"
    "strconv"
)

// Service define la interfaz para el servicio de búsqueda de cursos
type Service interface {
    Search(ctx context.Context, query string, offset int, limit int) ([]search_dto.SearchDto, error)
}

// CourseController maneja las solicitudes HTTP relacionadas con la búsqueda de cursos
type CourseController struct {
    service Service
}

// NewCourseController crea una nueva instancia de CourseController
func NewCourseController(service Service) CourseController {
    return CourseController{
        service: service,
    }
}

// Search maneja la búsqueda de cursos
func (controller CourseController) Search(c *gin.Context) {
    // Leer el parámetro de búsqueda "query" desde la URL
    query := c.Query("query")

    // Leer y parsear el parámetro "offset" desde la URL
    offset, err := strconv.Atoi(c.Query("offset"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("invalid offset value: %s", err),
        })
        return
    }

    // Leer y parsear el parámetro "limit" desde la URL
    limit, err := strconv.Atoi(c.Query("limit"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("invalid limit value: %s", err),
        })
        return
    }

    // Invocar el servicio para buscar cursos
    courses, err := controller.service.Search(c.Request.Context(), query, offset, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("error searching courses: %s", err.Error()),
        })
        return
    }

    // Enviar la respuesta con los resultados
    c.JSON(http.StatusOK, courses)
}
