package utils

import (
	"codecosta.com/palico/app/models"
	"github.com/bwmarrin/discordgo"
)

func InfoCommandChoices() []*discordgo.ApplicationCommandOptionChoice {
	choices := []*discordgo.ApplicationCommandOptionChoice{}

	for _, monster := range models.MonsterList {
		choice := &discordgo.ApplicationCommandOptionChoice{
			Name:  string(monster),
			Value: monster,
		}

		choices = append(choices, choice)
	}

	return choices
}
