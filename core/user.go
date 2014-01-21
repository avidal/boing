package core

import "github.com/avidal/boing/server"

// Represents a user of the bouncer; pulled from the configuration
type User struct {
    Username string
    Password UserPassword

    Servers map[string]server.Server
}

type UserPassword struct {
    Algorithm  string
    Iterations int
    Salt       string
    Hash       string
}

func (u *User) CheckPassword(p string) bool {
    hashed := u.Password.Hash

    // TODO: Implement hashed passwords with various algorithms
    return hashed == p
}
