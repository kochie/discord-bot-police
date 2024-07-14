package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/bwmarrin/discordgo"
)

func celebDetection(s *discordgo.Session, m *discordgo.MessageCreate) {
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
