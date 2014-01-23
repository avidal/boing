package core

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "path"
    "strconv"
    "strings"

    "github.com/BurntSushi/toml"
)

type Config struct {
    ConfigFile string

    Bind string
    Port int

    Admins []string
    Users  []User `toml:"user"`
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

func (c *Config) GetUser(name string) *User {
    // Attempts to find a user by username, returns a pointer to the User if
    // found, otherwise returns nil with an error
    for _, u := range c.Users {
        if name == u.Username {
            return &u
        }
    }
    return nil
}

// Unmarshaler for passwords in the config file
func (p *UserPassword) UnmarshalText(text []byte) error {

    // Split the text on the token
    toks := strings.SplitN(string(text), "$", 4)

    // If the length is not 4, bail
    if len(toks) != 4 {
        return errors.New("invalid password specification")
    }

    p.Algorithm = toks[0]
    p.Salt = toks[2]
    p.Hash = toks[3]

    if toks[1] != "" {
        iters, err := strconv.ParseInt(toks[1], 10, 8)
        if err != nil {
            return errors.New("invalid iteration count")
        }
        p.Iterations = int(iters)
    } else {
        p.Iterations = 0
    }

    return nil

}
