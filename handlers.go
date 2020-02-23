package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func playUserSoundHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "+bump") {
		return
	}

	// Find channel
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", "cannot determine user channel")
		return
	}

	// Find the guild(server) for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", "cannot determine guild")
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			message := fmt.Sprintf("Playing your sound... In channel %s", vs.ChannelID)

			if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
				Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
				return
			}

			// TODO: play sound
			return
		}
	}

	// User not in channel
	if _, err := s.ChannelMessageSend(m.ChannelID, "You're not in a channel"); err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", err)
		return
	}
}

func voiceStateUpdateHandler(s *discordgo.Session, u *discordgo.VoiceStateUpdate) {
	// Ignore all messages created by the bot itself
	if u.UserID == s.State.User.ID {
		return
	}

	// Find channel
	c, err := s.State.Channel(u.ChannelID)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine user channel")
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine guild")
		return
	}

	voiceConn, err := s.ChannelVoiceJoin(g.ID, c.ID, false, true)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine guild")
		return
	}
	defer func() {
		if err := voiceConn.Disconnect(); err != nil {
			Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "error while disconnecting", err)
			return
		}
	}()


	time.Sleep(50 * time.Millisecond)
	voiceConn.Speaking(true)

	// Send the buffer data.
	for _, buff := range Announcement {
		voiceConn.OpusSend <- buff
	}

	voiceConn.Speaking(false)
	time.Sleep(250 * time.Millisecond)
}
