package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var conf config

type config struct {
	AppName      string
	AppDomain    string
	AppEnv       string
	ApiPort      string
	DBName       string
	DBUser       string
	DBPass       string
	DBHost       string
	DBPort       string
	JWTSecretKey string
}

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(fmt.Sprintf("failed to setup config, error: %s", err))
	} else {
		conf.AppName = getEnv("APP_NAME", "CMS")
		conf.AppDomain = getEnv("APP_DOMAIN", "localhost")
		conf.AppEnv = getEnv("APP_ENV", "local") // local or stg or prod
		conf.ApiPort = getEnv("API_PORT", "3000")
		conf.DBName = getEnv("DB_NAME", "cms")
		conf.DBUser = getEnv("DB_USER", "postgres")
		conf.DBPass = getEnv("DB_PASSWORD", "postgres")
		conf.DBHost = getEnv("DB_HOST", "localhost")
		conf.DBPort = getEnv("DB_PORT", "5432")
		conf.JWTSecretKey = getEnv("JWT_SECRET_KEY", "cms")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func AppName() string {
	return conf.AppName
}

func AppDomain() string {
	return conf.AppDomain
}

func AppEnv() string {
	return conf.AppEnv
}

func IsLocal() bool {
	return AppEnv() == "local"
}

func ApiPort() string {
	return conf.ApiPort
}

func DBName() string {
	return conf.DBName
}

func DBUser() string {
	return conf.DBUser
}

func DBPass() string {
	return conf.DBPass
}

func DBHost() string {
	return conf.DBHost
}

func DBPort() string {
	return conf.DBPort
}

func JWTSecretKey() string {
	return conf.JWTSecretKey
}
