package sessionmanager

import (
	"fmt"
	"sync"
)

// SessionManager is the struct implementation for session manager
type SessionManager struct {
	DefaultSession ISession
	Sessions       map[string]ISession
	m              *sync.RWMutex
}

// Verify that SessionManager implements ISessionManager
var _ ISessionManager = (*SessionManager)(nil)

// NewSessionManager is the constructor for session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		Sessions: make(map[string]ISession),
		m:        &sync.RWMutex{},
	}
}

// Get a session by session id
func (sm *SessionManager) GetSession(sessionId string) (ISession, error) {
	sm.m.RLock()
	defer sm.m.RUnlock()
	if session, ok := sm.Sessions[sessionId]; ok {
		return session, nil
	}
	return nil, fmt.Errorf("Session ID %s not found", sessionId)
}

// Create a new session
func (sm *SessionManager) CreateSession() (ISession, error) {
	sm.m.Lock()
	defer sm.m.Unlock()
	session := NewSession(nil)
	sm.Sessions[session.SessionId()] = session
	return sm.Sessions[session.SessionId()], nil
}

// Destroy a session
func (sm *SessionManager) DestroySession(sessionId string) error {
	sm.m.Lock()
	defer sm.m.Unlock()
	if _, ok := sm.Sessions[sessionId]; !ok {
		return fmt.Errorf("Session ID %s not found", sessionId)
	}
	delete(sm.Sessions, sessionId)
	return nil
}

// SetDefaultSession sets the default session for not require session id for get a current session
func (sm *SessionManager) SetAsDefaultSession(sessionId string) error {
	sm.m.Lock()
	defer sm.m.Unlock()
	if session, ok := sm.Sessions[sessionId]; ok {
		sm.DefaultSession = session
		return nil
	}
	return fmt.Errorf("Session ID %s not found", sessionId)
}

// GetDefaultSession gets the default session for not require session id
func (sm *SessionManager) GetDefaultSession() (ISession, error) {
	sm.m.RLock()
	defer sm.m.RUnlock()
	if sm.DefaultSession == nil {
		return nil, fmt.Errorf("default session not set")
	}
	return sm.DefaultSession, nil
}

// GetAllSessions gets all sessions stored in session manager
func (sm *SessionManager) GetAllSessions() map[string]ISession {
	sm.m.RLock()
	defer sm.m.RUnlock()
	if sm.Sessions == nil {
		sm.Sessions = make(map[string]ISession)
	}
	return sm.Sessions
}

// DestroyAllSessions destroys all sessions stored in session manager
//   - Important: this method never fails, but in future it can be changed
func (sm *SessionManager) DestroyAllSessions() error {
	sm.m.Lock()
	defer sm.m.Unlock()
	sm.Sessions = make(map[string]ISession)
	return nil
}
