package config

import "log"

type Environment struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Config      *Configuration
}
