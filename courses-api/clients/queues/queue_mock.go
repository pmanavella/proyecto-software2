package queues

import "courses-api/dto/courses"

type Mock struct{}

func NewMock() Mock {
	return Mock{}
}

func (Mock) Publish(CourseNew courses.CourseNew) error {
	return nil
}
