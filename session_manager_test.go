package sessionmanager_test

import (
	"errors"
	"testing"

	sessionmanager "github.com/solrac97gr/session-manager"
)

func TestSessionManager_NewSessionManager(t *testing.T) {
	cases := map[string]struct {
	}{
		"empty": {},
	}

	for name := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()

			if sessionManager.Sessions != nil {
				t.Error("Sessions is not nil")
			}
		})
	}
}

func TestSessionManager_GetSession(t *testing.T) {
	cases := map[string]struct {
		sessions func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession
		id       string
		expected sessionmanager.ISession
		err      error
	}{
		"empty": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{}
			},
			id:       "id",
			expected: nil,
			err:      errors.New("Session ID id not found"),
		},

		"with data": {
			sessions: func(id string, session sessionmanager.ISession) map[string]sessionmanager.ISession {
				return map[string]sessionmanager.ISession{
					id: session,
				}
			},
			id:       "id",
			expected: sessionmanager.NewSession(nil),
			err:      nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			sessionManager.Sessions = tc.sessions(tc.id, tc.expected)

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

			if session != tc.expected {
				t.Errorf("Expected session: %v, Actual session: %v", tc.expected, session)
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
