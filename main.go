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

	state := config.State{Config: &file}
	commands := config.Commands{Handlers: make(map[string]func(*config.State, config.Command) error)}

	commands.Register("login", config.HandlerLogin)

	osArgs := os.Args
	if len(osArgs) == 1 {
		fmt.Println("Please provide arguments when running the program")
		os.Exit(1)
	}
	args := osArgs[1:]

	cmd := config.Command{
		Name: args[0],
		Args: args[1:],
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)

}
