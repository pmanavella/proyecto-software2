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
	GetCourseByID(ctx context.Context, id string) (dto.CourseResponse_Full, errors.ApiError)
	Create(ctx context.Context, course dto.CourseResponse_Full) (dto.CourseResponse_Full, errors.ApiError)
	Update(ctx context.Context, id string, course dto.CourseResponse_Full) errors.ApiError
	Delete(ctx context.Context, id string) errors.ApiError
	SearchByTitle(ctx context.Context, title string) ([]dto.CourseResponse_Full, errors.ApiError)
	SearchByCategory(ctx context.Context, category string) ([]dto.CourseResponse_Full, errors.ApiError)
	SearchByDescription(ctx context.Context, description string) ([]dto.CourseResponse_Full, errors.ApiError)
    //RegisterUserToCourse(ctx context.Context, token string, courseID string) (dto.CourseResponse_Registration, errors.ApiError)	
	GetAll(ctx context.Context) ([]dto.CourseResponse_Full, errors.ApiError)
}

// type Queue interface {
// 	Publish(CourseNew course.CourseNew) error
// }

/// new
var (
	CourseService CourseServiceInterface
)

type courseService struct {
	repo   repositories.Mongo
	rabbit queues.Rabbit
}

func NewCourseService(repo repositories.Mongo, rabbit queues.Rabbit) CourseServiceInterface {
	return &courseService{repo: repo, rabbit: rabbit}
}

func (s *courseService) GetAll(ctx context.Context) ([]dto.CourseResponse_Full, errors.ApiError) {
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

func (s *courseService) GetCourseByID(ctx context.Context, id string) (dto.CourseResponse_Full, errors.ApiError) {
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


func (s *courseService) Create(ctx context.Context, course dto.CourseResponse_Full) (dto.CourseResponse_Full, errors.ApiError) {
    // Implementación del método Create
    daoCourse := repositories.Course{
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
        return dto.CourseResponse_Full{}, errors.NewInternalServerApiError("error creating course", err)
    }

    course.ID_Course = id
    return course, nil
}

// func (s *courseService) Create(ctx context.Context, course dto.CourseResponse_Full) (string, errors.ApiError) {
// 	daoCourse := dao.Course{
// 		Title:        course.Title,
// 		Description:  course.Description,
// 		Category:     course.Category,
// 		ImageURL:     course.ImageURL,
// 		Duration:     course.Duration,
// 		Instructor:   course.Instructor,
// 		Points:       course.Points,
// 		Capacity:     course.Capacity,
// 		Requirements: course.Requirements,
// 	}

// 	id, err := s.repo.Create(ctx, daoCourse)
// 	if err != nil {
// 		return "", errors.NewInternalServerApiError("error creating course", err)
// 	}

// 	CourseResponse_Full := dto.CourseResponse_Full{
// 		ID_Course:    id,
// 		Title:        daoCourse.Title,
// 		Description:  daoCourse.Description,
// 		Category:     daoCourse.Category,
// 		ImageURL:     daoCourse.ImageURL,
// 		Duration:     daoCourse.Duration,
// 		Instructor:   daoCourse.Instructor,
// 		Points:       daoCourse.Points,
// 		Capacity:     daoCourse.Capacity,
// 		Requirements: daoCourse.Requirements,
// 	}

// 	s.rabbit.Notify(CourseResponse_Full)
// 	return daoCourse.ID_Course, nil
// }

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

/// segunda opcion sobre DELETE

// func (s *courseService) Delete(ctx context.Context, id string) errors.ApiError {
//     // Captura ambos valores devueltos por s.repo.Delete
//     result, err := s.repo.Delete(ctx, id)
//     if err != nil {
//         return errors.NewInternalServerApiError("error deleting course", err)
//     }

//     // Puedes manejar el resultado si es necesario, por ejemplo, verificar si se eliminó algún documento
//     if result.DeletedCount == 0 {
//         return errors.NewNotFoundApiError("course not found")
//     }

//     return nil
// }

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

func (s *courseService) SearchByCategory(ctx context.Context, category string) ([]dto.CourseResponse_Full, errors.ApiError) {
	courses, err := s.repo.SearchByCategory(ctx, category)
	if err != nil {
		return nil, errors.NewInternalServerApiError("error searching courses by category", err)
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


func (s *courseService) SearchByDescription(ctx context.Context, description string) ([]dto.CourseResponse_Full, errors.ApiError) {
	courses, err := s.repo.SearchByDescription(ctx, description)
	if err != nil {
		return nil, errors.NewInternalServerApiError("error searching courses by description", err)
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

// func (s *courseService) RegisterUserToCourse(ctx context.Context, token string, courseID string) (dto.CourseResponse_Registration, errors.ApiError) {
// 	// Validate token
// 	user, err := s.repo.GetUserByToken(ctx, token)
// 	if err != nil {
// 		return dto.CourseResponse_Registration{}, errors.NewUnauthorizedApiError("invalid token")
// 	}

// 	// Validate course
// 	course, err := s.repo.GetCourseByID(ctx, courseID)
// 	if err != nil {
// 		return dto.CourseResponse_Registration{}, errors.NewNotFoundApiError("course not found")
// 	}
	
// 	// Validate user registration
// 	registration, err := s.repo.GetRegistration(ctx, user.ID_User, course.ID_Course)
// 	if err == nil {
// 		return dto.CourseResponse_Registration{}, errors.NewBadRequestApiError("user already registered to course")
// 	}

// 	// Register user to course
// 	err = s.repo.RegisterUserToCourse(ctx, user.ID_User, course.ID_Course)
// 	if err != nil {
// 		return dto.CourseResponse_Registration{}, errors.NewInternalServerApiError("error registering user to course", err)
// 	}

// 	response := dto.CourseResponse_Registration{
// 		ID_User:    user.ID_User,
// 		ID_Course:  course.ID_Course,
// 	}

// 	return response, nil
// }
