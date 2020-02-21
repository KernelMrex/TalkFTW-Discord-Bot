package main

import (
	"TalkFTWDiscordBot/config"
	"github.com/DiscordBotList/go-dbl"
	"log"
	"os"
)

var Env *config.Environment

func init() {


	Env = &config.Environment{
		ErrorLogger: log.New(os.Stderr, "ERROR | ", log.Lshortfile|log.Ltime),
		InfoLogger:  log.New(os.Stdout, "INFO  | ", log.Lshortfile|log.Ltime),
	}
}

func main() {
	listener := dbl.NewListener("AuthToken", NewUserInChannelHandler)

	if err := listener.Serve(":9090"); err != nil {
		Env.ErrorLogger.Fatalln("[ main ]", err)
	}
}
