package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
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
	// Find channel
	c, err := s.State.Channel(u.ChannelID)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", "cannot determine user channel")
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		Env.ErrorLogger.Println("[ playUserSoundHandler ]", "cannot determine guild")
		return
	}

	// Example: sending message when user connected to a voice channel
	for _, ch := range g.Channels {
		if ch.Type == discordgo.ChannelTypeGuildText {
			message := fmt.Sprintf("User %s connected voice channel %s", u.UserID, u.ChannelID)
			if _, err := s.ChannelMessageSend(ch.ID, message); err != nil {
				Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", err)
				return
			}
			break
		}
	}
}
