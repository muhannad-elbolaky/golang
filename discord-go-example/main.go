package main

import (
	"discord-weather-bot/bot"
	"discord-weather-bot/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	botToken := utils.GetEnvVar("BOT_TOKEN")
	openWeatherApiKey := utils.GetEnvVar("OPENWEATHER_API_KEY")

	bot.BotToken = botToken
	bot.OpenWeatherApiKey = openWeatherApiKey
	bot.Initualize()
}
