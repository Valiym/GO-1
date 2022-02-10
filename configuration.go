package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

type Configuration struct {
	Port        int    `json:"port"`
	DbURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   int    `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func Load() (*Configuration, error) {
	var config = Configuration{}
	var fileName = ""

	flag.IntVar(&config.Port, "port", 8080, "port 1000-65535")
	flag.StringVar(&config.DbURL, "dbURL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
	flag.StringVar(&config.JaegerURL, "jaegerURL", "http://jaeger:16686", "jaeger url")
	flag.StringVar(&config.SentryURL, "sentryURL", "http://sentry:9000", "sentry url")
	flag.StringVar(&config.KafkaBroker, "kafka", "http://kafka:9092", "kafka broker")
	flag.IntVar(&config.SomeAppID, "someAppID", 5002, "some app id")
	flag.StringVar(&config.SomeAppKey, "someAppKey", "someAppKey", "some app key")
	flag.StringVar(&fileName, "filename", "", "json config file")
	flag.Parse()

	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("can't open file: %w", err)
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatalf("can't close file: %v", err)
			}
		}()

		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			return nil, fmt.Errorf("can't decode json file: %w", err)
		}
	}
	err := validate(&config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}

func validate(cfg *Configuration) error {
	if cfg.Port > 65535 || cfg.Port < 1000 {
		return fmt.Errorf("invalid port value")
	}
	if !validURL(cfg.DbURL) {
		return fmt.Errorf("invalid dbURL value")
	}
	if !validURL(cfg.JaegerURL) {
		return fmt.Errorf("invalid jaegerURL value")
	}
	if !validURL(cfg.SentryURL) {
		return fmt.Errorf("invalid sentryURL value")
	}
	if !validURL(cfg.KafkaBroker) {
		return fmt.Errorf("invalid kafkaBroker value")
	}
	if cfg.SomeAppID < 1000 {
		return fmt.Errorf("invalid someAppID value")
	}
	return nil
}

func validURL(str string) bool {
	p, err := url.Parse(str)
	return err == nil || p.Scheme != "" || p.Host != ""
}
