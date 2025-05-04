package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config AppConfig

type AppConfig struct {
	Port                    string `mapstructure:"APP_PORT"`
	AppName                 string `mapstructure:"APP_NAME"`
	AppEnv                  string `mapstructure:"APP_ENV"`
	SignatureKey            string `mapstructure:"SIGNATURE_KEY"`
	RateLimiterMaxRequest   int    `mapstructure:"RATE_LIMITER_MAX_REQUEST"`
	RateLimiterTimeSecond   int    `mapstructure:"RATE_LIMITER_TIME_SECOND"`
	JWTSecretKey            string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationTime       int    `mapstructure:"JWT_EXPIRATION_TIME"`
	DbHost                  string `mapstructure:"DB_HOST"`
	DbPort                  string `mapstructure:"DB_PORT"`
	DbUsername              string `mapstructure:"DB_USER"`
	DbPassword              string `mapstructure:"DB_PASS"`
	DbName                  string `mapstructure:"DB_NAME"`
	DbMaxOpenConnection     int    `mapstructure:"DB_MAX_OPEN_CONNECTION"`
	DbMaxLifetimeConnection int    `mapstructure:"DB_MAX_LIFETIME_CONNECTION"`
	DbMaxIdleConnection     int    `mapstructure:"DB_MAX_IDLE_CONNECTION"`
	DbMaxIddleTime          int    `mapstructure:"DB_MAX_IDDLE_TIME"`
}

func Init() {
	// Load file .env
	viper.SetConfigFile(".env")
	viper.SetConfigType("env") // optional, tapi membantu
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error reading .env file: %v", err)
	}

	viper.AutomaticEnv() // baca juga dari ENV OS, kalau ada

	// Binding ke struct
	if err := viper.Unmarshal(&Config); err != nil {
		logrus.Fatalf("failed to bind config: %v", err)
	}

	logrus.Infof("App Name: %s", Config.AppName)
}

