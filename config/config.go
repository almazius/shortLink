package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Postgres struct {
		User     *string `json:"user"`
		Password *string `json:"password"`
		Host     *string `json:"host"`
		Port     *string `json:"port"`
		DbName   *string `json:"dbName"`
	} `json:"postgres"`
	Redis struct {
		Network  *string `json:"network"`
		Host     *string `json:"host"`
		Port     *string `json:"port"`
		User     *string `json:"user"`
		Password *string `json:"password"`
		DbName   *int    `json:"dbName"`
	} `json:"redis"`
}

// LoadConfig Loads and reads the config
func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to read config, %v", err)
		return nil, err
	}
	return v, nil
}

// ParseConfig parse config
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
