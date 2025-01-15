package main

import (
	"os"
)

type Config struct {
	LibraryPath string
}

func LoadConfig() Config {
	var cfg Config

	cfg.LibraryPath = os.Getenv("LIBRARY_PATH")

	return cfg
}
