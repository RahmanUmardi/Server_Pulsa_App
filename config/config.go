package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName       string `json:"IssuerName"`
	JwtSignatureKy   []byte `json:"JwtSignatureKy"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type Config struct {
	DBConfig
	ApiConfig
	TokenConfig
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}
	c.DBConfig = DBConfig{
		Host:     getEnv("DB_HOST", "167.172.91.111"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "rahasia"),
		Name:     getEnv("DB_NAME", "server_pulsa_db"),
		Driver:   getEnv("DB_DRIVER", "postgres"),
	}

	c.ApiConfig = ApiConfig{ApiPort: getEnv("API_PORT", "8080")}

	tokenExpire, _ := strconv.Atoi(getEnv("TOKEN_EXPIRE", "120"))
	c.TokenConfig = TokenConfig{
		IssuerName:       getEnv("TOKEN_ISSUE", "Enigma Camp Incubation Class"),
		JwtSignatureKy:   []byte(getEnv("TOKEN_SECRET", "Golang Incubation Class")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	}

	if c.Host == "" || c.Port == "" || c.User == "" || c.Name == "" || c.Driver == "" || c.ApiPort == "" ||
		c.IssuerName == "" || c.JwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil

}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
