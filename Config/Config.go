package Config

import (
	"crypto/sha1"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

const filename = "config.yaml"

var instance *Config

func init() {
	instance = getConfig()
}

type Config struct {
	Address     string
	Debug       bool
	MasterToken string
}

func Address() string {
	return instance.Address
}

func Debug() bool {
	return instance.Debug
}

func MasterToken() string {
	return instance.MasterToken
}

func Version() string {
	return "1.0.0"
}

func NewMasterToken() string {

	s := strconv.FormatInt(time.Now().UnixNano(), 10)
	r := strconv.Itoa(int(rand.Uint64()))
	h := sha1.New()
	h.Write([]byte(s))
	h.Write([]byte(r))
	bs := h.Sum(nil)

	instance.MasterToken = fmt.Sprintf("%x", bs)
	saveConfig(instance)

	return instance.MasterToken

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
		Address:     ":1211",
		Debug:       true,
		MasterToken: "",
	}
	return config
}
