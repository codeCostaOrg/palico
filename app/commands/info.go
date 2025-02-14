package commands

import (
	"codecosta.com/palico/app/apis"
	"codecosta.com/palico/app/utils"
	"github.com/bwmarrin/discordgo"
)

func GetMonsterInfo(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "/info")

	options := interaction.ApplicationCommandData().Options
	monsterName := options[0].Value.(string)

	res, err := apis.GetAPIMonsterInfo(monsterName, interaction.Member.User.Username)
	if err != nil {
		utils.LogSystemError("GetMonsterInfo.GetMonsterInfo", err.Error())
	}
	if res.Error != "" {
		utils.DiscordRespondWithError(discord, interaction, res.Error)
		return
	}

	// embed image
	embeds := []*discordgo.MessageEmbed{
		{
			Title: res.Name,
			Image: &discordgo.MessageEmbedImage{
				URL: res.ImgURL,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: res.AvatarURL,
			},
			Description: "**This is currently a Work In Progress. Some information may be missing or incorrect. Please contact the admins if you have any information to add.**",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Element",
					Value:  utils.ConvertElementsToString(res.Elements),
					Inline: true,
				},
				{
					Name:   "Weaknesses",
					Value:  utils.ConvertElementsToString(res.Weaknesses),
					Inline: true,
				},
				{
					Name:   "Resistances",
					Value:  utils.ConvertElementsToString(res.Resistances),
					Inline: true,
				},
				{
					Name:  "Traps Allowed",
					Value: utils.ConvertTrapsToString(res.ShockTrappable, res.PitfallTrappable),
				},
				{
					Name:  "Tail Severable",
					Value: utils.ConvertBoolToString(res.TailSeverable),
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: string(res.Size) + " " + string(res.Species),
			},
		},
	}

	err = discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title:  res.Name,
			Embeds: embeds,
		},
	})
	if err != nil {
		utils.LogDiscordError("GetMonsterInfo.InteractionRespond", err.Error())
	}
}
