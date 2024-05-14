package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

const ENDPOINT string = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func getWeather(messageContent string) *discordgo.MessageSend {
	regex, err := regexp.Compile(`\d{5}`)
	if err != nil {
		log.Fatal(err)
	}

	zipCode := regex.FindString(messageContent)
	if zipCode == "" {
		return &discordgo.MessageSend{
			Content: "Please supply a zip code.",
		}
	}

	weatherURL := fmt.Sprintf("%szip=%s&appid=%s&units=metric", ENDPOINT, zipCode, OpenWeatherApiKey)

	// * Http Client
	client := http.Client{Timeout: 5 * time.Second}

	res, err := client.Get(weatherURL)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error. Please try again later.",
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error. Please try again later.",
		}
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return &discordgo.MessageSend{
			Content: "Sorry, Could found this location.",
		}
	}

	var data WeatherData
	json.Unmarshal([]byte(body), &data)

	city := data.Name
	conditions := data.Weather[0].Description
	temperature := strconv.FormatFloat(data.Main.Temp, 'f', 2, 64)
	humidaty := strconv.FormatFloat(data.Main.Humidity, 'f', 2, 64)
	wind := strconv.FormatFloat(data.Wind.Speed, 'f', 2, 64)

	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Current weather",
				Description: "Weather in " + city,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Conditions",
						Value:  conditions,
						Inline: true,
					},
					{
						Name:   "Temperature",
						Value:  temperature + "°C",
						Inline: true,
					},
					{
						Name:   "Humidity",
						Value:  humidaty + "%",
						Inline: true,
					},
					{
						Name:   "Wind",
						Value:  wind + "m/s",
						Inline: true,
					},
				},
			},
		},
	}

	return embed
}
