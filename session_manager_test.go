package sessionmanager_test

import (
	"errors"
	"testing"
	"time"

	sessionmanager "github.com/solrac97gr/session-manager"
	"github.com/stretchr/testify/assert"
)

func TestSessionManager_NewSessionManager(t *testing.T) {
	cases := map[string]struct {
	}{
		"empty": {},
	}

	for name := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()

			if sessionManager.Sessions == nil {
				t.Error("Sessions is nil")
			}
		})
	}
}

func TestSessionManager_GetSession(t *testing.T) {
	cases := map[string]struct {
		sessions     func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		id           string
		injected     sessionmanager.ISession
		expected     *sessionmanager.Session
		err          error
		avoidExpired bool
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			id:       "id",
			injected: nil,
			err:      errors.New("Session ID id not found"),
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:       "id",
			injected: sessionmanager.NewSession(nil),
			err:      nil,
		},

		"with data and expired": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				session.(*sessionmanager.Session).SetExpirationTime(time.Now().Add(-1 * time.Second))
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:           "id",
			injected:     sessionmanager.NewSession(nil),
			err:          errors.New("Session ID id is expired"),
			expected:     nil,
			avoidExpired: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.SetAvoidExpired(tc.avoidExpired)

			sessionManager.Sessions = tc.sessions(tc.id, tc.injected)

			session, err := sessionManager.GetSession(tc.id)

			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {
				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %s, Actual error: %s", tc.err, err)
			}

			if session != tc.injected && tc.expected != nil {
				t.Errorf("Injected session: %v, Actual session: %v", tc.injected, session)
			}

			if tc.expected != nil {
				assert.Equal(t, tc.expected, session)
			}
		})
	}

}

func TestSessionManager_CreateSession(t *testing.T) {
	cases := map[string]struct {
		sessions                   func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		err                        error
		withAlreadyExistingSession bool
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			err: nil,
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			err: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()

			sessionManager.Sessions = tc.sessions("id", sessionmanager.NewSession(nil))

			session, err := sessionManager.CreateSession()

			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {
				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %s, Actual error: %s", tc.err, err)
			}

			if session == nil {
				t.Error("Session is nil")
			}

			if session != nil && session.SessionId() == "" {
				t.Error("Session ID is empty")
			}

			if session != nil && session.SessionId() != "" && sessionManager.Sessions[session.SessionId()] == nil {
				t.Error("Session not found in sessions")
			}

		})
	}
}

func TestSessionManager_DestroySession(t *testing.T) {
	cases := map[string]struct {
		sessions func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		id       string
		err      error
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			id:  "id",
			err: errors.New("Session ID id not found"),
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:  "id",
			err: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.Sessions = tc.sessions(tc.id, sessionmanager.NewSession(nil))

			err := sessionManager.DestroySession(tc.id)

			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {
				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %s, Actual error: %s", tc.err, err)
			}

			if tc.err == nil && sessionManager.Sessions[tc.id] != nil {
				t.Errorf("Session %s not deleted", tc.id)
			}
		})
	}
}

func TestSessionManager_DestroyAllSessions(t *testing.T) {
	cases := map[string]struct {
		sessions func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		err      error
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			err: nil,
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					"id1": session,
					"id2": session,
				}
			},
			err: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.Sessions = tc.sessions("id", sessionmanager.NewSession(nil))

			err := sessionManager.DestroyAllSessions()

			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {
				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %s, Actual error: %s", tc.err, err)
			}

			if tc.err == nil && len(sessionManager.Sessions) != 0 {
				t.Errorf("Sessions not deleted")
			}
		})
	}
}

func TestSessionManager_GetAllSessions(t *testing.T) {
	cases := map[string]struct {
		sessions       func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		withNilSession bool
		err            error
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			err: nil,
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					"id1": session,
					"id2": session,
				}
			},
			err: nil,
		},

		"nil session": {
			sessions:       func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession { return nil },
			withNilSession: true,
			err:            nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()

			if tc.withNilSession {
				sessionManager.Sessions = nil
				t.Log(sessionManager.Sessions)
				sessions := sessionManager.GetAllSessions()

				if len(sessions) != 0 {
					t.Errorf("Expected 0 sessions, Actual: %d", len(sessions))
				}

				return

			}

			sessionManager.Sessions = tc.sessions("id", sessionmanager.NewSession(nil))

			sessions := sessionManager.GetAllSessions()

			if tc.err == nil && len(sessions) != len(sessionManager.Sessions) {
				t.Errorf("Sessions not found")
			}
		})
	}
}

func TestSessionManager_SetAsDefaultSession(t *testing.T) {
	cases := map[string]struct {
		sessions     func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		id           func(session sessionmanager.ISession) string
		err          error
		avoidExpired bool
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			id:  func(session sessionmanager.ISession) string { return session.SessionId() },
			err: errors.New("not found"),
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:  func(session sessionmanager.ISession) string { return session.SessionId() },
			err: nil,
		},
		"with data and expired session": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				session.(*sessionmanager.Session).SetExpirationTime(time.Now().Add(-10 * time.Minute))
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:           func(session sessionmanager.ISession) string { return session.SessionId() },
			err:          errors.New("is expired"),
			avoidExpired: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.SetAvoidExpired(tc.avoidExpired)

			s, err := sessionManager.CreateSession()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			sessionManager.Sessions = tc.sessions(s.SessionId(), s)

			err = sessionManager.SetAsDefaultSession(tc.id(s))
			t.Log(err)

			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {

				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				assert.Contains(t, err.Error(), tc.err.Error())
			}

			if tc.err == nil && sessionManager.DefaultSession != s {
				t.Errorf("Default session not set")
			}
		})
	}
}

func TestSessionManager_GetDefaultSession(t *testing.T) {
	cases := map[string]struct {
		sessions     func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		id           func(s sessionmanager.ISession) string
		err          error
		avoidExpired bool
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			id:  func(s sessionmanager.ISession) string { return s.SessionId() },
			err: errors.New("default session not set"),
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:  func(s sessionmanager.ISession) string { return s.SessionId() },
			err: nil,
		},

		"with data and expired session": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				session.(*sessionmanager.Session).SetExpirationTime(time.Now().Add(-10 * time.Minute))

				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id: func(s sessionmanager.ISession) string {
				return s.SessionId()
			},
			err:          errors.New("default session is expired"),
			avoidExpired: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			s, err := sessionManager.CreateSession()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			sessionManager.Sessions = tc.sessions(s.SessionId(), s)

			sessionManager.SetAsDefaultSession(s.SessionId())

			sessionManager.SetAvoidExpired(tc.avoidExpired)

			session, err := sessionManager.GetDefaultSession()
			if err != nil && tc.err == nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if err == nil && tc.err != nil {
				t.Errorf("Expected error: %s", tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				assert.Contains(t, err.Error(), tc.err.Error())
			}

			if tc.err == nil && session != s {
				t.Errorf("Default session not set")
			}

			if tc.err == nil && sessionManager.DefaultSession != s {
				t.Errorf("Default session not set")
			}
		})
	}

}

func TestSessionManager_SetAvoidExpired(t *testing.T) {
	cases := map[string]struct {
		avoidExpired bool
	}{
		"true": {
			avoidExpired: true,
		},
		"false": {
			avoidExpired: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.SetAvoidExpired(tc.avoidExpired)

			if sessionManager.AvoidExpired != tc.avoidExpired {
				t.Errorf("Expected AvoidExpired: %t, Actual: %t", tc.avoidExpired, sessionManager.AvoidExpired)
			}
		})
	}
}
