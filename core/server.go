package core

type Server struct {
    ServerName string `toml:"server"`

    RealName string `toml:"real"`
    Nick     string
    Nick2    string
    Nick3    string
}
