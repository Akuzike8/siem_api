package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WAZUH_API_IP     string
	WAZUH_API_PORT   string
	WAZUH_CONNECTION_STRING string
	VEL_CONFIG_PATH  string
}

func LoadConfig() *Config {
	// Load environment variables from a .env file, if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	return &Config{
		WAZUH_API_IP:     getEnv("WAZUH_API_IP", "172.31.30.70"),
		WAZUH_API_PORT:   getEnv("WAZUH_API_PORT", "55000"),
		WAZUH_CONNECTION_STRING: getEnv("WAZUH_CONNECTION_STRING","https://172.31.30.70:55000"),
		VEL_CONFIG_PATH: getEnv("VEL_CONFIG_PATH","api.config.yaml"),
	}
}

// getEnv reads an environment variable and falls back to a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}