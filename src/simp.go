package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func simpDetaction(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Author.ID == os.Getenv("SIMP_ID") {
		if strings.Contains(processedString, "caitlin") {
			filename := simps[rand.Intn(len(simps))]

			file, err := os.Open(filename)
			if err != nil {
				log.Println(err)
				return true
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
	return false
}
