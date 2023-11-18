package app

import (
	"log"

	"github.com/agusheryanto182/go-todo/models/config"
	"github.com/joeshaw/envdecode"
)

func NewConfig() *config.Global {
	var c config.Global
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
