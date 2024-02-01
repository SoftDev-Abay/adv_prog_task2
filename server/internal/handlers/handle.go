package handlers

import "renting/internal/config"

type Handlers struct {
	c *config.Config
}

func NewHandlers(config *config.Config) *Handlers {
	return &Handlers{c: config}
}
