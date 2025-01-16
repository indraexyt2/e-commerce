package helpers

import (
	"github.com/joho/godotenv"
)

var (
	EnvMap map[string]string
	err    error
)

func SetupConfig() {
	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		Logger.Error("Error loading .env file: ", err)
	}
}

func GetEnv(key string) string {
	return EnvMap[key]
}
