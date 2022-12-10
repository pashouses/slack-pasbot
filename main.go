package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	token := os.Getenv("SLACK_BOT_TOKEN")
	sClient := slack.New(token)
	http.HandleFunc("/api/slack", func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.Body = io.NopCloser(io.TeeReader(r.Body, &verifier))
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = verifier.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Printf("Command: %s %s\n", s.Command, s.Text)
		switch s.Command {
		case "/shuffle":
			resTxt := handleShuffle(sClient, s)
			respondToSlack(w, resTxt, s.ChannelID)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)
}
