package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/namsral/flag"
	"net/url"
)

type Configuration struct {
	Port        int    `envconfig:"PORT" default:"8080"`
	DbUrl       string `envconfig:"DB_URL" default:"postgres: //db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	JaegerUrl   string `envconfig:"JAEGER_URL" default:"http: //jaeger:16686"`
	SentryUrl   string `envconfig:"SENTRY_URL" default:"http: //sentry:9000"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" default:"kafka:9092"`
	SomeAppId   int    `envconfig:"SOME_APP_ID" default:"5000"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY" default:"someAppKey"`
}

func Load() (*Configuration, error) {
	var config = Configuration{}
	err := envconfig.Process("configuration", &config)
	if err != nil {
		fmt.Println("config could not be load: ", err)
	}
	flag.IntVar(&config.Port, "port", 8080, "port 1000-65535")
	flag.StringVar(&config.DbUrl, "dbUrl", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
	flag.StringVar(&config.JaegerUrl, "jaeger", "http://jaeger:16686", "jaeger url")
	flag.StringVar(&config.SentryUrl, "sentry", "http://sentry:9000", "sentry url")
	flag.StringVar(&config.KafkaBroker, "kafka", "http://kafka:9092", "kafka broker")
	flag.IntVar(&config.SomeAppId, "someAppId", 5000, "some app id")
	flag.StringVar(&config.SomeAppKey, "someAppKey", "someAppKey", "some app key")
	flag.Parse()
	err = validate(&config)
	if err != nil {
		return &config, err
	}

	return &config, err
}

func validate(cfg *Configuration) error {
	if cfg.Port > 65535 || cfg.Port < 1000 {
		fmt.Println("invalid port value")
	}
	if !valid(cfg.DbUrl) {
		fmt.Println("invalid dbUrl value")
	}
	if !valid(cfg.JaegerUrl) {
		fmt.Println("invalid jaegerUrl value")
	}
	if !valid(cfg.SentryUrl) {
		fmt.Println("invalid sentryUrl value")
	}
	if !valid(cfg.KafkaBroker) {
		fmt.Println("invalid kafkaBroker value")
	}
	if cfg.SomeAppId < 1000 {
		fmt.Println("invalid someAppId value")
	}
	return nil
}

func valid(str string) bool {
	p, err := url.Parse(str)
	return err == nil || p.Scheme != "" || p.Host != ""
}
