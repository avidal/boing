package server

type Server struct {
    // TODO: Better name for the socket addr?
    ServerName string `toml:"server"`

    RealName string `toml:"real"`
    Nick     string
    Nick2    string
    Nick3    string
}
