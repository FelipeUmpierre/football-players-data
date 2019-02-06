package extractor

import (
	"os"
	"testing"

	"github.com/FelipeUmpierre/onefootball-test/pkg/config"
	"github.com/FelipeUmpierre/onefootball-test/pkg/teams"
)

func initEnvVars() {
	os.Setenv("TEAMS_URL", "https://vintagemonster.onefootball.com/api/teams/en/%d.json")
	os.Setenv("PLAYER_URL", "https://vintagemonster.onefootball.com/api/player/en/%s.json")
}

func TestExtractPlayers(t *testing.T) {
	initEnvVars()

	tests := []struct {
		scenario string
		function func(*testing.T, *config.Configuration)
	}{
		{
			scenario: "Success",
			function: testExtractPlayersSuccess,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			tt.function(t, config.New())
		})
	}
}

func testExtractPlayersSuccess(t *testing.T, c *config.Configuration) {
	players, errCh := ExtractPlayers(c, teams.TeamsList)

	go func() {
		for {
			err, ok := <-errCh
			if !ok {
				return
			}

			t.Errorf("failed to extract the players, %v", err)
		}
	}()

	if len(players) <= 0 {
		t.Errorf("expected to collect players, got %d", len(players))
	}
}
