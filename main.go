package main

import (
	"errors"
	"github.com/benbunsford/gator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	currentState := state{
		cfg: &cfg,
	}

	cmds := commands{
		cmdMap: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	userArgs := os.Args
	if len(userArgs) < 2 {
		log.Fatal(errors.New("No command provided. Provide a command after 'gator'."))
	}

	userCmd := command{
		name: userArgs[1],
		args: userArgs[2:],
	}

	err = cmds.run(&currentState, userCmd)
	if err != nil {
		log.Fatal(err)
	}
}
