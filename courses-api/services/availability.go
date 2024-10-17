package services

import "sync"

type Availability struct {
	CourseID   string
	Available bool
}

func GetAvailability(courseIDs []string) map[string]bool {
	result := make(map[string]bool)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(courseIDs))
	ch := make(chan Availability)
	go func() {
		for {
			availability := <-ch
			result[availability.CourseID] = availability.Available
		}
	}()
	for _, courseID := range courseIDs {
		go IsAvailableAsync(courseID, &waitGroup, ch)
	}
	waitGroup.Wait()
	return result
}

func IsAvailableAsync(courseID string, group *sync.WaitGroup, ch chan Availability) {
	defer group.Done()
	ch <- Availability{
		CourseID:   courseID,
		Available: IsAvailable(courseID),
	}
}

func IsAvailable(hotelID string) bool {
	// implement
	return true
}