package main

import (
	"github.com/ajswetz/go-gator/internal/config"
	"github.com/ajswetz/go-gator/internal/database"
)

type state struct {
	configuration *config.Config
	db *database.Queries
}
