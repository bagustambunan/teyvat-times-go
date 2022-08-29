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

func getENV(key, defaultVal string) string {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultVal
	}
	return envVal
}

var Config = AppConfig{
	ENV:                getENV("ENV", "testing"),
	AppName:            "Project Nemesis - Prototype",
	AppPort:            8080,
	JWTSecretKey:       []byte("very-secret"),
	JWTExpiryInMinutes: 60,
	DBConfig: dbConfig{
		DBHost:     getENV("DBHost", "localhost"),
		DBPort:     getENV("DBPort", "5432"),
		DBUser:     getENV("DBUser", "postgres"),
		DBPassword: getENV("DBPassword", "123456"),
		DBName:     getENV("DBName", "nemesis-4"),
	},
}
