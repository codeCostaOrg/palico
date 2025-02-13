package commands

import (
	"fmt"

	"codecosta.com/palico/app/models"
	"codecosta.com/palico/app/utils"
	"github.com/bwmarrin/discordgo"
)

func GetMonsterInfo(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "/info")

	options := interaction.ApplicationCommandData().Options
	monsterName := options[0].Value.(string)
	var mapImgURL string

	switch monsterName {
	case string(models.RATHALOS):
		mapImgURL = "https://cdn.discordapp.com/attachments/1137468055216193647/1339390723690070048/MHW-Rathalos_Render_001.webp?ex=67ae8c51&is=67ad3ad1&hm=57f5f6415ae839475c63dede155d2f41b98622ba2e053427baddca5d25d548ae&"
	default:
		errMessage := fmt.Sprintf("Sorry %s is not supported", monsterName)
		utils.LogDiscordError("GetMonsterInfo.monsterName.default", errMessage)
		err := utils.DiscordRespondWithError(discord, interaction, errMessage)
		if err != nil {
			utils.LogDiscordError("GetMonsterInfo.monsterName.default.InteractionRespond", err.Error())
		}

		return
	}

	// embed map
	embeds := []*discordgo.MessageEmbed{
		{
			Title: monsterName,
			Image: &discordgo.MessageEmbedImage{
				URL: mapImgURL,
			},
		},
	}

	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title:  monsterName,
			Embeds: embeds,
		},
	})
	if err != nil {
		utils.LogDiscordError("GetMonsterInfo.InteractionRespond", err.Error())
	}
}
