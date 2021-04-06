package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/bwmarrin/discordgo"
)

type UserWarnings struct {
	Username    string    `json:username`
	Warnings    int       `json:warnings`
	LastWarning time.Time `json:last_warning`
}

var reg *regexp.Regexp

func init() {
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	reg = r
}

func main() {
	token := os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

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
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	processedString := reg.ReplaceAllString(strings.ToLower(m.Content), "")

	if strings.Contains(processedString, "anime") {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "😡")
		file, err := os.Open("assets/anime_ban_1.jpg")
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
	}

	if m.ChannelID == "794074793388408832" {
		isSteveBuscemi := false
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
				contentType := http.DetectContentType(img)
				base64img := base64.StdEncoding.EncodeToString(img)

				result, err := rek.RecognizeCelebrities(context.Background(), &rekognition.RecognizeCelebritiesInput{
					Image: &types.Image{
						Bytes: []byte(fmt.Sprintf("data:%s;base64,%s", contentType, base64img)),
					},
				})
				if err != nil {
					log.Println(err)
					continue
				}

				for _, celeb := range result.CelebrityFaces {
					if *celeb.Name == "Steve Buscemi" {
						isSteveBuscemi = true
					}
				}
			}
		}

		if len(m.Attachments) > 0 && !isSteveBuscemi {
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

var rek *rekognition.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	rek = rekognition.NewFromConfig(cfg)
}
