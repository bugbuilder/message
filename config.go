package main

import (
	"io/ioutil"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultFile = "/data/config.yml"
)

// Config represent the config to use
type Config struct {
	Message string
	Factor  struct {
		min int
		max int
	}
	Teams map[string][]string
}

func (c *Config) Read(file string) (*Config, error) {
	if file == "" {
		file = defaultFile
		glog.V(3).Infoln("Using default file: %v", defaultFile)
	}

	yml, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	glog.V(3).Infoln("Decoding configuration file")

	err = yaml.Unmarshal(yml, c)
	return c, nil
}
