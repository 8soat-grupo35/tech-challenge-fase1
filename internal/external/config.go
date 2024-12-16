package external

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost     string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var (
	runOnce sync.Once
	config  Config
)

func GetConfig() Config {
	runOnce.Do(func() {
		cfg, err := initConfig()
		if err != nil {
			fmt.Println(context.Background(), err, "could not load usecase configuration")
		}
		config = Config{
			ServerHost: cfg.GetString("server.host"),
			DatabaseConfig: DatabaseConfig{
				Host:     cfg.GetString("DATABASE_HOST"),
				Port:     cfg.GetString("DATABASE_PORT"),
				User:     cfg.GetString("DATABASE_USER"),
				Password: cfg.GetString("DATABASE_PASSWORD"),
				DbName:   cfg.GetString("DATABASE_DBNAME"),
			},
		}

		fmt.Println(config)
	})

	return config
}

func initConfig() (viper.Viper, error) {
	cfg := viper.New()
	var err error
	initDefaults(cfg)
	cfg.AutomaticEnv()
	// workaround because viper does not resolve envs when unmarshalling
	for _, key := range cfg.AllKeys() {
		val := cfg.Get(key)
		cfg.Set(key, val)
	}
	fmt.Println(cfg)
	return *cfg, err
}

func initDefaults(config *viper.Viper) {
	config.SetDefault("server.host", "0.0.0.0:8000")
	config.SetDefault("DATABASE_HOST", "postgres")
	config.SetDefault("DATABASE_PORT", "5432")
	config.SetDefault("DATABASE_USER", "root")
	config.SetDefault("DATABASE_PASSWORD", "root")
	config.SetDefault("DATABASE_DBNAME", "root")
}
