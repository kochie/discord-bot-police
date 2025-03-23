package directives

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/database"
	"github.com/kochie/discord-bot-police/src/util"
	"log"
)

func CommieDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if util.CompareList(processedString, []string{
		"communist",
		"commie",
		"communism",
		"redscare",
		"lenin",
		"marx",
		"stalin",
		"kermitthefrog",
	}) {
		//filename := "assets/Joseph_McCarthy.jpg"

		score := database.UpdateCommieScore(m.Author.ID, 1)
		if score > 20 {
			_, err := s.ChannelMessageSendReply(m.ChannelID, "Known communist sympathiser found", m.Reference())
			if err != nil {
				log.Println(err)
				return
			}
		}

		err := s.MessageReactionAdd(m.ChannelID, m.ID, "ussr:906835494598496306")
		if err != nil {
			log.Println(err)
			return
		}

		//if m.Author.ID == os.Getenv("KNOWN_COMMUNIST_ID") {
		//	filename = "assets/commie.mp4"
		//}
		//
		//file, err := os.Open(filename)
		//if err != nil {
		//	log.Println(err)
		//}
		//_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		//	Files: []*discordgo.File{{
		//		Name:        filepath.Base(filename),
		//		ContentType: util.GetContentType(filename),
		//		Reader:      file,
		//	}},
		//})
	}
}
