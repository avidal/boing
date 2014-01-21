package commands

import (
    "fmt"
    "log"

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
}

func startUserProxy(u *core.User) {
}
