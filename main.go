package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = file.SetUser("Apolo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Read())

}
