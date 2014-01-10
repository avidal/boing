package commands

import (
	"fmt"

	"github.com/avidal/boing/boinglib"
	"github.com/spf13/cobra"
)

var Config *boinglib.Config

var BoingCmd = &cobra.Command{
	Use:   "boing",
	Short: "boing is a flexible, modern IRC bouncer.",
	Long: `A modern IRC bouncer built with love by avidal and friends in Go.
    
Complete documentation is available at https://github.com/avidal/boing`,

	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		fmt.Printf("Hello! Binding to %v:%v\n", Config.Bind, Config.Port)
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
	Config = boinglib.SetupConfig(&CfgFile)
}
