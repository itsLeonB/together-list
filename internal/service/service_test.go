package service

import (
	"testing"
)

// MockService for testing
type MockService struct {
	closed bool
}

func (m *MockService) Close() error {
	m.closed = true
	return nil
}

func TestServiceInterface(t *testing.T) {
	var s Service = &MockService{}
	
	if err := s.Close(); err != nil {
		t.Errorf("Close() returned error: %v", err)
	}
	
	mock := s.(*MockService)
	if !mock.closed {
		t.Error("Service was not properly closed")
	}
}