package services

import (
	"cursos-ucc/dto"
	e "cursos-ucc/utils/errors"
	//"strconv"
	//registerclient "cursos-ucc/clients/register"
	//"github.com/golang-jwt/jwt"
	//"github.com/jinzhu/gorm"
)

type courseService struct {
	HTTPClient utils.HttpClientInterface
}

type CourseServiceInterface interface {
	GetCourses() (dto.CoursesResponse_Full, e.ApiError)
	GetCourseById(id int) (dto.CourseResponse_Full, e.ApiError)
	GetCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError)
	GetAvailableCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, e.ApiError)
	CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError)
	UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError)
	DeleteCourse(courseId int) e.ApiError
	RegisterUserToCourse(tokenString string, courseId int) (dto.CourseResponse_Registration, e.ApiError)
}

var CourseService courseServiceInterface
