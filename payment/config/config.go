package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return GetEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return GetEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := GetEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("%s is invalid port value", portStr)
	}
	return port
}

func GetEnvironmentValue(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s environment variable is missing", key)
	}

	return val
}
