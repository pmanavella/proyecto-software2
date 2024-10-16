package courses

import (
	"cursos-ucc/dto"
	service "cursos-ucc/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCourses(c *gin.Context) {

	var coursesDto dto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetCourses()
	fmt.Println(coursesDto, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCoursesByUser(c *gin.Context) {
	var tokenDto dto.CourseRequest_Token
	_ = c.BindJSON(&tokenDto)

	var coursesDto dto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetCoursesByUser(tokenDto.Token)

	log.Debug(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Debug(coursesDto)
	c.JSON(http.StatusOK, coursesDto)
}