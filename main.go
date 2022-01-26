package main

import (
	"fmt"
	"github.com/Valiym/GO/Lesson-8/configuration"
	"os"
)

func main() {

	config, err := configuration.Load()
	if err != nil {
		fmt.Println("config could not be loade: %W", err)
		os.Exit(1)
	}
	fmt.Println(config)
}
