package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

type SlackClientIFace interface {
	GetUserGroupMembers(userGroup string) ([]string, error)
	GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error)
	GetUsersInfo(users ...string) (*[]slack.User, error)
}

// Format <!Name^VALUE|Text>
// Example:
// <!subteam^TEAMID|@team>
// <!here>
// <!channel|@channel>
type SlackTag struct {
	Name  string
	Value string
	Text  string
}

func parseSlackTag(txt string) SlackTag {
	tag := SlackTag{}
	if strings.HasPrefix(txt, "@") {
		tag.Name = strings.Replace(txt, "@", "", 1)
		return tag
	}
	r := regexp.MustCompile(`^<!(?P<tagname>\w+)(?P<tagvalue>\^[A-Z0-9]+)?(?P<text>\|@?.+)?>$`)
	submatch := r.FindStringSubmatch(txt)
	if len(submatch) == 0 {
		return tag
	}
	tag.Name = r.ReplaceAllString(txt, "${tagname}")
	tag.Value = strings.Replace(r.ReplaceAllString(txt, "${tagvalue}"), "^", "", 1)
	tag.Text = strings.Replace(r.ReplaceAllString(txt, "${text}"), "|", "", 1)
	return tag
}

func respondToSlack(w http.ResponseWriter, msg string, sClient slack.Client, s slack.SlashCommand) {
	channelID, timestamp, err := sClient.PostMessage(s.ChannelID, slack.MsgOptionText(msg, false))
	fmt.Printf("Slack message is sent successfully to %s at %s: %s", channelID, timestamp, msg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	w.WriteHeader(200)
}
