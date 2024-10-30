package handlers

import (
    "courses-api/dao"
    "courses-api/repositories"
    "courses-api/queues"
    "encoding/json"
    "net/http"
    "sync"
)

type Handler struct {
    repo   *repositories.CourseRepository
    rabbit *queues.Rabbit
}

func NewHandler(repo *repositories.CourseRepository, rabbit *queues.Rabbit) *Handler {
    return &Handler{repo: repo, rabbit: rabbit}
}

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
    var course dao.Course
    err := json.NewDecoder(r.Body).Decode(&course)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.repo.CreateCourse(course)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.rabbit.Notify(course)
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) CalculateAvailability(w http.ResponseWriter, r *http.Request) {
    // Implementar lógica para calcular disponibilidad usando Go Routines
}

func (h *Handler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
    // Implementar lógica para crear inscripción
}