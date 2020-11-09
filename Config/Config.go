package Config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const filename = "config.yaml"

var instance *Config

func init() {
	instance = getConfig()
}

type Config struct {
	Address string
	Debug   bool
}

func Address() string {
	return instance.Address
}

func Debug() bool {
	return instance.Debug
}

func Version() string {
	return "1.0.0"
}

func getConfig() (conf *Config) {

	if configYaml, err := ioutil.ReadFile(filename); err == nil {
		conf = &Config{}
		if err = yaml.Unmarshal(configYaml, conf); err == nil {
			return
		}
	}

	conf = defaultConfig()
	saveConfig(conf)

	return
}

func saveConfig(conf *Config) {
	if configYaml, err := yaml.Marshal(conf); err == nil {
		err = ioutil.WriteFile(filename, configYaml, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func defaultConfig() *Config {
	config := &Config{
		Address: ":1211",
		Debug:   true,
	}
	return config
}
