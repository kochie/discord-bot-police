package commands

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

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
		},
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
			err := rdb.HSet(ctx, "settings", "AOE_TAUNTS", "true").Err()
			if err != nil {
				log.Println(err)
				errorResponse(err, s, i)
				return
			}
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
			// Disable AOE taunts
			err := rdb.HSet(ctx, "settings", "AOE_TAUNTS", "false").Err()
			if err != nil {
				log.Println(err)
				errorResponse(err, s, i)
				return
			}
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
}
