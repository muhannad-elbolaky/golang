package bot

import "fmt"

var (
	BotToken          string
	OpenWeatherApiKey string
)

func Initualize() {
	fmt.Println("Got Keys: ", OpenWeatherApiKey, BotToken)
}
