package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func commieDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if compareList(processedString, []string{
		"communist",
		"commie",
		"communism",
		"redscare",
		"lenin",
		"marx",
		"stalin",
		"kermitthefrog",
	}) {
		filename := "assets/Joseph_McCarthy.jpg"

		if m.Author.ID == os.Getenv("KNOWN_COMMUNIST_ID") {
			filename = "assets/commie.mp4"
		}

		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
		}
		_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Files: []*discordgo.File{{
				Name:        filepath.Base(filename),
				ContentType: getContentType(filename),
				Reader:      file,
			}},
		})
		if err != nil {
			log.Println(err)
		}
	}
}
