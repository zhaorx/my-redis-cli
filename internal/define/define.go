package define

var ConfigName = "my-redis-cli.conf"

type Config struct {
    Connections []*Connection `json:"connections"`
}

type Connection struct {
    Identity string `json:"identity"`
    Name     string `json:"name"`
    Addr     string `json:"addr"`
}
