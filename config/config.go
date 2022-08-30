package config

import "os"

type dbConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type AppConfig struct {
	ENV                string
	AppName            string
	AppPort            int
	JWTSecretKey       []byte
	JWTExpiryInMinutes int64
	DBConfig           dbConfig
}

func GetENV(key, defaultVal string) string {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultVal
	}
	return envVal
}

var Config = AppConfig{
	ENV:                GetENV("ENV", "testing"),
	AppName:            "Teyvat Times - API",
	AppPort:            8080,
	JWTSecretKey:       []byte("very-secret"),
	JWTExpiryInMinutes: 60,
	DBConfig: dbConfig{
		DBHost:     GetENV("DBHost", "localhost"),
		DBPort:     GetENV("DBPort", "5432"),
		DBUser:     GetENV("DBUser", "postgres"),
		DBPassword: GetENV("DBPassword", "123456"),
		DBName:     GetENV("DBName", "teyvat-1"),
	},
}
