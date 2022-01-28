package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	Port        int    `envconfig:"PORT" default:"8080"`
	DbUrl       string `envconfig:"DB_URL" default:"postgres: //db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	JaegerUrl   string `envconfig:"JAEGER_URL" default:"http: //jaeger:16686"`
	SentryUrl   string `envconfig:"SENTRY_url" default:"http: //sentry:9000"`
	KafkaBroker string `envconfig:"kafka_broker" default:"kafka:9092"`
	SomeAppId   string `envconfig:"some_appId" default:"testid"`
	SomeAppKey  string `envconfig:"some_appKey" default:"testkey"`
}

func Load() (Configuration, error) {
	var config Configuration
	err := envconfig.Process("configuration", &config)
	if err != nil {
		fmt.Println("config невозможно загрузить: %W", err)
		os.Exit(1)
	}
	return config, err
}
