package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/slack-go/slack"
)

type SlackClientIFace interface {
	GetUserGroupMembers(userGroup string) ([]string, error)
}

type SlackSubteam struct {
	Name string
	ID   string
}

func parseSubteam(txt string) SlackSubteam {
	r := regexp.MustCompile(`^<!subteam\^(?P<teamid>[A-Z]\w+)\|(?P<teamname>@.+)>$`)
	subteam := SlackSubteam{}
	submatch := r.FindStringSubmatch(txt)
	if len(submatch) == 0 {
		return subteam
	}
	matchGroup := r.SubexpNames()
	if len(matchGroup) == 3 {
		subteam.ID = r.ReplaceAllString(txt, "${teamid}")
		subteam.Name = r.ReplaceAllString(txt, "${teamname}")
	}
	return subteam
}

func respondToSlack(w http.ResponseWriter, msg string, sClient slack.Client, s slack.SlashCommand) {
	channelID, timestamp, err := sClient.PostMessage(s.ChannelID, slack.MsgOptionText(msg, false))
	fmt.Printf("Slack message is sent successfully to %s at %s: %s", channelID, timestamp, msg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	w.WriteHeader(200)
}
