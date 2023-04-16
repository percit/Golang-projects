package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token") //./discord-name -t BOT_TOKEN
	flag.Parse()
}

func main() {
	// Create a new Discord session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(helloWorldTest)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer dg.Close()

	// BOT LOGIC ###################################################################

}

func helloWorldTest(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore all messages created by the bot itself
	}
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, world!")
	}
	s.ChannelMessageSend(m.ChannelID, "Hello World!")
	s.ChannelMessageSend(m.ChannelID, m.Content)
}
