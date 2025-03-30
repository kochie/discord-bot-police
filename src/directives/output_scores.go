package directives

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kochie/discord-bot-police/src/database"
	"log"
	"strconv"
)

func OutputScores(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commies := database.GetAllCommieScores()
	furries := database.GetAllFurryScores()
	dirty := database.GetDirtyScores()

	// Get the unique keys of users for both maps
	userIds := make([]string, 1)
	for id, _ := range commies {
		userIds = append(userIds, id)
	}
	for id, _ := range furries {
		// Check if the user is already in the list
		found := false
		for _, userId := range userIds {
			if userId == id {
				found = true
				break
			}
		}
		if !found {
			userIds = append(userIds, id)
		}
	}
	for id, _ := range dirty {
		// Check if the user is already in the list
		found := false
		for _, userId := range userIds {
			if userId == id {
				found = true
				break
			}
		}
		if !found {
			userIds = append(userIds, id)
		}
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Name", "Commie\nScore", "Furry\nScore", "Cum\nScore", "Total Degeneracy\nScore", "JWDS\nScore"})
	for _, userId := range userIds {
		st, err := s.User(userId)
		if err != nil {
			// If the user is not found, skip this user
			continue
		}

		commieScore := int64(0)
		if score, ok := commies[userId]; ok {
			commieScore, err = strconv.ParseInt(score, 10, 64)
		}

		furryScore := int64(0)
		if score, ok := furries[userId]; ok {
			furryScore, err = strconv.ParseInt(score, 10, 64)
		}

		dirtyScore := int64(0)
		if score, ok := dirty[userId]; ok {
			dirtyScore, err = strconv.ParseInt(score, 10, 64)
		}

		jwds := 2*commieScore + 3*furryScore + dirtyScore
		t.AppendRow(table.Row{st.Username, commieScore, furryScore, dirtyScore, commieScore + furryScore + dirtyScore, jwds})
	}

	t.SortBy([]table.SortBy{
		{Name: "Total Degeneracy Score", Mode: table.DscNumeric},
	})
	
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("```%s```", t.Render()),
		},
	})
	if err != nil {
		// Handle error
		log.Println(err)
	}
}
