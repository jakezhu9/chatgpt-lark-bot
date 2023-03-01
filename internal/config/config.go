package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	AppID             string `yaml:"app_id"`
	AppSecret         string `yaml:"app_secret"`
	VerificationToken string `yaml:"verification_token"`
	EventEncryptKey   string `yaml:"event_encrypt_key"`
	OpenAIKey         string `yaml:"open_ai_key"`
}

var defaultConf = []byte(`app_id: 
app_secret: 
verification_token: 
event_encrypt_key: 
open_ai_key: 
`)

func LoadConfig(filename string) (*Config, error) {
	conf := &Config{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func WriteDefaultConfig(filename string) error {
	return os.WriteFile(filename, defaultConf, 0644)
}
