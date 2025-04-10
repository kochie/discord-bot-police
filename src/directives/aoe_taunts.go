package directives

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var reg1 *regexp.Regexp
var mutex *sync.Mutex
var timer *time.Timer

func init() {
	r1, err := regexp.Compile(`^\d+$`)
	if err != nil {
		log.Fatal(err)
	}
	reg1 = r1

	mutex = &sync.Mutex{}
}

func AoeTaunts(processedString string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID == os.Getenv("PIZZAGATE_ID") {

		if reg1.MatchString(m.Content) {
			taunt, err := strconv.ParseInt(processedString, 10, 32)
			if err != nil {
				return
			}
			entries, err := os.ReadDir("assets/taunts")
			if err != nil {
				log.Println(err)
			}

			var tauntFile os.DirEntry
			for _, file := range entries {
				if strings.HasPrefix(file.Name(), processedString+"_") {
					tauntFile = file
					break
				}
			}

			if tauntFile != nil {
				serverId := os.Getenv("SERVER_ID")

				mutex.Lock()
				if dgv, ok := s.VoiceConnections[serverId]; ok {
					log.Println("Joined Voice chat to play", tauntFile.Name())
					timer.Reset(10 * time.Second)
					dgvoice.PlayAudioFile(dgv, "assets/taunts/"+tauntFile.Name(), make(chan bool))
					log.Println("Spoke")
				} else {
					fmt.Println(taunt, "selected, joining", os.Getenv("FOW_ID"))
					dgv, err := s.ChannelVoiceJoin(serverId, os.Getenv("FOW_ID"), false, true)

					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Joined Voice chat to play", tauntFile.Name())
					dgvoice.PlayAudioFile(dgv, "assets/taunts/"+tauntFile.Name(), make(chan bool))
					fmt.Println("Spoke")

					timer = time.AfterFunc(10*time.Second, func() {
						err = dgv.Disconnect()
						if err != nil {
							fmt.Println(err)
						}
						dgv.Close()
					})
				}
				mutex.Unlock()

			} else {
				log.Printf("Taunt not found")
			}
		}
	}
}
