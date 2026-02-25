package main

import (
	"github.com/benbunsford/gator/internal/config"
	"github.com/benbunsford/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
