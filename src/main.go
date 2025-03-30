package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/commands"
	"github.com/kochie/discord-bot-police/src/database"
	"github.com/kochie/discord-bot-police/src/directives"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var reg *regexp.Regexp
var ServerId = os.Getenv("SERVER_ID")

func init() {

	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	reg = r

}

func addCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func main() {
	token := os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(addCommands)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	log.Println("Adding commands...")
	// Need to overwrite the commands to get the command ID
	commands.SyncCommands(dg, ServerId, commands.Commands)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		return
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	processedString := reg.ReplaceAllString(strings.ToLower(m.Content), "")

	settings := database.GetSettings()

	if enabled, ok := settings["DETECT_ANIME"]; ok && enabled == "true" {
		directives.AnimeDetection(processedString, s, m)
	}
	if enabled, ok := settings["DETECT_SIMP"]; ok && enabled == "true" {
		directives.SimpDetection(processedString, s, m)
	}
	if enabled, ok := settings["DETECT_COMMIE"]; ok && enabled == "true" {
		directives.CommieDetection(processedString, s, m)
	}
	if enabled, ok := settings["DETECT_FURRY"]; ok && enabled == "true" {
		directives.FurryDetection(processedString, s, m)
	}
	if enabled, ok := settings["AOE_TAUNTS"]; ok && enabled == "true" {
		directives.AoeTaunts(processedString, s, m)
	}
	if enabled, ok := settings["DETECT_CELEB"]; ok && enabled == "true" {
		directives.CelebDetection(s, m)
	}

}
