package cursos

import (
	"context"
	courseDto "courses-api/dto/courses"
	service "courses-api/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	//dto "courses-api/dto/courses"
)

type Service interface {
	GetCourseByID(ctx context.Context, id string) (courseDto.CourseResponse_Full, error)
	Create(ctx context.Context, curso courseDto.CourseResponse_Full) (string, error)
	Update(ctx context.Context, curso courseDto.CourseResponse_Full) error
	Delete(courseID int) error
	GetCourses() (courseDto.CoursesResponse_Full, error)
	SearchByTitle(title string) (courseDto.CoursesResponse_Full, error)
	SearchByCategory(category string) (courseDto.CoursesResponse_Full, error)
	SearchByDescription(description string) (courseDto.CoursesResponse_Full, error)
	RegisterUserToCourse(token string, courseID int) (courseDto.CourseResponse_Registration, error)
	GetAll() (courseDto.CoursesResponse_Full, error)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetCourseByID(ctx *gin.Context) {
	// Validate ID param
	cursoID := strings.TrimSpace(ctx.Param("id"))
	if cursoID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "course ID is required",
		})
		return
	}

	// Get course by ID using the service
	curso, err := controller.service.GetCourseByID(ctx.Request.Context(), cursoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error getting course: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, curso)
}

func (controller Controller) Create(ctx *gin.Context) {
	var curso courseDto.CourseResponse_Full
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request body: %s", err.Error()),
		})
		return
	}

	// Create course using the service
	cursoID, err := controller.service.Create(ctx.Request.Context(), curso)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating course: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusCreated, gin.H{
		"id": cursoID,
	})
}

func (controller Controller) Update(ctx *gin.Context) {
	var curso courseDto.CourseResponse_Full
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request body: %s", err.Error()),
		})
		return
	}

	// Update course using the service
	if err := controller.service.Update(ctx.Request.Context(), curso); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error updating course: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "course updated successfully",
	})
}

func GetCourses(c *gin.Context) {

	var coursesDto courseDto.CoursesResponse_Full
	coursesDto, err := service.courseService.GetCourses()
	fmt.Println(coursesDto, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCourseByTitle(c *gin.Context) {

	//coursesDto, err := courseService.GetCourseByIdUser(id) ---> ver si va o no

	var title string
	title = c.Param(title)
	var CoursesResponse_Full courseDto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchByTitle(title)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func GetCourseByCategory(c *gin.Context) {

	var category string
	category = c.Param(category)
	var CoursesResponse_Full courseDto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchByCategory(category)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func GetCourseByDescription(c *gin.Context) {

	var description string
	description = c.Param(description)
	var CoursesResponse_Full courseDto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchByDescription(description)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func PostCourse(c *gin.Context) {
	var courseDto courseDto.CourseResponse_Full
	err := c.BindJSON(&courseDto)

	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.CreateCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)
}

func PutCourse(c *gin.Context) {
	var courseDto courseDto.CourseResponse_Full

	err := c.BindJSON(&courseDto)

	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.UpdateCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)

}

func DeleteCourse(c *gin.Context) {
	idParam := c.Param("id")
	courseID, err := strconv.Atoi(idParam)
	if err != nil {
		//log.Error("Invalid course ID: " + idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = service.CourseService.DeleteCourse(courseID)
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func RegisterUserToCourse(c *gin.Context) {
	var crr courseDto.CourseRequest_Registration
	_ = c.BindJSON(&crr)

	var CourseResponseDto courseDto.CourseResponse_Registration
	CourseResponseDto, err := service.CourseService.RegisterUserToCourse(crr.Token, crr.ID_Course)
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CourseResponseDto)
}

func GetAll(c *gin.Context) {
	var coursesDto courseDto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetAllCourses()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

// func GetCoursesByUser(c *gin.Context) {
// 	var tokenDto dto.CourseRequest_Token
// 	_ = c.BindJSON(&tokenDto)

// 	var coursesDto dto.CoursesResponse_Full
// 	coursesDto, err := service.CourseService.GetCoursesByUser(tokenDto.Token)

// 	log.Debug(err)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	log.Debug(coursesDto)
// 	c.JSON(http.StatusOK, coursesDto)
// }

// func GetAvailableCoursesByUser(c *gin.Context) {
// 	var tokenDto dto.CourseRequest_Token
// 	_ = c.BindJSON(&tokenDto)

// 	var coursesDto dto.CoursesResponse_Full
// 	coursesDto, err := service.CourseService.GetAvailableCoursesByUser(tokenDto.Token)

// 	log.Debug(err)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	log.Debug(coursesDto)
// 	c.JSON(http.StatusOK, coursesDto)
// }
