package config

type DatabaseInterface interface {
	URL() string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDatabaseConfig() *DatabaseConfig {
	host, _ := getEnvVar("DB_HOST", "localhost")
	port, _ := getEnvVar("DB_PORT", "5432")
	user, err := getEnvVar("DB_USER", "")

	if err != nil {
		panic(err)
	}

	password, err := getEnvVar("DB_PASSWORD", "")

	if err != nil {
		panic(err)
	}

	database, err := getEnvVar("DB_DATABASE", "")

	if err != nil {
		panic(err)
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}

func (c *DatabaseConfig) URL() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database
}
