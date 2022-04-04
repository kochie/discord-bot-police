package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	// "time"

	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

// type UserWarnings struct {
// 	Username    string    `json:username`
// 	Warnings    int       `json:warnings`
// 	LastWarning time.Time `json:last_warning`
// }

var reg *regexp.Regexp
var reg1 *regexp.Regexp

var rek *rekognition.Client

func init() {
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	r1, err := regexp.Compile("^\\d+$")
	if err != nil {
		log.Fatal(err)
	}
	reg = r
	reg1 = r1

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	rek = rekognition.NewFromConfig(cfg)
	mutex = &sync.Mutex{}
}

// func dafoe(dg *discordgo.Session) {
// 	file, err := os.Open("assets/dafoe.gif")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = dg.ChannelMessageSendComplex(os.Getenv("DAFOE_CHANNEL_ID"), &discordgo.MessageSend{
// 		Files: []*discordgo.File{{
// 			Name:        "dafoe.gif",
// 			ContentType: "image/gif",
// 			Reader:      file,
// 		}},
// 	})
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

var simps = []string{
	"assets/simp1.gif",
	"assets/simp2.gif",
	"assets/simp1.jpg",
	"assets/simp2.jpg",
	"assets/simp3.jpg",
	"assets/simp4.jpg",
}

func getContentType(filename string) string {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".mp4":
		return "video/mp4"
	default:
		return ""
	}
}

func main() {
	fmt.Println("Hello, World,4")
	token := os.Getenv("DISCORD_TOKEN")

	fmt.Println("Hello", token)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// ticker := time.NewTicker(24 * time.Hour)
	// done := make(chan bool)

	// go func() {
	// 	dafoe(dg)
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case <-ticker.C:
	// 			dafoe(dg)
	// 		}
	// 	}
	// }()

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		return
	}
}

func compareList(phrase string, comparisionList []string) bool {
	for _, word := range comparisionList {
		if strings.Contains(phrase, word) {
			return true
		}
	}

	return false
}

var timer *time.Timer
var mutex *sync.Mutex

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	processedString := reg.ReplaceAllString(strings.ToLower(m.Content), "")

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

	if m.Author.ID == os.Getenv("SIMP_ID") {
		if strings.Contains(processedString, "caitlin") {
			filename := simps[rand.Intn(len(simps))]

			file, err := os.Open(filename)
			if err != nil {
				log.Println(err)
				return
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

	if m.ChannelID == os.Getenv("PIZZAGATE_ID") {
		//fmt.Println("HELLO THERE")
		if reg1.MatchString(m.Content) {
			taunt, err := strconv.ParseInt(processedString, 10, 32)
			if err != nil {
				log.Println(err)
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

				if dgv, ok := s.VoiceConnections[serverId]; ok {
					mutex.Lock()
					fmt.Println("Joined Voice chat to play", tauntFile.Name())
					timer.Reset(10 * time.Second)
					dgvoice.PlayAudioFile(dgv, "assets/taunts/"+tauntFile.Name(), make(chan bool))
					mutex.Unlock()
					fmt.Println("Spoke")
				} else {
					fmt.Println(taunt, "selected, joining", os.Getenv("FOW_ID"))
					mutex.Lock()
					dgv, err := s.ChannelVoiceJoin(serverId, os.Getenv("FOW_ID"), false, true)
					mutex.Unlock()
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("Joined Voice chat to play", tauntFile.Name())
					dgvoice.PlayAudioFile(dgv, "assets/taunts/"+tauntFile.Name(), make(chan bool))
					fmt.Println("Spoke")

					timer = time.AfterFunc(10*time.Second, func() {
						err = dgv.Disconnect()
						if err != nil {
							fmt.Println(err)
							return
						}
						dgv.Close()
					})
				}

			}
		}
	}

	if m.ChannelID == os.Getenv("BUSCEMI_ID") {
		isSteveBuscemi := false
		celebLength := 0
		unknownFacesLength := 0
		for _, atta := range m.Attachments {
			if strings.HasSuffix(atta.Filename, ".png") ||
				strings.HasSuffix(atta.Filename, ".jpeg") ||
				strings.HasSuffix(atta.Filename, ".jpg") {
				resp, err := http.Get(atta.URL)
				if err != nil {
					fmt.Println("Error retrieving the file, ", err)
					continue
				}
				img, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading the response, ", err)
					continue
				}
				//contentType := http.DetectContentType(img)
				//base64img := base64.StdEncoding.EncodeToString(img)
				//
				//log.Println(fmt.Sprintf("data:%s;base64,%s", contentType, base64img))
				result, err := rek.RecognizeCelebrities(context.Background(), &rekognition.RecognizeCelebritiesInput{
					Image: &types.Image{
						Bytes: img,
					},
				})
				if err != nil {
					log.Println(err)
					continue
				}
				celebLength += len(result.CelebrityFaces)
				unknownFacesLength += len(result.UnrecognizedFaces)

				for _, celeb := range result.CelebrityFaces {
					if *celeb.Name == "Steve Buscemi" {
						isSteveBuscemi = true
						break
					}
				}
			}
		}

		if len(m.Attachments) > 0 && (!isSteveBuscemi && celebLength > 0) {
			err := s.ChannelMessageDelete(m.ChannelID, m.ID)
			if err != nil {
				log.Println(err)
			}

			file, err := os.Open("assets/steve_1.jpg")
			if err != nil {
				log.Println(err)

			}
			_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Files: []*discordgo.File{{
					Name:        "steve.jpg",
					ContentType: "image/jpeg",
					Reader:      file,
				}},
				Content: "Mr Buscemi is not happy with you.",
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
