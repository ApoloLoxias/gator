package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	currentState := state{Config: &file}
	commands := commands{Handlers: make(map[string]func(*state, command) error)}

	commands.Register("login", handlerLogin)

	osArgs := os.Args
	if len(osArgs) == 1 {
		fmt.Println("Please provide arguments when running the program")
		os.Exit(1)
	}
	args := osArgs[1:]

	cmd := command{
		Name: args[0],
		Args: args[1:],
	}

	err = commands.run(&currentState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)

}
