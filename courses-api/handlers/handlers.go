package handlers

import (
    "courses-api/services"
    "courses-api/dto/courses"
    "utils/errors"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Handler struct {
    service services.CourseServiceInterface
}

func NewHandler(service services.CourseServiceInterface) *Handler {
    return &Handler{service: service}
}

func (h *Handler) GetCourses(c *gin.Context) {
    courses, err := h.service.GetCourses()
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCourseByID(c *gin.Context) {
    id := c.Param("id")
    course, err := h.service.GetCourseById(id)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, course)
}

func (h *Handler) CreateCourse(c *gin.Context) {
    var courseRequest courses.CourseNewRequest
    if err := c.ShouldBindJSON(&courseRequest); err != nil {
        c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
        return
    }

    id, err := h.service.CreateCourse(courseRequest)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) UpdateCourse(c *gin.Context) {
    id := c.Param("id")
    var courseRequest courses.CourseNewRequest
    if err := c.ShouldBindJSON(&courseRequest); err != nil {
        c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
        return
    }

    err := h.service.UpdateCourse(id, courseRequest)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "course updated successfully"})
}

func (h *Handler) DeleteCourse(c *gin.Context) {
    id := c.Param("id")
    err := h.service.DeleteCourse(id)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "course deleted successfully"})
}

func (h *Handler) SearchCoursesByTitle(c *gin.Context) {
    title := c.Query("title")
    courses, err := h.service.SearchCoursesByTitle(title)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, courses)
}

func (h *Handler) CreateEnrollment(c *gin.Context) {
    var enrollmentRequest struct {
        CourseID string `json:"course_id"`
        UserID   string `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&enrollmentRequest); err != nil {
        c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
        return
    }

    err := h.service.CreateEnrollment(enrollmentRequest.CourseID, enrollmentRequest.UserID)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "enrollment created successfully"})
}

func (h *Handler) CalculateAvailability(c *gin.Context) {
    err := h.service.CalculateAvailability()
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "availability calculated successfully"})
}