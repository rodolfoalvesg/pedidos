package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is a struct that holds the configuration for the application.
type Config struct {
	HTTP  HTTPConfig
	DB    ConfigPostgres
	Redis ConfigRedis
	ES    ConfigElastic
	User  ConfigUserAPI
}

type HTTPConfig struct {
	Host string `envconfig:"HTTP_HOST" default:"app-order"`
	Port string `envconfig:"HTTP_PORT" default:"3001"`
}

type ConfigPostgres struct {
	DriverName string `envconfig:"DB_DRIVER" default:"postgres"`
	Host       string `envconfig:"DB_HOST" default:"db_user"`
	Port       string `envconfig:"DB_PORT" default:"5432"`
	User       string `envconfig:"DB_USER" default:"postgres"`
	Password   string `envconfig:"DB_PASSWORD" default:"postgres"`
	Name       string `envconfig:"DB_NAME" default:"dev"`
}

type ConfigRedis struct {
	Host     string `envconfig:"REDIS_HOST" default:"redis_user"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
}

type ConfigElastic struct {
	Host      string `envconfig:"ELASTICSEARCH_HOST" default:"http://elasticsearch:9200"`
	Username  string `envconfig:"ELASTICSEARCH_USER" default:"elastic"`
	Password  string `envconfig:"ELASTICSEARCH_PASSWORD" default:"changeme"`
	NameIndex string `envconfig:"ELASTICSEARCH_NAME" default:"orders"`
}

type ConfigUserAPI struct {
	APIURL string `envconfig:"API_URL" required:"true" default:"http://app-user:3000"`
}

func LoadConfig() (Config, error) {
	var noPrefix = ""

	// Load the configuration from the environment.
	var cfg Config
	if err := envconfig.Process(noPrefix, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (db *ConfigPostgres) DSN() string {
	return "host=" + db.Host + " port=" + db.Port + " user=" + db.User + " password=" + db.Password + " dbname=" + db.Name + " sslmode=disable"
}
