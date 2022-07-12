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
	ServerPort int    `yaml:"server_port"`
	DSN        string `yaml:"dsn"`       // data source name for connecting to the database.
	APIKey     string `yaml:"api_key"`   // API key for the currency api.
	Intervalo  int    `yaml:"intervalo"` // Intervalo in seconds to request the currency API.
	Timeout    int    `yaml:"timeout"`   // Timeout in seconds to configure the timeout for requesting the currency API.
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.DSN, v.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger *zap.SugaredLogger) (*Config, error) {
	// default config
	c := Config{
		ServerPort: DefaultServerPort, // the default server port is 8080
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
