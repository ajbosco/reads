package goodreads

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Config contains configuration for authenticating with Goodreads API
type Config struct {
	DeveloperKey    string `yaml:"DeveloperKey"`
	DeveloperSecret string `yaml:"DeveloperSecret"`
	AccessToken     string `yaml:"AccessToken"`
	AccessSecret    string `yaml:"AccessSecret"`
}

// ReadConfig parses the config file into a Config object
func ReadConfig(path string) (*Config, error) {
	c := new(Config)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not read config from filepath: %v", path))
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// WriteConfig writes a Config object to the config file
func WriteConfig(config *Config, path string) error {
	c, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(path, c, 0644); err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not write config to filepath: %v", path))
	}
	return nil
}
