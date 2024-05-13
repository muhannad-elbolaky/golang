package main

import (
	"discord-weather-bot/bot"
	"discord-weather-bot/utils"
)

func main() {
	botToken := utils.GetEnvVar("BOT_TOKEN")
	openWeatherApiKey := utils.GetEnvVar("OPEN_WEATHER_API_KEY")

	bot.BotToken = botToken
	bot.OpenWeatherApiKey = openWeatherApiKey
	bot.Initualize()
}
