package main

import (
	"database/sql"
	"errors"
	"github.com/benbunsford/gator/internal/config"
	"github.com/benbunsford/gator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	currentState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		cmdMap: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
