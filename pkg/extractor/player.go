package extractor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// collectDataFromPlayer respnsible to make the HTTP call to the player url and parse it to the PlayerTeams struct.
// it will then, generate a new Player struct with the content found from the player.
func collectDataFromPlayer(player Player, url string) (Player, error) {
	response, err := http.Get(fmt.Sprintf(url, player.ID))
	if err != nil {
		return Player{}, err
	}

	var teams PlayerTeams
	if err := json.NewDecoder(response.Body).Decode(&teams); err != nil {
		return Player{}, err
	}

	return Player{
		ID:        player.ID,
		FirstName: player.FirstName,
		LastName:  player.LastName,
		FullName: strings.Join(sanitizeSlice([]string{
			player.FirstName,
			player.LastName,
		}), " "),
		Age: player.Age,
		Teams: sanitizeSlice([]string{
			teams.Data.Info.NationalTeam.Name,
			teams.Data.Info.ClubTeam.Name,
		}),
	}, nil
}

// sanitizeSlice remove any string that might be empty.
func sanitizeSlice(originalSlice []string) []string {
	var newSlice []string
	for _, s := range originalSlice {
		if s == "" {
			continue
		}

		newSlice = append(newSlice, s)
	}
	return newSlice
}
