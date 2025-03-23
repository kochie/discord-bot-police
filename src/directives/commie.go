package directives

import (
	"github.com/bwmarrin/discordgo"
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

		err := s.MessageReactionAdd(m.ChannelID, m.ID, "ussr:695565659303706745")

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
		if err != nil {
			log.Println(err)
		}
	}
}
