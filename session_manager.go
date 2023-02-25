package sessionmanager

import (
	"fmt"
)

// SessionManager is the struct implementation for session manager
type SessionManager struct {
	Sessions map[string]ISession
}

// Verify that SessionManager implements ISessionManager
var _ ISessionManager = (*SessionManager)(nil)

// NewSessionManager is the constructor for session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

// Get a session by session id
func (sm *SessionManager) GetSession(sessionId string) (ISession, error) {
	if session, ok := sm.Sessions[sessionId]; ok {
		return session, nil
	}
	return nil, fmt.Errorf("Session ID %s not found", sessionId)
}

// Create a new session
func (sm *SessionManager) CreateSession() (ISession, error) {
	session := NewSession(nil)
	sm.Sessions[session.SessionId()] = session
	return session, nil
}

// Destroy a session
func (sm *SessionManager) DestroySession(sessionId string) error {
	if _, ok := sm.Sessions[sessionId]; !ok {
		return fmt.Errorf("Session ID %s not found", sessionId)
	}
	delete(sm.Sessions, sessionId)
	return nil
}
