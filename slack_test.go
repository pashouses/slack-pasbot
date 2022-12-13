package main

import (
	"testing"
)

func TestParseTag(t *testing.T) {
	tag := parseSlackTag("<!subteam^S03LCFJ9HL5|@backend-review>")
	if tag.Name != "subteam" || tag.Value != "S03LCFJ9HL5" || tag.Text != "@backend-review" {
		t.Errorf("Failed to parse tag: %v\n", tag)
	}

	tag = parseSlackTag("<!here|@here>")
	if tag.Name != "here" || tag.Value != "" || tag.Text != "@here" {
		t.Errorf("Failed to parse tag: %v\n", tag)
	}

	tag = parseSlackTag("<!here>")
	if tag.Name != "here" || tag.Value != "" || tag.Text != "" {
		t.Errorf("Failed to parse tag: %v\n", tag)
	}

	tag = parseSlackTag("@here")
	if tag.Name != "here" || tag.Value != "" || tag.Text != "" {
		t.Errorf("Failed to parse tag: %v\n", tag)
	}
}
