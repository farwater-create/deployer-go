package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Resolver = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		name := i.ApplicationCommandData().Name
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Command failed for ", name, " ", r)
			}
		}()
		handler, ok := Handlers[name]
		if !ok {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Something went wrong while trying to run that command :(",
				},
			})
			return
		}
		handler(s, i)
	}()
}
