package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
)

var (
	EnvMap map[string]string
	err    error
)

func SetupConfig() {
	EnvMap, err = godotenv.Read(".env")
	fmt.Printf("Error loading .env file: %v\n", err)
}

func GetEnv(key string) string {
	return EnvMap[key]
}
