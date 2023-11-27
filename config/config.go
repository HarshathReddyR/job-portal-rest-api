package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig      appConfig
	DatabaseCOnfig databaseConfig
	RedisDB        redisDB
	Keys           keys
}

type appConfig struct {
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT,required=true"`
}
type databaseConfig struct {
	// Host string `env:"DB_HOST,required=true"`
	// User string `env:"DB_USER,required=true"`
	// Password string `env:"DB_PASSWORD,required=true"`
	// Dbname string `env:"DB_DBNAME,required=true"`
	// Port string `env:"DB_PORT,required=true"`
	// Sslmode string `env:"DB_SSLMODE,required=true"`
	// Timezone string `env:"DB_TIMEZONE,required=true"`
	DB_DSN string `env:"DB_DSN,required=true"`
}
type redisDB struct {
	Addr     string `env:"REDIS_ADRESS,required=true"`
	Password string `env:"REDIS_PASSWORD,required=true"`
	Db       string `env:"REDIS_DB,required=true"`
}
type keys struct {
	Public  string `env:"PUBLIC_KEY,required=true"`
	Private string `env:"PRAVIT_KEY,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
