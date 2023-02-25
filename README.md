# Session Manager for Go 📂

This is a simple session manager for a go application. For example, you can use it to manage the user's login status.

Helps to manage user login information in the standard library of go. A lots of frameworks have their own session management, but some times I ue the standard library of go, so I write this package.

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
# Work in progress
- [x] Create a new session
- [x] Get a session
- [x] Delete a session
- [ ] Use concurrent map

# License
MIT License