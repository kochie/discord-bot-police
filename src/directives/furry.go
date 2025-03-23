package directives

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/database"
	"github.com/kochie/discord-bot-police/src/util"
	"log"
)

// FurryDetection is a function that detects furry content in a message, it then adds the tally of how furry a message is to the user's furry score
func FurryDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// If the processed string contains any of these words, add to the furry score
	if util.CompareList(processedString, []string{
		"furry",
		"furries",
		"uwu",
		"owo",
		"yiff",
		"yiffing",
		"yiffed",
		"fursuit",
		"fursona",
		"fur",
		"furs",
		"fursonas",
		"furries",
		"fursona",
		"fursonas",
	}) {
		// Add a reaction to the message
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🦊")
		if err != nil {
			log.Println(err)
		}

		score := database.UpdateFurryScore(m.Author.ID, 1)

		if score > 20 {
			// If the user has a furry score greater than 20, add a role to the user
			_, err = s.ChannelMessageSendReply(m.ChannelID, "You went to RainFurrest didn't you?", m.Reference())
			if err != nil {
				log.Println(err)
				return
			}
		} else if score > 10 {
			// If the user has a furry score greater than 10, add a role to the user
			err = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "1353376852185710696")
			if err != nil {
				log.Println(err)
				return
			}

		} else if score > 5 {
			// reply to the user
			_, err = s.ChannelMessageSendReply(m.ChannelID, "Naughty Kitty", m.Reference())
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
