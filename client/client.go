package client

import "net"

type Client struct {
    Conn *net.Conn
}
