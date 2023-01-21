package commands

import "github.com/bwmarrin/discordgo"

func init() {
	ApplicationCommands = append(ApplicationCommands, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "sends ping",
	})
	Handlers["ping"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		})
	}
}
