package service

// Service is the base interface that all services should implement
// It provides a common cleanup method for resource management
type Service interface {
	Close() error
}