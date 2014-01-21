package commands

import (
    "fmt"
    "log"
    "net"
    "strings"

    "github.com/avidal/boing/core"
    "github.com/spf13/cobra"
)

var Config *core.Config

var BoingCmd = &cobra.Command{
    Use:   "boing",
    Short: "boing is a flexible, modern IRC bouncer.",
    Long: `A modern IRC bouncer built with love by avidal and friends in Go.
    
Complete documentation is available at https://github.com/avidal/boing`,

    Run: func(cmd *cobra.Command, args []string) {
        InitializeConfig()
        fmt.Printf("Hello! Binding to %v:%v\n", Config.Bind, Config.Port)
        fmt.Printf("Admins: %#v\n", Config.Admins)

        start()
    },
}

var CfgFile string

func Execute() {
    BoingCmd.Execute()
}

func init() {
    BoingCmd.PersistentFlags().StringVarP(&CfgFile, "config", "c",
        "config.toml", "configuration file")
}

func InitializeConfig() {
    Config = core.SetupConfig(&CfgFile)
}

func start() {
    // Starts the proxy server by:
    // - Connect to all servers in the configuration
    // - Open a listening socket for client connections

    // Each user in the configuration is handled by a single goroutine, each of
    // those goroutines will then spawn additional ones to deal with
    // communication

    for _, user := range Config.Users {
        log.Println("Creating goroutine for user " + user.Username)
        log.Printf("Servers: %#v\n", user.Servers)
        go startUserProxy(&user)
    }

    // Setup our listener for clients to connect to
    bind := fmt.Sprintf("%s:%d", Config.Bind, Config.Port)
    listener, err := net.Listen("tcp", bind)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println("Listening on", bind)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalln("Error accepting connection:", err)
        }

        go accept(conn)
    }
}

func accept(c net.Conn) {
    // Deals with the initial handshake with a newly connected client, then
    // sets up a communication channel
    log.Printf("Accepting connection from %s\n", c.RemoteAddr())

    // Setup a buffer to read socket data into
    var buf [512]byte

    // Loop until we read a PASS command, which will tell us what client and
    // server they are attempting to connect to
    // TODO: We should handle CAP LS here, I think. Or buffer everything that
    // comes before the PASS command and let the Proxy instance deal with it.
    for {
        l, err := c.Read(buf[0:])
        if err != nil {
            log.Println("Connection closed.")
            return
        }
        msg := string(buf[0:l])
        log.Println("RECV:", msg)
        if strings.HasPrefix(msg, "PASS ") == false {
            continue
        }

        // We read a PASS command, split on spaces, second part is the password
        p := strings.Fields(msg)
        p0 := strings.SplitN(p[1], ":", 2)
        p1 := strings.SplitN(p0[1], "@", 2)

        username := p0[0]
        passwd := p1[0]
        server := p1[1]

        log.Printf("Username: %s, Password: %s, Server: %s", username, passwd, server)

        // Now, let's see if this user exists, and if so, does the password
        // match and is the server configured
        u, err := Config.GetUser(username)
        log.Println("User:", u)
        log.Println("Error?", err)

        if u == nil {
            log.Println("No user found.")
            c.Write([]byte(":localhost NOTICE AUTH :*** No such user.\r\n"))
            c.Close()
        }

        if u.CheckPassword(passwd) == false {
            log.Println("Password mismatch!")
            m := []byte(":localhost NOTICE AUTH :*** Invalid password!\r\n")
            c.Write(m)
            c.Close()
        }

    }
}

func startUserProxy(u *core.User) {}
