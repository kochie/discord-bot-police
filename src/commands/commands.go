package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/database"
	"github.com/kochie/discord-bot-police/src/directives"
	"log"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "status",
		Description: "Get the status of the bot",
	},
	{
		Name:        "enable",
		Description: "Enable a bot function",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "taunts",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Enable AOE voice taunts",
			},
			{
				Name:        "furry",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Enable furry detection",
			},
			{
				Name:        "anime",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Enable anime detection",
			},
			{
				Name:        "simp",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Enable simp detection",
			},
			{
				Name:        "commie",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Enable commie detection",
			},
		},
	},
	{
		Name:        "disable",
		Description: "Disable a bot function",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "taunts",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Disable AOE voice taunts",
			},
			{
				Name:        "furry",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Disable furry detection",
			},
			{
				Name:        "anime",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Disable anime detection",
			},
			{
				Name:        "simp",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Disable simp detection",
			},
			{
				Name:        "commie",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Disable communist detection",
			},
		},
	},
	{
		Name:        "scores",
		Description: "Get the scores of the server",
	},
}

func errorResponse(err error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(err)
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "An error occurred",
		},
	})

	if err != nil {
		log.Println(err)
	}
}

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"status": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "I'm good bro, stop asking",
			},
		})

		if err != nil {
			log.Println(err)
		}
	},
	"enable": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		switch options[0].Name {
		case "taunts":
			// Enable AOE taunts
			database.UpdateSettings("AOE_TAUNTS", "true")
		case "furry":
			// Enable furry detection
			database.UpdateSettings("DETECT_FURRY", "true")
		case "anime":
			// Enable anime detection
			database.UpdateSettings("DETECT_ANIME", "true")
		case "simp":
			// Enable simp detection
			database.UpdateSettings("DETECT_SIMP", "true")
		case "commie":
			// Enable commie detection
			database.UpdateSettings("DETECT_COMMIE", "true")
		default:
			errorResponse(fmt.Errorf("unknown option"), s, i)
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("I have enabled %s", options[0].Name),
			},
		})

		if err != nil {
			log.Println(err)
		}
	},
	"disable": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		switch options[0].Name {
		case "taunts":
			// Enable AOE taunts
			database.UpdateSettings("AOE_TAUNTS", "false")
		case "furry":
			// Enable furry detection
			database.UpdateSettings("DETECT_FURRY", "false")
		case "anime":
			// Enable anime detection
			database.UpdateSettings("DETECT_ANIME", "false")
		case "simp":
			// Enable simp detection
			database.UpdateSettings("DETECT_SIMP", "false")
		case "commie":
			// Enable commie detection
			database.UpdateSettings("DETECT_COMMIE", "false")
		default:
			errorResponse(fmt.Errorf("unknown option"), s, i)
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("I have disabled %s", options[0].Name),
			},
		})

		if err != nil {
			log.Println(err)
		}
	},
	"scores": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		directives.OutputScores(s, i)
	},
}
