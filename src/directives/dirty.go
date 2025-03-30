package directives

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/database"
	"github.com/kochie/discord-bot-police/src/util"
)

func DirtyDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if the message contains any of the dirty words
	dirtyWords := []string{"dirty", "filthy", "naughty", "cum"}
	if util.CompareList(processedString, dirtyWords) {
		database.UpdateDirtyScore(m.Author.ID, 1)
	}
}
