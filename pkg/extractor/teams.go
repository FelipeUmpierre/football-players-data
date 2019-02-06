package extractor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// collectDataFromTeam will make the HTTP call to the team url and parse it to the PlayersList struct
// and it will return a slice of all the players found in the team
func collectDataFromTeam(id int, url string) ([]Player, error) {
	response, err := http.Get(fmt.Sprintf(url, id))
	if err != nil {
		return nil, err
	}

	var playersList PlayersList
	if err := json.NewDecoder(response.Body).Decode(&playersList); err != nil {
		return nil, err
	}

	return playersList.Data.Team.Players, nil
}
