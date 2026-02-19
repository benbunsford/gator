package main

import (
	"fmt"
	"log"

	"github.com/benbunsford/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("ben")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cfg)
}
