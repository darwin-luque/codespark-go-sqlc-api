package config

type ServerInterface interface {
	Port() string
}

type ServerConfig struct {
	port string
}

func NewServerConfig() ServerInterface {
	port, _ := getEnvVar("SERVER_PORT", "8080")

	return &ServerConfig{port}
}

func (s *ServerConfig) Port() string {
	return s.port
}
