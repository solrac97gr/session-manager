package sessionmanager_test

import (
	"errors"
	"testing"
	"time"

	sessionmanager "github.com/solrac97gr/session-manager"
)

func TestSession_NewSession(t *testing.T) {
	cases := map[string]struct {
		data map[string]interface{}
	}{
		"empty": {
			data: map[string]interface{}{},
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)

			if session.ID == "" {
				t.Error("Session ID is empty")
			}

			if len(session.Data) != len(tc.data) {
				t.Errorf("Session data length is not equal to expected data length. Expected: %d, Actual: %d", len(tc.data), len(session.Data))
			}
		})
	}

}

func TestSession_Get(t *testing.T) {
	cases := map[string]struct {
		data     map[string]interface{}
		key      string
		expected interface{}
	}{
		"empty": {
			data:     map[string]interface{}{},
			key:      "key",
			expected: nil,
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
			key:      "key",
			expected: "value",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)

			actual, err := session.Get(tc.key)
			if err != nil {
				if err.Error() != "key not found" {
					t.Error(err)
				}
			}
			if actual != tc.expected {
				t.Errorf("Session data is not equal to expected data. Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}

func TestSession_Set(t *testing.T) {
	cases := map[string]struct {
		data  map[string]interface{}
		key   string
		value interface{}
		err   error
	}{
		"empty": {
			data:  map[string]interface{}{},
			key:   "key",
			value: "value",
			err:   nil,
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
			key:   "key2",
			value: "value2",
			err:   nil,
		},

		"with data and key exists": {
			data: map[string]interface{}{
				"key": "value",
			},
			key:   "key",
			value: "value2",
			err:   errors.New("key key already exists, for replace delete it first"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)

			err := session.Set(tc.key, tc.value)
			if err != nil {
				if err.Error() != tc.err.Error() {
					t.Error(err)
				}
			}
		})
	}
}

func TestSession_Delete(t *testing.T) {
	cases := map[string]struct {
		data map[string]interface{}
		key  string
		err  error
	}{
		"empty": {
			data: map[string]interface{}{},
			key:  "key",
			err:  errors.New("key not found"),
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
			key: "key",
			err: nil,
		},

		"with data and key not exists": {
			data: map[string]interface{}{
				"key": "value",
			},
			key: "key2",
			err: errors.New("key not found"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)

			err := session.Delete(tc.key)
			if err != nil {
				if err.Error() != tc.err.Error() {
					t.Error(err)
				}
			}
		})
	}
}

func TestSession_SessionId(t *testing.T) {
	cases := map[string]struct {
		data map[string]interface{}
	}{
		"empty": {
			data: map[string]interface{}{},
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)

			if session.SessionId() == "" {
				t.Error("Session ID is empty")
			}
		})
	}
}

func TestSession_SetExpirationTime(t *testing.T) {
	cases := map[string]struct {
		data map[string]interface{}
	}{
		"empty": {
			data: map[string]interface{}{},
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)
			expected := time.Now().Add(80 * time.Minute)

			session.SetExpirationTime(expected)

			if session.ExpirationTime != expected {
				t.Errorf("Session expiration time is not equal to expected expiration time. Expected: %v, Actual: %v", expected, session.ExpirationTime)
			}
		})
	}
}

func TestSession_IsExpired(t *testing.T) {
	cases := map[string]struct {
		data    map[string]interface{}
		expired bool
	}{
		"empty": {
			data: map[string]interface{}{},
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
		},

		"already expired": {
			data: map[string]interface{}{
				"key": "value",
			},
			expired: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			session := sessionmanager.NewSession(tc.data)
			if tc.expired {
				session.Expired = true
				if !session.IsExpired() {
					t.Error("Session is not expired")
				}
				return
			}
			session.SetExpirationTime(time.Now().Add(-80 * time.Minute))

			if !session.IsExpired() {
				t.Error("Session is not expired")
			}
		})
	}
}

func TestSessionManager_IsActive(t *testing.T) {
	cases := map[string]struct {
		data map[string]interface{}
	}{
		"empty": {
			data: map[string]interface{}{},
		},

		"with data": {
			data: map[string]interface{}{
				"key": "value",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			sessionManager := sessionmanager.NewSessionManager()
			s, _ := sessionManager.CreateSession()
			s.Set("data", tc.data)

			if !s.IsActive() {
				t.Error("Session is not active")
			}
		})
	}
}
