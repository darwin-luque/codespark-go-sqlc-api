package config

type AuthInterface interface {
	Secret() string
}

type AuthConfig struct {
	secret string
}

func NewAuthConfig() (config AuthInterface, err error) {
	secret, err := getEnvVar("AUTH_SECRET", "random-secret")

	if err != nil {
		return nil, err
	}

	return &AuthConfig{
		secret,
	}, nil
}

func (s *AuthConfig) Secret() string {
	return s.secret
}
