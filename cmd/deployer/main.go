package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/humbertovnavarro/deployer-go/pkg/commands"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}
	s.AddHandlerOnce(func(s *discordgo.Session, c *discordgo.Ready) {
		_, err := s.ApplicationCommandBulkOverwrite(os.Getenv("APP_ID"), os.Getenv("GUILD_ID"), commands.ApplicationCommands)
		if err != nil {
			panic(err)
		}
	})
	s.AddHandler(commands.Resolver)
	go func() {
		err = s.Open()
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			fmt.Println(input)
		}
	}()
	forever := make(chan bool)
	<-forever

}
