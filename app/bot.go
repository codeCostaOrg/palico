package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"codecosta.com/palico/app/commands"
	"codecosta.com/palico/app/utils"
	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session
var DiscordAppID string
var DiscordBotToken string

func Run() {
	// init discord
	Discord, err := discordgo.New("Bot " + DiscordBotToken)
	if err != nil {
		log.Fatal(err)
	}

	// create discord connection
	Discord.Open()
	defer Discord.Close()
	fmt.Println("Bot started...")

	fmt.Println("Setting up commands...")
	_, err = Discord.ApplicationCommandBulkOverwrite(DiscordAppID, "", commands.CommandList)
	if err != nil {
		log.Fatal(err)
	}

	Discord.AddHandler(func(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
		switch interaction.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commands.CommandHandlers[interaction.ApplicationCommandData().Name]; ok {
				h(discord, interaction)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := commands.CommandHandlers[interaction.MessageComponentData().CustomID]; ok {
				h(discord, interaction)
			}
		}
	})

	// create a list so we can remove the commands later
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.CommandList))
	for i, v := range commands.CommandList {
		cmd, err := Discord.ApplicationCommandCreate(Discord.State.User.ID, "", v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
		fmt.Println("Command created:", cmd.Name)
		registeredCommands[i] = cmd
	}
	fmt.Println("Commands set up!")

	// create a channel to receive OS signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // notify the channel when signals are received
	<-c                            // send the signal to this function

	// remove commands before exiting
	fmt.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := Discord.ApplicationCommandDelete(Discord.State.User.ID, "", v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
	fmt.Println("Commands removed!")
}

func handleIncomingMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// ignore self messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	username := message.Author.Username
	channelID := message.ChannelID

	utils.LogCommand(username, message.Content)

	// public commmands
	switch message.Content {
	case "me":
		utils.SendUserMessage(discord, channelID, username, message.Author.ID)
	case "ping":
		utils.SendUserMessage(discord, channelID, username, "pong")
	}

	if username != "luan.me" {
		utils.SendUserMessage(discord, channelID, username, "Sorry, I'm not allowed to talk to you.")
		return
	}
}
