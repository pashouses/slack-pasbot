package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/slack-go/slack"
)

var frontendReviewMembers = []string{"U034TPRC8JV", "U039001AU49", "U03FVQUP8NB"}

func TestHandleShuffleDoNotIncludeSelf(t *testing.T) {
	sClient := &SlackClientMock{}
	teamName := "@frontend-review"
	s := slack.SlashCommand{
		/*
			/shuffle @frontend-review
			# shuffle from @frontend review
			# Default behavior: do not include self (U034TPRC8JV|@sendia)
			# and return 1 member
		*/
		Command:  "/shuffle",
		Text:     fmt.Sprintf("<!subteam^S03LFDE08N6|%s>", teamName),
		UserName: "@asendia",
		UserID:   "U034TPRC8JV",
	}
	resTxt := handleShuffle(sClient, s)
	if strings.Contains(resTxt, "U034TPRC8JV") {
		t.Error("U034TPRC8JV (self) should not exist in the response text")
	}
	if !(strings.Contains(resTxt, "<@U039001AU49>") || strings.Contains(resTxt, "<@U03FVQUP8NB>")) {
		t.Error("<@U039001AU49> or <@U03FVQUP8NB> should be exist in the response string")
	}
	if !strings.Contains(resTxt, fmt.Sprintf("from %s, nominated by %s", teamName, s.UserName)) {
		t.Error("Response should contain group name & caller username")
	}
}

func TestHandleShuffleDoNotIncludeSelfAndReturn3members(t *testing.T) {
	sClient := &SlackClientMock{}
	teamName := "@frontend-review"
	s := slack.SlashCommand{
		/*
			/shuffle @frontend-review 0 3
			# shuffle from @frontend review, do not include self (U034TPRC8JV|@sendia)
			# and return at max 3 members
		*/
		Command:  "/shuffle",
		Text:     fmt.Sprintf("<!subteam^S03LFDE08N6|%s> 0 3", teamName),
		UserName: "@asendia",
		UserID:   "U034TPRC8JV",
	}
	resTxt := handleShuffle(sClient, s)
	if strings.Contains(resTxt, "U034TPRC8JV") {
		t.Error("U034TPRC8JV (self) should not exist in the response text")
	}
	if !strings.Contains(resTxt, "<@U039001AU49>") || !strings.Contains(resTxt, "<@U03FVQUP8NB>") {
		t.Error("<@U039001AU49> & <@U03FVQUP8NB> should be exist in the response string")
	}
	if !strings.Contains(resTxt, fmt.Sprintf("from %s, nominated by %s", teamName, s.UserName)) {
		t.Error("Response should contain group name & caller username")
	}
}

func TestHandleShuffleIncludeSelfAndReturn3members(t *testing.T) {
	sClient := &SlackClientMock{}
	teamName := "@frontend-review"
	s := slack.SlashCommand{
		/*
			/shuffle @frontend-review 1 3
			# shuffle from @frontend review, include self (U034TPRC8JV|@sendia)
			# and return at max 3 members
		*/
		Command:  "/shuffle",
		Text:     fmt.Sprintf("<!subteam^S03LFDE08N6|%s> 1 3", teamName),
		UserName: "@asendia",
		UserID:   "U034TPRC8JV",
	}
	resTxt := handleShuffle(sClient, s)
	if !strings.Contains(resTxt, "U034TPRC8JV") {
		t.Error("U034TPRC8JV (self) should exist in the response text")
	}
	if !strings.Contains(resTxt, "<@U039001AU49>") || !strings.Contains(resTxt, "<@U03FVQUP8NB>") {
		t.Error("<@U039001AU49> & <@U03FVQUP8NB> should be exist in the response string")
	}
	if !strings.Contains(resTxt, fmt.Sprintf("from %s, nominated by %s", teamName, s.UserName)) {
		t.Error("Response should contain group name & caller username")
	}
}

func TestHandleShuffleHere(t *testing.T) {
	sClient := &SlackClientMock{}
	s := slack.SlashCommand{
		/*
			/shuffle @frontend-review 1 3
			# shuffle from @frontend review, include self (U034TPRC8JV|@sendia)
			# and return at max 3 members
		*/
		Command:   "/shuffle",
		Text:      "<!here|@here>",
		UserName:  "@asendia",
		UserID:    "U034TPRC8JV",
		ChannelID: "CHANNELID",
	}
	resTxt := handleShuffle(sClient, s)
	if strings.Contains(resTxt, "<@U034TPRC8JV>") {
		t.Error("Should not contain self")
	}
	if !(strings.Contains(resTxt, "<@U039001AU49>") || strings.Contains(resTxt, "<@U03FVQUP8NB>")) {
		t.Error("<@U034TPRC8JV> or <@U039001AU49> or <@U03FVQUP8NB> should be exist in the response string")
	}
	if !strings.Contains(resTxt, fmt.Sprintf("from %s, nominated by %s", "here", s.UserName)) {
		t.Error("Response should contain group name & caller username")
	}
}

func TestHandleShuffleHereAndSelf(t *testing.T) {
	sClient := &SlackClientMock{}
	s := slack.SlashCommand{
		/*
			/shuffle @frontend-review 1 3
			# shuffle from @frontend review, include self (U034TPRC8JV|@sendia)
			# and return at max 3 members
		*/
		Command:   "/shuffle",
		Text:      "<!here|@here> 1",
		UserName:  "@asendia",
		UserID:    "U034TPRC8JV",
		ChannelID: "CHANNELID",
	}
	resTxt := handleShuffle(sClient, s)
	if !(strings.Contains(resTxt, "<@U034TPRC8JV>") || strings.Contains(resTxt, "<@U039001AU49>") || strings.Contains(resTxt, "<@U03FVQUP8NB>")) {
		t.Error("<@U034TPRC8JV> or <@U039001AU49> or <@U03FVQUP8NB> should be exist in the response string")
	}
	if !strings.Contains(resTxt, fmt.Sprintf("from %s, nominated by %s", "here", s.UserName)) {
		t.Error("Response should contain group name & caller username")
	}
}

type SlackClientMock struct{}

func (m *SlackClientMock) GetUserGroupMembers(userGroup string) ([]string, error) {
	return frontendReviewMembers, nil
}

func (m *SlackClientMock) GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
	return frontendReviewMembers, "", nil
}
