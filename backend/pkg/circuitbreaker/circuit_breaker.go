package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	name            string
	maxRequests     int32
	resetTimeout    time.Duration
	state           CircuitState
	mutex           sync.Mutex
	failureCount    int32
	lastFailureTime time.Time
}

func NewCircuitBreaker(name string, maxRequests int32, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		name:         name,
		maxRequests:  maxRequests,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	if cb.state == StateOpen {
		return nil, errors.New("circuit breaker is open")
	}

	result, err := req()
	if err != nil {
		cb.recordFailure()
		return nil, err
	}

	cb.recordSuccess()
	return result, nil
}

func (cb *CircuitBreaker) recordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.failureCount++
	cb.lastFailureTime = time.Now()

	if cb.failureCount >= cb.maxRequests {
		cb.state = StateOpen
		go cb.startResetTimer()
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.failureCount = 0
	cb.state = StateClosed
}

func (cb *CircuitBreaker) startResetTimer() {
	time.Sleep(cb.resetTimeout)
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.state = StateHalfOpen
	cb.failureCount = 0
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	return cb.state
}

// GetName returns the name of the circuit breaker
func (cb *CircuitBreaker) GetName() string {
	return cb.name
}

// GetFailureCount returns the current failure count
func (cb *CircuitBreaker) GetFailureCount() int32 {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	return cb.failureCount
}

// Wrapper for external search service
type SearchService struct {
	circuitBreaker *CircuitBreaker
}

func NewSearchService() *SearchService {
	return &SearchService{
		circuitBreaker: NewCircuitBreaker("search", 5, 30*time.Second),
	}
}

func (s *SearchService) Search(query string) ([]string, error) {
	var result []string

	res, err := s.circuitBreaker.Execute(func() (interface{}, error) {
		// In a real implementation, this would call an external search API
		// For this example, we'll simulate a failure rate
		if time.Now().Unix()%5 == 0 { // Simulate 20% failure rate
			return nil, errors.New("search service unavailable")
		}

		// Mock search results
		searchResults := []string{
			"Result 1 for " + query,
			"Result 2 for " + query,
			"Result 3 for " + query,
		}
		return searchResults, nil
	})

	if err != nil {
		// Fallback logic when search service fails
		result = []string{
			"Fallback result 1 for " + query,
			"Fallback result 2 for " + query,
		}
		return result, nil
	}

	if res != nil {
		result = res.([]string)
	}

	return result, nil
}
