package main

import (
	"log"
	"os"

	bot "codecosta.com/palico/app"
	"codecosta.com/palico/app/apis"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordAppID := os.Getenv("PALICO_DISCORD_APP_ID")
	discordBotToken := os.Getenv("PALICO_DISCORD_BOT_TOKEN")

	apis.CostaAPIBotToken = os.Getenv("COSTA_API_BOT_TOKEN")
	apis.CostaAPIPort = os.Getenv("COSTA_API_PORT")

	bot.DiscordAppID = discordAppID
	bot.DiscordBotToken = discordBotToken
	bot.Run()
}
