package configs

import (
	"log"

	"github.com/spf13/viper"
)

type APIConfig struct {
	Port string
}

type DatabaseConfig struct {
	Name string
	User string
	Pass string
	Host string
	Port string
}

type config struct {
	API      APIConfig
	Database DatabaseConfig
}

var cfg = new(config)

func Init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	Init()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error: %v", err)
		}
	}
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.Database = DatabaseConfig{
		Name: viper.GetString("database.name"),
		User: viper.GetString("database.user"),
		Pass: viper.GetString("database.pass"),
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
	}
	return nil
}

func GetAPiInformations() string {
	return cfg.API.Port
}

func GetDatabaseInformations() DatabaseConfig {
	return cfg.Database
}
