package directives

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kochie/discord-bot-police/src/util"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

var simps = []string{
	"assets/simp1.gif",
	"assets/simp2.gif",
	"assets/simp1.jpg",
	"assets/simp2.jpg",
	"assets/simp3.jpg",
	"assets/simp4.jpg",
}

func SimpDetection(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) bool {
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
					ContentType: util.GetContentType(filename),
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
