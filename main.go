package main

import (
	"fmt"
	"github.com/Valiym/GO/Lesson-8/configuration"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		fmt.Println("config could not be load: ", err)
	}
	fmt.Println(config)
}
