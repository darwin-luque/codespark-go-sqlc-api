package config

type ServerInterface interface {
	Port() string
	Mode() string
}

type ServerConfig struct {
	port string
	mode string
}

func NewServerConfig() ServerInterface {
	port, _ := getEnvVar("SERVER_PORT", "8080")
	mode, _ := getEnvVar("SERVER_MODE", "debug")

	return &ServerConfig{port, mode}
}

func (s *ServerConfig) Port() string {
	return s.port
}

func (s *ServerConfig) Mode() string {
	return s.mode
}
