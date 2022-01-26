package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	port        int    `envconfig:"PORT" default:"8080"`
	dbUrl       string `envconfig:"DB_URL" default:"postgres: //db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	jaegerUrl   string `envconfig:"JAEGER_URL" default:"http: //jaeger:16686"`
	sentryUrl   string `envconfig:"SENTRY_url" default:"http: //sentry:9000"`
	kafkaBroker int    `envconfig:"kafka_broker" default:"kafka:9092"`
	someAppId   string `envconfig:"some_appId" default:"testid"`
	someAppKey  string `envconfig:"some_appKey" default:"testkey"`
}

func Load() (*Configuration, error) { //, error
	config := Configuration{}
	err := envconfig.Process("configuration", &config)
	if err != nil {
		fmt.Println("config невозможно загрузить: %W", err)
		os.Exit(1)
	}
	return &Configuration{port: 8181}, err
}
