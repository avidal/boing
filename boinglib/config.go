package boinglib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ConfigFile string

	Bind string
	Port int
}

var c Config

func SetupConfig(f *string) *Config {
	cfg := c.findConfigFile(*f)
	c.ConfigFile = cfg

	c.Bind = "0.0.0.0"
	c.Port = 6667

	c.readInConfig()

	return &c
}

func (c *Config) readInConfig() {
	file, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return
	}

	if _, err := toml.Decode(string(file), &c); err != nil {
		fmt.Printf("Error parsing config: %s", err)
		os.Exit(1)
	}
}

func (c *Config) findConfigFile(f string) string {
	// If it's an absolute path, use it
	if path.IsAbs(f) {
		return f
	}

	// Otherwise, join up the pwd with the filename and use that
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get working directory to find configuration file.")
	}

	return path.Join(pwd, f)
}
