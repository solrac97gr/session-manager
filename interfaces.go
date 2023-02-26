package sessionmanager

import "time"

// ISessionManager is the interface for session manager
type ISessionManager interface {
	// Get a session by session id
	GetSession(sessionId string) (ISession, error)
	// Create a new session
	CreateSession() (ISession, error)
	// Destroy a session
	DestroySession(sessionId string) error
	// SetDefaultSession sets the default session
	SetAsDefaultSession(sessionId string) error
	// GetDefaultSession gets the default session
	GetDefaultSession() (ISession, error)
	// GetAllSessions gets all sessions
	GetAllSessions() map[string]ISession
	// DestroyAllSessions destroys all sessions
	DestroyAllSessions() error
	// SetAvoidExpired sets if session manager avoid expired sessions
	SetAvoidExpired(avoidExpired bool)
}

// Session is the interface for session
type ISession interface {
	// Get a value from session
	Get(key string) (interface{}, error)
	// Set a value to session
	Set(key string, value interface{}) error
	// Delete a value from session
	Delete(key string) error
	// Get session id
	SessionId() string
	// Set ExpirationTime
	SetExpirationTime(expirationTime time.Time)
	// IsExpired checks if session is expired
	IsExpired() bool
	// IsActive checks if session is active
	IsActive() bool
}
