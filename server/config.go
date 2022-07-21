package server

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DBConfig stores config information for a single DB
type DBConfig struct {
	Name              string `yaml:"db_name"`
	User              string `yaml:"db_user"`
	PasswordEncrypted string `yaml:"db_password_encrypted"`
	Hostname          string `yaml:"db_host"`
}

// IConfig is an interface that Config implements
type IConfig interface {
	ReadConfig(path string) (*Config, error)
}

// Config contains vauthgo configuration information
type Config struct {
	BindAddress    string   `yaml:"bind_address"`
	CrossServerKey string   `yaml:"cross_server_key"`
	DBConfig       DBConfig `yaml:"db_config"`
	InstallVersion string   `yaml:"install_version"`
	InstallGitHash string   `yaml:"install_githash"`
}

// ReadConfig reads config from path to yaml file
func ReadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	c := Config{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return &c, nil
}
