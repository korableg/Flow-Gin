package Config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var instance *Config

func init() {
	instance = getConfig()
}

type Config struct {
	Address string
}

func Address() string {
	return instance.Address
}

func Version() string {
	return "1.0.0"
}

func getConfig() (conf *Config) {

	const filename = "config.yaml"

	if configYaml, err := ioutil.ReadFile(filename); err == nil {
		conf = &Config{}
		if err = yaml.Unmarshal(configYaml, conf); err == nil {
			return
		}
	}

	conf = defaultConfig()
	if configYaml, err := yaml.Marshal(conf); err == nil {
		ioutil.WriteFile(filename, configYaml, 0644)
	}

	return
}

func defaultConfig() *Config {
	config := &Config{
		Address: ":1211",
	}
	return config
}
