package services

import (
	"context"
	"courses-api/clients/queues"
	dao "courses-api/dao/courses"
	dto "courses-api/dto/courses"
	"courses-api/repositories"
	errors "courses-api/utils/errors"
	//"sync"
)

type CourseServiceInterface interface {
	GetCourses(ctx context.Context) ([]dto.CourseResponse_Full, errors.ApiError)
	GetById(ctx context.Context, id string) (dto.CourseResponse_Full, errors.ApiError)
	Create(ctx context.Context, course dto.CourseResponse_Full) (string, errors.ApiError)
	Update(ctx context.Context, id string, course dto.CourseResponse_Full) errors.ApiError
	Delete(ctx context.Context, id string) errors.ApiError
	SearchByTitle(ctx context.Context, title string) ([]dto.CourseResponse_Full, errors.ApiError)
}

type courseService struct {
	repo   repositories.Mongo
	rabbit queues.Rabbit
}

func NewCourseService(repo repositories.Mongo, rabbit queues.Rabbit) CourseServiceInterface {
	return &courseService{repo: repo, rabbit: rabbit}
}

func (s *courseService) GetCourses(ctx context.Context) ([]dto.CourseResponse_Full, errors.ApiError) {
	courses, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.NewInternalServerApiError("error getting courses", err)
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

func (s *courseService) GetById(ctx context.Context, id string) (dto.CourseResponse_Full, errors.ApiError) {
	course, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return dto.CourseResponse_Full{}, errors.NewNotFoundApiError("course not found")
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

func (s *courseService) Create(ctx context.Context, course dto.CourseResponse_Full) (string, errors.ApiError) {
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

	id, err := s.repo.Create(ctx, daoCourse)
	if err != nil {
		return "", errors.NewInternalServerApiError("error creating course", err)
	}

	CourseResponse_Full := dto.CourseResponse_Full{
		ID_Course:    id,
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

func (s *courseService) Update(ctx context.Context, id string, course dto.CourseResponse_Full) errors.ApiError {
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

	err := s.repo.Update(ctx, daoCourse)
	if err != nil {
		return errors.NewInternalServerApiError("error updating course", err)
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

func (s *courseService) Delete(ctx context.Context, id string) errors.ApiError {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return errors.NewInternalServerApiError("error deleting course", err)
	}

	return nil
}

func (s *courseService) SearchByTitle(ctx context.Context, title string) ([]dto.CourseResponse_Full, errors.ApiError) {
	courses, err := s.repo.SearchByTitle(ctx, title)
	if err != nil {
		return nil, errors.NewInternalServerApiError("error searching courses by title", err)
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

