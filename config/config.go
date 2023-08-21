package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type MySqlDBConfig struct {
	Host     string `mapstructure:"MYSQLDB_HOST"`
	Port     string `mapstructure:"MYSQLDB_PORT"`
	Username string `mapstructure:"MYSQLDB_USERNAME"`
	Password string `mapstructure:"MYSQLDB_PASSWORD"`
	Name     string `mapstructure:"MYSQLDB_NAME"`
}

type GoogleOauthConfig struct {
	ClientSecret       string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	ClientId           string `mapstructure:"GOOGLE_CLIENT_ID"`
	AccesTokenUrl      string `mapstructure:"GOOGLE_ACCESS_TOKEN_URL"`
	UserInformationUrl string `mapstructure:"GOOGLE_USER_INFORMATION_URL"`
	RedirectUri        string `mapstructure:"GOOGLE_REDIRECT_URI"`
}

type GithubOauthConfig struct {
	ClientSecret       string `mapstructure:"GITHUB_CLIENT_SECRET"`
	ClientId           string `mapstructure:"GITHUB_CLIENT_ID"`
	AccesTokenUrl      string `mapstructure:"GITHUB_ACCESS_TOKEN_URL"`
	UserInformationUrl string `mapstructure:"GITHUB_USER_INFORMATION_URL"`
}

type AppConfig struct {
	Port        int               `mapstructure:"APP_PORT"`
	JwtKey      string            `mapstructure:"JWT_SECRET_KEY"`
	MySqlDB     MySqlDBConfig     `mapstructure:",squash"`
	GithubOauth GithubOauthConfig `mapstructure:",squash"`
	GoogleOauth GoogleOauthConfig `mapstructure:",squash"`
}

var Config = &AppConfig{}

func LoadConfig() {

	if _, err := os.Stat("env/.env"); err != nil {
		panic(fmt.Sprintf("Unable to find env file : %v", err))
	}

	viper.AddConfigPath("env/")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Unable to read env: %v", err))
	}

	if err := viper.Unmarshal(Config); err != nil {
		panic(fmt.Sprintf("Unable to parse env: %v", err))
	}

}
