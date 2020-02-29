package main

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func voiceStateUpdateHandler(s *discordgo.Session, u *discordgo.VoiceStateUpdate) {
	// Ignore all messages created by the bot itself
	if u.UserID == s.State.User.ID {
		return
	}

	// Find channel
	c, err := s.State.Channel(u.ChannelID)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine user channel", err)
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine guild", err)
		return
	}

	Env.ServersVoice.ServerLock(g.ID)
	voiceConn, err := s.ChannelVoiceJoin(g.ID, c.ID, false, true)
	if err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot determine guild", err)
		return
	}
	defer func() {
		Env.ServersVoice.ServerUnlock(g.ID)
		if err := voiceConn.Disconnect(); err != nil {
			Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "error while disconnecting", err)
			return
		}
		voiceConn.Close()
	}()

	time.Sleep(50 * time.Millisecond)
	if err := voiceConn.Speaking(true); err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot start speaking:", err)
		return
	}

	// Send the buffer data.
	for _, buff := range Announcement {
		voiceConn.OpusSend <- buff
	}

	if err := voiceConn.Speaking(false); err != nil {
		Env.ErrorLogger.Println("[ voiceStateUpdateHandler ]", "cannot stop speaking:", err)
		return
	}

	time.Sleep(250 * time.Millisecond)
}
