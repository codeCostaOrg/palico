package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendUserMessage(discord *discordgo.Session, channelID string, user string, message string) {
	log.Printf("<- %s: %s", user, message)
	discord.ChannelMessageSend(channelID, message)
}

func DiscordRespondWithError(discord *discordgo.Session, interaction *discordgo.InteractionCreate, responseError string) error {
	return discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: responseError,
		},
	})
}

func EnforceDMOnly(discord *discordgo.Session, interaction *discordgo.InteractionCreate) bool {
	if interaction.User == nil {
		DiscordRespondWithError(discord, interaction, "This command is only available in DMs")
		return false
	}
	return true
}
