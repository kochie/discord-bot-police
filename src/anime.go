package main

import (
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func animeDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(processedString, "anime") {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "😡")
		if err != nil {
			log.Println(err)
		}

		filename := ""
		if rand.Float64() > 0.5 {
			filename = "assets/anime_ban_1.jpg"
		} else {
			filename = "assets/anime_ban_2.jpg"
		}
		file, err := os.Open(filename)

		if err != nil {
			log.Println(err)
		}
		_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Files: []*discordgo.File{{
				Name:        "anime.jpg",
				ContentType: "image/jpeg",
				Reader:      file,
			}},
		})
		if err != nil {
			log.Println(err)
		}
		file.Close()
	}
}
