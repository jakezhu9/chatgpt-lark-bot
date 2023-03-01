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
	LarkBaseUrl       string `yaml:"lark_base_url"`
	Port              int    `yaml:"port"`
}

var defaultConf = []byte(`app_id: 
app_secret: 
verification_token: 
event_encrypt_key: 
open_ai_key: 
lark_base_url: https://open.larksuite.com # for feishu, use https://open.feishu.cn
port: 3000
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
