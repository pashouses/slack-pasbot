package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/slack-go/slack"
)

func handleShuffle(sClient SlackClientIFace, s slack.SlashCommand) string {
	params := strings.Split(s.Text, " ")
	tag := parseSlackTag(params[0])
	var source string
	var userIDs []string
	if tag.Name == "subteam" && tag.Value != "" && tag.Text != "" {
		// Subteam
		source = tag.Text
		var err error
		userIDs, err = sClient.GetUserGroupMembers(tag.Value)
		if err != nil {
			errText := fmt.Sprintf("Error: %v\n", err)
			fmt.Print(errText)
			return errText
		}
	} else if tag.Name == "here" || tag.Name == "channel" {
		// Current channel
		source = tag.Name
		var err error
		params := slack.GetUsersInConversationParameters{ChannelID: s.ChannelID}
		userIDs, _, err = sClient.GetUsersInConversation(&params)
		if err != nil {
			errText := fmt.Sprintf("Error: %v\n", err)
			fmt.Print(errText)
			return errText
		}
	} else {
		return "Incorrect format, should be /shuffle @team-name [1 = include self] [number]"
	}
	fmt.Printf("Members of %s: %v\n", source, userIDs)
	includeSelf := len(params) > 1 && params[1] == "1"
	idsToBeShuffled := []string{}
	for _, mID := range userIDs {
		if includeSelf || mID != s.UserID {
			idsToBeShuffled = append(idsToBeShuffled, mID)
		}
	}
	if len(idsToBeShuffled) == 0 {
		return fmt.Sprintf("Source %s does not have any members", source)
	}
	shuffle(idsToBeShuffled)
	fmt.Printf("Shuffle result: %v\n", idsToBeShuffled)
	number := 1
	if len(params) == 3 {
		var err error
		number, err = strconv.Atoi(params[2])
		if err != nil {
			number = 1
		}
		if number < 1 {
			number = 1
		} else if number > len(idsToBeShuffled) {
			number = len(idsToBeShuffled)
		}
	}
	selectedIDs := []string{}
	for _, userID := range idsToBeShuffled[:number] {
		selectedIDs = append(selectedIDs, fmt.Sprintf("<@%s>", userID))
	}
	selectedStr := strings.Join(selectedIDs, ", ")
	fmt.Printf("Selected member(s): %s\n", selectedStr)
	return fmt.Sprintf("%s from %s, nominated by %s", selectedStr, source, s.UserName)
}
