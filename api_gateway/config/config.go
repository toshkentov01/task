package config

import (
	"os"
	"sync"

	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

//Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	AppURL           string
	Environment      string
	ServerPort       int
	ServerHost       string
	LogLevel         string
	ServiceDir       string

	RedisHost     string
	RedisPort     int
	RedisPassword string

	CasbinConfigPath    string
	MiddlewareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	CrudServiceHost string
	CrudServicePort int

	DataServiceHost string
	DataServicePort int

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
}

func load() *Configuration {
	return &Configuration{
		AppURL:              cast.ToString(getOrReturnDefault("APP_URL", "localhost:8000")),
		ServerHost:          cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:          cast.ToInt(getOrReturnDefault("SERVER_PORT", "8000")),
		Environment:         cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		LogLevel:            cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug")),
		ServiceDir:          cast.ToString(getOrReturnDefault("CURRENT_DIR", "")),
		CtxTimeout:          cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		RedisHost:           cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost")),
		RedisPort:           cast.ToInt(getOrReturnDefault("REDIS_PORT", 0)),
		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf")),
		MiddlewareRolesPath: cast.ToString(getOrReturnDefault("MIDDLEWARE_ROLES_PATH", "./config/models.csv")),

		CrudServiceHost:  cast.ToString(getOrReturnDefault("CRUD_SERVICE_HOST", "")),
		CrudServicePort:  cast.ToInt(getOrReturnDefault("CRUD_SERVICE_PORT", 9000)),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
