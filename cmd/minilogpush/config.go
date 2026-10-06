package main

import (
	"log"
	"os"
)

type Config struct {
	Addr     string
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func loadConfig() Config {
	config := Config{
		Addr: "80",
	}
	return config
}

func (app *Config) setupLogger() {
	app.InfoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime)
}
