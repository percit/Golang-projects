package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Token string
var WeatherApiKey string
const openWeatherMapURL = "http://api.openweathermap.org/data/2.5/weather"

type OpenWeatherMapResponse struct {
	Name      string `json:"name"`
	Main      Main   `json:"main"`
	Weather   []Weather `json:"weather"`
	Timestamp int64     `json:"dt"`
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token") //./discord-bot -t BOT_TOKEN
	flag.StringVar(&WeatherApiKey, "w", "", "Weather api key") //./discord-bot -t BOT_TOKEN -d WEATHER-API-KEY
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

func parseCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore all messages created by the bot itself
	}
	if m.Content == "!test" {
		s.ChannelMessageSend(m.ChannelID, "hello1")
	} else if m.Content == "!weather" {
		s.ChannelMessageSend(m.ChannelID, getWeatherData())
	}
	if m.Content == "" {
		s.ChannelMessageSend(m.ChannelID, "Error, no content")
	}
	// s.ChannelMessageSend(m.ChannelID, string(m.Content))
}
func getWeatherData() string {
	city := "New York"

	weather, err := GetWeather(city, WeatherApiKey)
	if err != nil {
		fmt.Println("Error getting weather:", err)
		return "Error getting weather" + err.Error()
	}
	return weather
}

//HELPER FUNCTIONS

func GetWeather(city string, apiKey string) (string, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherMapURL, city, apiKey)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var openWeatherMapResponse OpenWeatherMapResponse

	if err := json.NewDecoder(resp.Body).Decode(&openWeatherMapResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	weather := openWeatherMapResponse.Weather[0].Description
	temperature := fmt.Sprintf("%.1f", openWeatherMapResponse.Main.Temperature)
	cityName := openWeatherMapResponse.Name

	return fmt.Sprintf("Current weather in %s: %s, Temperature: %s℃", cityName, weather, temperature), nil
}


//todo przenies rzeczy do paczki package weather jak zacznie dzialac