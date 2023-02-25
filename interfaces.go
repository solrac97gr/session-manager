package sessionmanager

// ISessionManager is the interface for session manager
type ISessionManager interface {
	// Get a session by session id
	GetSession(sessionId string) (ISession, error)
	// Create a new session
	CreateSession() (ISession, error)
	// Destroy a session
	DestroySession(sessionId string) error
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
}
