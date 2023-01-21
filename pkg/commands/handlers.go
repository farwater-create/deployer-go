package commands

import "github.com/bwmarrin/discordgo"

type commandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

var Handlers map[string]commandHandler = make(map[string]commandHandler)
