package handlers

import (
    "context"
    "courses-api/services"
    "courses-api/dto/courses"
    "courses-api/utils/errors"
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
        c.JSON(http.StatusBadRequest, errors.NewBadRequestApiError("invalid json body"))
        return
    }

    id, err := h.service.Create(ctx, courseRequest)
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
        c.JSON(http.StatusBadRequest, errors.NewBadRequestApiError("invalid json body"))
        return
    }

    err := h.service.Update(id, courseRequest)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "course updated successfully"})
}

func (h *Handler) DeleteCourse(c *gin.Context) {
    id := c.Param("id")
    err := h.service.Delete(ctx, id)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "course deleted successfully"})
}

func (h *Handler) SearchByTitle(c *gin.Context) {
    title := c.Query("title")
    courses, err := h.service.SearchByTitle(ctx, title)
    if err != nil {
        c.JSON(err.Status(), err)
        return
    }
    c.JSON(http.StatusOK, courses)
}
