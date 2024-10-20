/*package clients_search

import (dao_search "search-api/dao"

log "github.com/sirupsen/logrus"
)


type coursesClient struct{}


type CoursesClientInterface interface {
    GetCourseById(id string) dao_search.Search       // Buscar curso por ID
    GetCourses() []dao_search.Search                 // Obtener todos los cursos
    GetCourseByTitle(query string) []dao_search.Search // Buscar cursos por nombre
}

var (
	CoursesClient CoursesClientInterface
)

func init() {
	CoursesClient = &coursesClient{}
}

func (s *coursesClient) GetCourseById(id string) (dao_search.Search, error) {
    var course dao_search.Search
    err := clients.Db.Where("id_course = ?", id).First(&course).Error 
        log.Error("Error loading course by ID: ", err)
        return course, err 
    }
    log.Debug("Course: ", course)
    return course, nil
}



*/
