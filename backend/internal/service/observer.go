package service

import (
	"log"
	"sync"
)

type PostEvent struct {
	EventType string      `json:"event_type"`
	PostID    int64       `json:"post_id"`
	Data      interface{} `json:"data"`
}

type Observer interface {
	Update(event PostEvent) error
}

type PostService struct {
	observers []Observer
	mu        sync.Mutex
}

func NewPostService() *PostService {
	return &PostService{
		observers: make([]Observer, 0),
	}
}

func (s *PostService) Subscribe(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, observer)
}

func (s *PostService) Unsubscribe(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *PostService) Notify(event PostEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, observer := range s.observers {
		go func(obs Observer) {
			if err := obs.Update(event); err != nil {
				log.Printf("Error notifying observer: %v", err)
			}
		}(observer)
	}
}

// Example Observers
type SearchIndexObserver struct{}

func (o *SearchIndexObserver) Update(event PostEvent) error {
	// In a real implementation, this would update the search index
	log.Printf("Updating search index for post %d: %v", event.PostID, event.EventType)
	return nil
}

type NotificationObserver struct{}

func (o *NotificationObserver) Update(event PostEvent) error {
	// In a real implementation, this would send notifications
	log.Printf("Sending notification for post %d: %v", event.PostID, event.EventType)
	return nil
}

type AnalyticsObserver struct{}

func (o *AnalyticsObserver) Update(event PostEvent) error {
	// In a real implementation, this would track analytics
	log.Printf("Tracking analytics for post %d: %v", event.PostID, event.EventType)
	return nil
}
