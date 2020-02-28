package config

import (
	"TalkFTWDiscordBot/voice"
	"log"
)

type Environment struct {
	InfoLogger   *log.Logger
	ErrorLogger  *log.Logger
	Config       *Configuration
	ServersVoice *voice.ServersVoiceActivity
}
