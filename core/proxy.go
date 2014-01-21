package core

import (
    "github.com/avidal/boing/client"
    "github.com/avidal/boing/server"
)

// Deals with proxying and routing data from one or more clients to the
// connected server. Each server for each user has its own Proxy instance.
type Proxy struct {
    Clients map[string]*client.Client
    Server  *server.Server
}
