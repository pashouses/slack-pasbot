package main

import (
	"encoding/json"
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

func respondToSlack(w http.ResponseWriter, msg string, channelID string) {
	resp := &slack.Msg{
		Text: msg,
		// https://api.slack.com/interactivity/slash-commands#responding_immediate_response
		ResponseType: "in_channel",
	}
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
