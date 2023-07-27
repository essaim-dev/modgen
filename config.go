package modgen

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// Config defines the format of a modgen configuration file.
type Config struct {
	Host    string         `yaml:"host"`
	Modules []ModuleConfig `yaml:"modules"`
}

// Config defines an individual module from a configuration file.
type ModuleConfig struct {
	Path    string `yaml:"path"`
	VCS     string `yaml:"vcs"`
	RepoURL string `yaml:"repo-url"`
}

// LoadConfig tries to load and parse a modgen configuration file at the given path.
func LoadConfig(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(raw, &config)

	return config, err
}

// withHost returns a copy of the ModuleConfig with the given host prepended to the Path field.
func (mc ModuleConfig) withHost(host string) ModuleConfig {
	mc.Path = path.Join(host, mc.Path)

	return mc
}
