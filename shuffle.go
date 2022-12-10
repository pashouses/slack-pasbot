package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/slack-go/slack"
)

func handleShuffle(sClient SlackClientIFace, s slack.SlashCommand) string {
	params := strings.Split(s.Text, " ")
	subteam := parseSubteam(params[0])
	if subteam.ID == "" {
		return "Incorrect format, should be /shuffle @team-name [1 = include self] [number]"
	}
	teamMemberIDs, err := sClient.GetUserGroupMembers(subteam.ID)
	if err != nil {
		errText := fmt.Sprintf("Error: %v\n", err)
		fmt.Print(errText)
		return errText
	}
	fmt.Printf("Members of %s: %v\n", subteam.Name, teamMemberIDs)
	includeSelf := len(params) > 1 && params[1] == "1"
	idsToBeShuffled := []string{}
	for _, mID := range teamMemberIDs {
		if includeSelf || mID != s.UserID {
			idsToBeShuffled = append(idsToBeShuffled, mID)
		}
	}
	if len(idsToBeShuffled) == 0 {
		return fmt.Sprintf("Group %s does not have any members", subteam.Name)
	}
	shuffle(idsToBeShuffled)
	fmt.Printf("Shuffle result: %v\n", idsToBeShuffled)
	number := 1
	if len(params) == 3 {
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
	return fmt.Sprintf("%s from %s, nominated by %s", selectedStr, subteam.Name, s.UserName)
}
