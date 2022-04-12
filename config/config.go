package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	IP   string
	Port int
	//FeatureFlag int
}

type DatabaseConfig struct {
	Host       string
	User       string
	Password   string
	Authsource string
	Port       int
	TimeOut    int
	DBNAME     string
}

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

func LoadConfig() (config Config, err error) {
	viper.SetDefault("App.Port", 9999)
	viper.SetDefault("App.IP", "0.0.0.0")
	viper.SetDefault("DB.Host", "localhost")
	viper.SetDefault("DB.User", "mongouser")
	viper.SetDefault("DB.Password", "changeme")
	viper.SetDefault("DB.Authsource", "admin")
	viper.SetDefault("DB.Port", 27017)
	viper.SetDefault("DB.TimeOut", 2)
	viper.SetDefault("DB.DBNAME", "icanhazquotes")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("ICHQ")

	viper.AutomaticEnv()

	viper.BindEnv("App.Port", "ICHQ_APP_PORT")
	viper.BindEnv("App.Hostname", "ICHQ_APP_IP")
	viper.BindEnv("DB.Host", "ICHQ_DB_HOST")
	viper.BindEnv("DB.User", "ICHQ_DB_USER")
	viper.BindEnv("DB.Password", "ICHQ_DB_PASSWORD")
	viper.BindEnv("DB.Authsource", "ICHO_DB_AUTHSOURCE")
	viper.BindEnv("DB.Port", "ICHO_DB_PORT")
	viper.BindEnv("DB.TimeOut", "ICHO_DB_TIMEOUT")
	viper.BindEnv("DB.DBNAME", "ICHO_DB_DBNAME")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	err = viper.Unmarshal(&config)
	log.Printf("App Port %d", config.App.Port)
	return
}
