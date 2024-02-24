package server

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host           string
	Port           string
	AllowedOrigins []string
	SqlitePath     string
	TwilioSID      string
	TwilioAuth     string
	TwilioFrom     string
}

func NewConfig() Config {
	host := getEnvDefault("HOST", "0.0.0.0")
	port := getEnvDefault("PORT", "8080")
	return Config{
		Host: host,
		Port: port,
		AllowedOrigins: strings.Split(
			getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",",
		),
		SqlitePath: getEnvDefault("SQLITE_PATH", "data.db"),
		TwilioSID:  getEnvDefault("TWILIO_SID", ""),
		TwilioAuth: getEnvDefault("TWILIO_AUTH", ""),
		TwilioFrom: getEnvDefault("TWILIO_FROM", ""),
	}
}

func getEnvDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
