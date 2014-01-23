package proxy

import (
    "github.com/avidal/boing/client"
    "github.com/avidal/boing/core"
    "github.com/avidal/boing/server"
)

// Deals with proxying and routing data from one or more clients to the
// connected server. Each server for each user has its own Proxy instance.
type ProxyServer struct {
    User    *core.User
    Clients map[string]*client.Client
    Server  *server.Server

    Connections chan *client.Client
}

func NewProxy(user *core.User, server *server.Server) ProxyServer {
    p := ProxyServer{
        User:        user,
        Clients:     make(map[string]*client.Client),
        Server:      server,
        Connections: make(chan *client.Client),
    }

    return p
}

func (p *ProxyServer) Start() {}
