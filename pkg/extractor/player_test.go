package extractor

import (
	"reflect"
	"strings"
	"testing"
)

const playerURL = "https://vintagemonster.onefootball.com/api/player/en/%s.json"

func TestCollectDataFromPlayer(t *testing.T) {
	tests := []struct {
		scenario string
		id       string
		teams    []string
	}{
		{
			scenario: "Search for an existing player",
			id:       "6672",
			teams:    []string{"Goztepe A.S."},
		},
		{
			scenario: "Search for an inexisting player",
			id:       "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			player, err := collectDataFromPlayer(Player{
				ID: tt.id,
			}, playerURL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tt.teams, player.Teams) {
				t.Errorf(
					"incorrect teams from player, expected %q, got %q",
					strings.Join(tt.teams, ", "),
					strings.Join(player.Teams, ", "),
				)
			}
		})
	}
}
