package services

import (
    //"context"
    dao "courses-api/dao/courses"
	dto "courses-api/dto/courses"
    "courses-api/repositories"
    "utils/errors"
    "courses-api/clients/queues"
    "sync"
)

type CourseServiceInterface interface {
    GetCourses() ([]dto.CourseResponse_Full, errors.ApiError)
    GetCourseById(id string) (dto.CourseResponse_Full, errors.ApiError)
    CreateCourse(course dto.CourseNewRequest) (string, errors.ApiError)
    UpdateCourse(id string, course dto.CourseNewRequest) errors.ApiError
    DeleteCourse(id string) errors.ApiError
    SearchCoursesByTitle(title string) ([]dto.CourseResponse_Full, errors.ApiError)
    CreateEnrollment(courseID string, userID string) errors.ApiError
    CalculateAvailability() errors.ApiError
}

type courseService struct {
    repo   *repositories.CourseRepository
    rabbit *queues.Rabbit
}

func NewCourseService(repo *repositories.CourseRepository, rabbit *queues.Rabbit) CourseServiceInterface {
    return &courseService{repo: repo, rabbit: rabbit}
}

func (s *courseService) GetCourses() ([]dto.CourseResponse_Full, errors.ApiError) {
    courses, err := s.repo.GetAllCourses()
    if err != nil {
        return nil, errors.NewInternalServerError("error getting courses")
    }

    var response []dto.CourseResponse_Full
    for _, course := range courses {
        response = append(response, dto.CourseResponse_Full{
            ID_Course:    course.ID_Course,
            Title:        course.Title,
            Description:  course.Description,
            Category:     course.Category,
            ImageURL:     course.ImageURL,
            Duration:     course.Duration,
            Instructor:   course.Instructor,
            Points:       course.Points,
            Capacity:     course.Capacity,
            Requirements: course.Requirements,
        })
    }

    return response, nil
}

func (s *courseService) GetCourseById(id string) (dto.CourseResponse_Full, errors.ApiError) {
    course, err := s.repo.GetCourse(id)
    if err != nil {
        return dto.CourseResponse_Full{}, errors.NewNotFoundError("course not found")
    }

    response := dto.CourseResponse_Full{
        ID_Course:    course.ID_Course,
        Title:        course.Title,
        Description:  course.Description,
        Category:     course.Category,
        ImageURL:     course.ImageURL,
        Duration:     course.Duration,
        Instructor:   course.Instructor,
        Points:       course.Points,
        Capacity:     course.Capacity,
        Requirements: course.Requirements,
    }

    return response, nil
}

func (s *courseService) CreateCourse(course dto.CourseResponse_Full) (string, errors.ApiError) {
    daoCourse := dao.Course{
        Title:        course.Title,
        Description:  course.Description,
        Category:     course.Category,
        ImageURL:     course.ImageURL,
        Duration:     course.Duration,
        Instructor:   course.Instructor,
        Points:       course.Points,
        Capacity:     course.Capacity,
        Requirements: course.Requirements,
    }

    err := s.repo.CreateCourse(daoCourse)
    if err != nil {
        return "", errors.NewInternalServerError("error creating course")
    }

    CourseResponse_Full := dto.CourseResponse_Full{
        ID_Course:    daoCourse.ID_Course,
        Title:        daoCourse.Title,
        Description:  daoCourse.Description,
        Category:     daoCourse.Category,
        ImageURL:     daoCourse.ImageURL,
        Duration:     daoCourse.Duration,
        Instructor:   daoCourse.Instructor,
        Points:       daoCourse.Points,
        Capacity:     daoCourse.Capacity,
        Requirements: daoCourse.Requirements,
    }

    s.rabbit.Notify(CourseResponse_Full)
    return daoCourse.ID_Course, nil
}

func (s *courseService) UpdateCourse(id string, course dto.CourseNewRequest) errors.ApiError {
    daoCourse := dao.Course{
        Title:        course.Title,
        Description:  course.Description,
        Category:     course.Category,
        ImageURL:     course.ImageURL,
        Duration:     course.Duration,
        Instructor:   course.Instructor,
        Points:       course.Points,
        Capacity:     course.Capacity,
        Requirements: course.Requirements,
    }

    err := s.repo.UpdateCourse(id, daoCourse)
    if err != nil {
        return errors.NewInternalServerError("error updating course")
    }

    CourseResponse_Full := dto.CourseResponse_Full{
        ID_Course:    daoCourse.ID_Course,
        Title:        daoCourse.Title,
        Description:  daoCourse.Description,
        Category:     daoCourse.Category,
        ImageURL:     daoCourse.ImageURL,
        Duration:     daoCourse.Duration,
        Instructor:   daoCourse.Instructor,
        Points:       daoCourse.Points,
        Capacity:     daoCourse.Capacity,
        Requirements: daoCourse.Requirements,
    }

    s.rabbit.Notify(CourseResponse_Full)
    return nil
}

func (s *courseService) DeleteCourse(id string) errors.ApiError {
    err := s.repo.DeleteCourse(id)
    if err != nil {
        return errors.NewInternalServerError("error deleting course")
    }

    return nil
}

func (s *courseService) SearchCoursesByTitle(title string) ([]dto.CourseResponse_Full, errors.ApiError) {
    courses, err := s.repo.SearchCoursesByTitle(title)
    if err != nil {
        return nil, errors.NewInternalServerError("error searching courses by title")
    }

    var response []dto.CourseResponse_Full
    for _, course := range courses {
        response = append(response, dto.CourseResponse_Full{
            ID_Course:    course.ID_Course,
            Title:        course.Title,
            Description:  course.Description,
            Category:     course.Category,
            ImageURL:     course.ImageURL,
            Duration:     course.Duration,
            Instructor:   course.Instructor,
            Points:       course.Points,
            Capacity:     course.Capacity,
            Requirements: course.Requirements,
        })
    }

    return response, nil
}

func (s *courseService) CreateEnrollment(courseID string, userID string) errors.ApiError {
    // Implementar lógica para crear inscripción
    return nil
}

func (s *courseService) CalculateAvailability() errors.ApiError {
    // Implementar lógica para calcular disponibilidad usando Go Routines
    var wg sync.WaitGroup
    courses, err := s.repo.GetAllCourses()
    if err != nil {
        return errors.NewInternalServerError("error getting courses")
    }

    for _, course := range courses {
        wg.Add(1)
        go func(course dao.Course) {
            defer wg.Done()
            // Lógica para calcular disponibilidad
        }(course)
    }

    wg.Wait()
    return nil
}