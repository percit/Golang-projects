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
// BOT LOGIC ###################################################################


	dg.AddHandler(helloWorldTest)
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates



//###################################################################

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
}

func helloWorldTest(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore all messages created by the bot itself
	}
	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "hello1")
	}
	if m.Content == "" {
		s.ChannelMessageSend(m.ChannelID, "Blad, content pusty")
	}
	s.ChannelMessageSend(m.ChannelID, string(m.Content))//"!test costam" gives "!test costam"
}
