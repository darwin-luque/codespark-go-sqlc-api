package config

type DatabaseInterface interface {
	DSN() string
}

type DatabaseConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	host, _ := getEnvVar("DB_HOST", "localhost")
	port, _ := getEnvVar("DB_PORT", "5432")
	user, err := getEnvVar("DB_USER", "")

	if err != nil {
		return nil, err
	}

	password, err := getEnvVar("DB_PASSWORD", "")

	if err != nil {
		return nil, err
	}

	name, err := getEnvVar("DB_NAME", "")

	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		name:     name,
	}, nil
}

func (c *DatabaseConfig) DSN() string {
	return "host=" + c.host + " port=" + c.port + " user=" + c.user + " password=" + c.password + " dbname=" + c.name + " sslmode=disable"
}
