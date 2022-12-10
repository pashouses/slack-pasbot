package main

import (
	"testing"
)

func TestParseSubteam(t *testing.T) {
	subteam := parseSubteam("<!subteam^S03LCFJ9HL5|@backend-review>")
	if subteam.ID != "S03LCFJ9HL5" || subteam.Name != "@backend-review" {
		t.Errorf("Failed to parse subteam: %v\n", subteam)
	}

	subteam2 := parseSubteam("@invalid-tag")
	if subteam2.ID != "" || subteam2.Name != "" {
		t.Errorf("Subteam should be invalid: %v\n", subteam2)
	}
}
