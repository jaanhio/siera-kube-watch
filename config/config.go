package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Webhook struct {
		Enabled bool   `yaml:"enabled"`
		Url     string `yaml:"url"`
	}

	Livenesscheck struct {
		Enabled  bool   `yaml:"enabled"`
		Interval string `yaml:"interval"`
	}

	Slack struct {
		Enabled  bool   `yaml:"enabled"`
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Channel  string `yaml:"channel"`
	}

	Telegram struct {
		Enabled bool   `yaml:"enabled"`
		Token   string `yaml:"token"`
		ChatID  string `yaml:"chatID"`
	}

	Workplace struct {
		Enabled   bool   `yaml:"enabled"`
		ThreadKey string `yaml:"thread.key"`
		Token     string `yaml:"token"`
	}

	ExcludedReasons []string `yaml:"excluded.reasons,flow"`
	IncludedReasons []string `yaml:"included.reasons,flow"`

	IncludedNamespace []string `yaml:"included.namespaces,flow"`
}

var GlobalConfig = &Config{}

func (config *Config) Load() (err error) {

	var configpath string

	if env := strings.ToLower(os.Getenv("ENV")); env == "dev" {
		configpath = "./config.yaml"
	} else {
		configpath = "/usr/src/app/etc/siera-kube-watch/config.yaml"
	}

	yamlFile, err := ioutil.ReadFile(configpath)

	if err != nil {
		log.Fatalf("Error read config file: %v", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error unmarshal: %v", err)
	}

	return
}
