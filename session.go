package sessionmanager

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// Session is the struct implementation for session
// ID is the unique id for session
// Data is the data for session
type Session struct {
	ID   string
	Data map[string]interface{}
	m    *sync.RWMutex
}

// Verify that Session implements ISession
var _ ISession = (*Session)(nil)

// NewSession is the constructor for session
func NewSession(data map[string]interface{}) *Session {
	sessionId := uuid.New().String()

	if data == nil {
		data = make(map[string]interface{})
	}

	return &Session{
		ID:   sessionId,
		Data: data,
		m:    &sync.RWMutex{},
	}
}

// Get a value from session
func (s *Session) Get(key string) (interface{}, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	if _, ok := s.Data[key]; !ok {
		return nil, errors.New("key not found")
	}
	return s.Data[key], nil
}

// Set a value to session
func (s *Session) Set(key string, value interface{}) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Data[key]; ok {
		return fmt.Errorf("key %s already exists, for replace delete it first", key)
	}
	s.Data[key] = value
	return nil
}

// Delete a value from session
func (s *Session) Delete(key string) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Data[key]; !ok {
		return errors.New("key not found")
	}
	delete(s.Data, key)
	return nil
}

// SessionId returns the session id
func (s *Session) SessionId() string {
	return s.ID
}
