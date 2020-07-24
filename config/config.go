package config

import (
	"encoding/json"
	"fmt"
)

var Conf = &Config{}

type IMConfig struct {
	AppName   string `json:"app_id"`
	AppID     string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	URL       string `json:"url"`
}

type RTCConfig struct {
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

type Config struct {
	IM  IMConfig  `json:"im"`
	RTC RTCConfig `json:"rtc"`
}

func (c *Config) Init(appID, appSecret string) error {
	ok, err := download(FILE, "", VERSION)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("下载动态链接库失败")
	}

	if err := _binding.init(); err != nil {
		return err
	}

	jsonData, err := _binding.deCrypto(appID, appSecret)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(jsonData), c); err != nil {
		return err
	}
	return nil
}
