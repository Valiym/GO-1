package main

import (
	"fmt"
	"github.com/Valiym/GO/Lesson-9/configuration"
)

func main() {
	app, err := configuration.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&app)
}
