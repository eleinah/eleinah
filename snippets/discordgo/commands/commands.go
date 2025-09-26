package commands

import "github.com/bwmarrin/discordgo"

func setupSlashCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{Name: "ping", Description: "pong!"},
		{Name: "greet", Description: "greet the ol' fashion way"},
		{Name: "add", Description: "add two numbers together",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "first",
					Description: "first number",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "second",
					Description: "second number",
					Required:    true,
				},
			},
		},
	}
}
