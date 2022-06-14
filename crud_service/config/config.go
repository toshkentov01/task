package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
)

//Config ...
type Config struct {
	Environment     string
	LogLevel        string
	RPCPort         string
	DataServiceHost string
	DataServicePort int
}

func load() *Config {
	return &Config{
		Environment:     cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		LogLevel:        cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug")),
		RPCPort:         cast.ToString(getOrReturnDefault("RPC_PORT", ":9000")),
		DataServiceHost: cast.ToString(getOrReturnDefault("DATA_SERVICE_HOST", "localhost")),
		DataServicePort: cast.ToInt(getOrReturnDefault("DATA_SERVICE_PORT", 9001)),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

var (
	instance *Config
	once     sync.Once
)

//Get ...
func Get() *Config {
	once.Do(func() {
		instance = load()
	})

	return instance
}
