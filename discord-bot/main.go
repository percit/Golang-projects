package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"encoding/json"
	"net/http"
	"bytes"

	"discord-bot/weatherApi"

	"github.com/bwmarrin/discordgo"
)

var Token string
var WeatherApiKey string
var ChatGPTAPIKey string


func init() {
	flag.StringVar(&Token, "t", "", "Bot Token") //./discord-bot -t BOT_TOKEN
	flag.StringVar(&WeatherApiKey, "w", "", "Weather api key") //./discord-bot -t BOT_TOKEN -d WEATHER-API-KEY
	flag.StringVar(&ChatGPTAPIKey, "c", "", "ChatGPT api key") //./discord-bot -t BOT_TOKEN -d WEATHER-API-KEY -c CHAT-GPT-KEY
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
	dg.AddHandler(parseCommands)
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates
//###################################################################

	err = dg.Open()// Open a websocket connection to Discord and begin listening.
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

func parseCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore all messages created by the bot itself
	}
	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "hello1")
	} else if m.Content == "!weather" {
		s.ChannelMessageSend(m.ChannelID, getWeatherData())
	} else if m.Content == "" {
		s.ChannelMessageSend(m.ChannelID, "Error, no content")
	}
	if strings.HasPrefix(m.Content, "!") {
		parts := strings.SplitN(m.Content[1:], " ", 2)
		if len(parts) == 2 {
			command := "!" + parts[0]
			message := parts[1]
			if command == "!chatGPT" {
				s.ChannelMessageSend(m.ChannelID, getChatGPTResponse(message))//not working, probably openAI api has some problem
			}
		}
	}
}
func getWeatherData() string {
	city := "Wrocław"

	weather, err := weatherApi.GetWeather(city, WeatherApiKey)
	if err != nil {
		fmt.Println("Error getting weather:", err)
		return "Error getting weather" + err.Error()
	}
	return weather
}

func getChatGPTResponse(parsedInput string) string {
	payload := map[string]string{
		"text": parsedInput,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("https://api.openai.com/v1/engine/"+ChatGPTAPIKey+"/completions", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	choices, ok := response["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		fmt.Println( "the case where choices is missing or empty")
		return ""
	}

	choice := choices[0].(map[string]interface{})
	text, ok := choice["text"].(string)
	if !ok {
		fmt.Println( "handle the case where text is missing or not a string")
		return ""
	}

	return text
} 