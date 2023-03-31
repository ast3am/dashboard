package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.skillbox.ru/denis_parovoi/profession-godev/finalWork/project/pkg/logging"
)

type Config struct {
	Listen struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Data      DataPath `yaml:"data_path"`
	CacheTime int      `yaml:"cache_time"`
}

type DataPath struct {
	SMS string `yaml:"sms"`
	MMS struct {
		URL string `yaml:"url"`
	} `yaml:"mms"`
	Email     string `yaml:"email"`
	Billing   string `yaml:"billing"`
	VoiceCall string `yaml:"voice_call"`
	Support   struct {
		URL string `yaml:"url"`
	} `yaml:"support"`
	Incidents struct {
		URL string `yaml:"url"`
	} `yaml:"incidents"`
}

var cfg *Config

func GetConfig(path string) *Config {
	logger := logging.GetLogger()
	logger.Info().Msg("Read config")
	cfg = &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		logger.Fatal().Msgf("%s", err)
	}
	return cfg
}
