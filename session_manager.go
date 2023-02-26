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
	AvoidExpired   bool
}

// Verify that SessionManager implements ISessionManager
var _ ISessionManager = (*SessionManager)(nil)

// NewSessionManager is the constructor for session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		Sessions:     make(map[string]ISession),
		m:            &sync.RWMutex{},
		AvoidExpired: false,
	}
}

// Get a session by session id
func (sm *SessionManager) GetSession(sessionId string) (ISession, error) {
	sm.m.RLock()
	defer sm.m.RUnlock()
	if session, ok := sm.Sessions[sessionId]; ok {
		if sm.AvoidExpired && session.IsExpired() {
			return nil, fmt.Errorf("Session ID %s is expired", sessionId)
		}
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

// SetAsDefaultSession sets the default session for not require session id for get a current session
func (sm *SessionManager) SetAsDefaultSession(sessionId string) error {
	sm.m.Lock()
	defer sm.m.Unlock()
	if session, ok := sm.Sessions[sessionId]; ok {
		if sm.AvoidExpired && session.IsExpired() {
			return fmt.Errorf("Session ID %s is expired", sessionId)
		}
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

	if sm.AvoidExpired && sm.DefaultSession.IsExpired() {
		return nil, fmt.Errorf("default session is expired")
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

// SetAvoidExpired sets the avoid expired flag
//   - If avoid expired is true, the session manager will not return expired sessions
//   - If avoid expired is false, the session manager will return expired sessions
func (sm *SessionManager) SetAvoidExpired(avoidExpired bool) {
	sm.m.Lock()
	defer sm.m.Unlock()
	sm.AvoidExpired = avoidExpired
}
