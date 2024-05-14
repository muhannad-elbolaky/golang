package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken          string
	OpenWeatherApiKey string
)

func Initualize() {
	session, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(Ready)
	session.AddHandler(MessageCreate)

	session.Open()
	defer session.Close()

	fmt.Println("Press CTRL-C to exit.")
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel
}

func Ready(session *discordgo.Session, context *discordgo.Ready) {
	fmt.Printf("%s(%s) is online!\n", context.User.Username, context.User.ID)
}

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, ".ping"):
		currentTime := time.Now().UTC().UnixMilli()
		messageTime := message.Message.Timestamp.UTC().UnixMilli()
		answer := fmt.Sprintf("Latency: %v", currentTime-messageTime)
		session.ChannelMessageSend(message.ChannelID, answer)
	case strings.Contains(message.Content, ".weather"):
		currentWeather := getWeather(message.Content)
		session.ChannelMessageSendComplex(message.ChannelID, currentWeather)
	}
}
