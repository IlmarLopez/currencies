package config

import (
	"io/ioutil"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config is the configuration for the server.
const (
	DefaultServerPort         = 8080
	DefaultJWTExpirationHours = 72
)

// Config represents the configuration of the server.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.DSN, v.Required),
		v.Field(&c.JWTSigningKey, v.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger zap.SugaredLogger) (*Config, error) {
	// default config
	c := Config{
		ServerPort:    DefaultServerPort,
		JWTExpiration: DefaultJWTExpirationHours,
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
