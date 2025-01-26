package config

type ServerInterface interface {
	Port() string
	Host() string
}

type ServerConfig struct {
	port string
	host string
}

func NewServerConfig() ServerInterface {
	port, _ := getEnvVar("SERVER_PORT", "8080")
	host, _ := getEnvVar("SERVER_HOST", "localhost")

	return &ServerConfig{port, host}
}

func (s *ServerConfig) Port() string {
	return s.port
}

func (s *ServerConfig) Host() string {
	return s.host
}
