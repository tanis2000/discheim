package service

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	session *discordgo.Session
	token string
	srv *Server
}

func NewBot(token string, srv *Server) *Bot {
	bot := &Bot {
		token: token,
		srv: srv,
	}
	return bot
}

func (bot *Bot)Create() error {
	var err error
	bot.session, err = discordgo.New("Bot " + bot.token)
	if err != nil {
		return err
	}
	return nil
}

func (bot * Bot)Connect() error {
	bot.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Print("bot up and running")
	})
	bot.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "!restart" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Restarting Valheim server")
			if err != nil {
				log.Println(err.Error())
				return
			}
			out, err := bot.srv.Restart()
			if err != nil {
				log.Println(err.Error())
				return
			}
			_, err = s.ChannelMessageSend(m.ChannelID, out)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}

		// If the message is "pong" reply with "Ping!"
		if m.Content == "!status" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Valheim server status:")
			if err != nil {
				log.Println(err.Error())
				return
			}
			out, err := bot.srv.Status()
			if err != nil {
				log.Println(err.Error())
				_, err = s.ChannelMessageSend(m.ChannelID, "Cannot retrieve the server status")
				if err != nil {
					log.Println(err.Error())
					return
				}
				return
			}
			_, err = s.ChannelMessageSend(m.ChannelID, out)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	})
	bot.session.Identify.Intents = discordgo.IntentsGuildMessages
	err := bot.session.Open()
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) Close() error {
	return bot.session.Close()
}