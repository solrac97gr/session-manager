# Session Manager for Go 📂
[![GoDoc](https://godoc.org/github.com/solrac97gr/session-manager?status.svg)](https://godoc.org/github.com/solrac97gr/session-manager)
[![Go Report Card](https://goreportcard.com/badge/github.com/solrac97gr/session-manager)](https://goreportcard.com/report/github.com/solrac97gr/session-manager)
![Build workflow](https://github.com/solrac97gr/session-manager/actions/workflows/workflows.yml/badge.svg)
[![codecov](https://codecov.io/gh/solrac97gr/session-manager/branch/main/graph/badge.svg)](https://codecov.io/gh/solrac97gr/session-manager)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a simple session manager for a go application. Helps you to store information of user and manage it, set a expiration time and set a default session of a group of different sessions created.

# Documentation

Visit the [GoDoc](https://godoc.org/github.com/solrac97gr/session-manager) page for the full documentation.



# Usage

Install the package using the go get command:

```bash
go get github.com/solrac97gr/session-manager
```

Import the package in your code:

```go
import "github.com/solrac97gr/session-manager"
```

And use it:

## Example: Create a new session
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/solrac97gr/session-manager"
)

func main() {
    user := struct{
        Name string
        Age int
    }{
        Name: "Solrac",
        Age: 20,
    }{}

    sm := sessionmanager.NewSessionManager()

    // Create a new session
    s,err := sm.CreateSession()
    if err != nil {
        panic(err)
    }

    s.Set("user", user)

    fmt.Println(s.Get("user"))
}
```

## Example: Get a session
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/solrac97gr/session-manager"
)

func main() {
    sm := sessionmanager.NewSessionManager()

    // Get a session
    s,err := sm.GetSession("session-id")
    if err != nil {
        panic(err)
    }

    fmt.Println(s.Get("user"))
}
```
## Example: Delete a session
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/solrac97gr/session-manager"
)

func main() {
    sm := sessionmanager.NewSessionManager()

    // Destroy a session
    err := sm.DestroySession("session-id")
    if err != nil {
        panic(err)
    }
}
```
## Example: Set a default session and get it
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/solrac97gr/session-manager"
)

func main() {
    user := struct{
        Name string
        Age int
    }{
        Name: "Solrac",
        Age: 20,
    }


    sm := sessionmanager.NewSessionManager()
    s,_:= sm.CreateSession()

    s.Set("user", user)

    // Set a default session
    sm.SetAsDefaultSession(s.SessionId())

    // Get a session
    s,err := sm.GetDefaultSession()
    if err != nil {
        panic(err)
    }

    fmt.Println(s.Get("user"))
}
```

## Example: Activated avoid expired sessions and set a expiration time in a session

By default, the session manager does not avoid expired sessions, so if you want to avoid expired sessions, you must activate it.

The defatul expiration time is 5 minutes, so if you want to change expiration time, you must set it.

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/solrac97gr/session-manager"
)

func main() {
    user := struct{
        Name string
        Age int
    }{
        Name: "Solrac",
        Age: 20,
    }

    sm := sessionmanager.NewSessionManager()

    // Activate avoid expired sessions
    sm.SetAvoidExpiredSessions(true)

    s,_:= sm.CreateSession()

    s.Set("user", user)
    s.SetExpirationTime(time.Now().Add(1440 * time.Minute))

    // Set a default session
    sm.SetAsDefaultSession(s.SessionId())

    // Get a session
    s,err := sm.GetDefaultSession()
    if err != nil {
        panic(err)
    }

    fmt.Println(s.Get("user"))
}
```


# Work in progress and completed
- [x] Create a new session
- [x] Get a session
- [x] Delete a session
- [x] Use concurrent map
- [x] Add a session expiration time
- [x] Add a active indicator

# License
MIT License
