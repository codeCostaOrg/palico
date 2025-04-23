package commands

import (
	"codecosta.com/palico/app/utils"
	"github.com/bwmarrin/discordgo"
)

var CommandList = []*discordgo.ApplicationCommand{
	{
		Name:        "info",
		Description: "Get information about a monster",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "monster",
				Description: "Monster Name",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices:     utils.InfoCommandChoices(),
				Required:    true,
			},
		},
	},
}

var CommandHandlers = map[string]func(discord *discordgo.Session, interaction *discordgo.InteractionCreate){
	"info": GetMonsterInfo,
}

func test(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "test")

	content := "test response"

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
