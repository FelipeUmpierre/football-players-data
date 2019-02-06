package extractor

import "testing"

const teamURL = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"

func TestCollectDataFromTeam(t *testing.T) {
	tests := []struct {
		scenario     string
		id           int
		totalPlayers int
	}{
		{
			scenario:     "Search for Barcelona Team",
			id:           5,
			totalPlayers: 34,
		},
		{
			scenario:     "Search for Team ID invalid",
			id:           75,
			totalPlayers: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			players, err := collectDataFromTeam(tt.id, teamURL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(players) != tt.totalPlayers {
				t.Errorf("different number of players, expected %d, got %d", tt.totalPlayers, len(players))
			}
		})
	}
}
