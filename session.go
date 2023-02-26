package sessionmanager

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Session is the struct implementation for session
// ID is the unique id for session
// Data is the data for session
type Session struct {
	ID             string
	Data           map[string]interface{}
	m              *sync.RWMutex
	ExpirationTime time.Time
	Expired        bool
	Active         bool
}

// Verify that Session implements ISession
var _ ISession = (*Session)(nil)

// NewSession is the constructor for session by default expiration time is 30 minutes
// and the session is active you can edit this values by setting the ExpirationTime and Active fields
func NewSession(data map[string]interface{}) *Session {
	sessionId := uuid.New().String()

	if data == nil {
		data = make(map[string]interface{})
	}

	return &Session{
		ID:             sessionId,
		Data:           data,
		m:              &sync.RWMutex{},
		Active:         true,
		ExpirationTime: time.Now().Add(time.Minute * 5),
		Expired:        false,
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

// SetExpirationTime sets the expiration time for session in case you
// want to change the default expiration time
func (s *Session) SetExpirationTime(expirationTime time.Time) {
	s.m.Lock()
	s.ExpirationTime = expirationTime
	s.m.Unlock()
}

// IsExpired returns true if the session is expired
func (s *Session) IsExpired() bool {
	s.m.RLock()
	if s.Expired {
		return true
	}
	s.m.RUnlock()

	if s.Active && time.Now().After(s.ExpirationTime) {
		s.m.Lock()
		s.Expired = true
		s.Active = false
		s.m.Unlock()
	}

	return s.Expired
}

// IsActive returns true if the session is active
func (s *Session) IsActive() bool {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.Active
}
