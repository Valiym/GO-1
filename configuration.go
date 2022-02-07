package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

type Configuration struct {
	Port        int    `json:"port"`
	DbUrl       string `json:"db_url"`
	JaegerUrl   string `json:"jaeger_url"`
	SentryUrl   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppId   int    `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func Load() (Configuration, error) {
	file, err := os.Open("configuration.json")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	app := Configuration{}

	err = json.NewDecoder(file).Decode(&app)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	err = validate(app)
	if err != nil {
		fmt.Println("can't validate config", err)
	}
	return app, nil
}

func validate(cfg Configuration) error {
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
