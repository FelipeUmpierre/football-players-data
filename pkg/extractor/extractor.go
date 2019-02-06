package extractor

import (
	"sort"
	"sync"
)

type (
	configuration interface {
		TeamsURL() string
		PlayerURL() string
	}

	// Players represent a slice of Player
	Players []Player
)

func (p Players) Len() int           { return len(p) }
func (p Players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Players) Less(i, j int) bool { return p[i].FullName < p[j].FullName }

// ExtractPlayers responsible to make the HTTP calls and extract the content from the API response
// to structs, it execute two goroutines, one for collecting concurrently all the data from all the
// teams in the list, and another, to collect the players and append it to the Players type
func ExtractPlayers(config configuration, list []int) (Players, <-chan error) {
	playerCh := make(chan Player)
	errCh := make(chan error)
	defer close(playerCh)
	defer close(errCh)

	var (
		players Players
		mu      sync.RWMutex
	)

	go func() {
		for {
			select {
			case player := <-playerCh:
				mu.Lock()
				players = append(players, player)
				mu.Unlock()
			}
		}
	}()

	wg := new(sync.WaitGroup)
	for _, id := range list {
		wg.Add(1)
		go extractTeamContent(wg, config, id, playerCh, errCh)
	}

	wg.Wait()

	mu.Lock()
	players = filterPlayers(players)
	sort.Sort(players)
	mu.Unlock()

	return players, errCh
}

// filterPlayers will remove the redundancy of players that might appear
func filterPlayers(originalPlayers Players) Players {
	filterPlayersMap := make(map[string]Player, len(originalPlayers))
	for _, player := range originalPlayers {
		filterPlayersMap[player.FullName] = player
	}

	players := make(Players, 0, len(filterPlayersMap))
	for _, player := range filterPlayersMap {
		players = append(players, player)
	}

	return players
}

// extractTeamContent responsible to extract the data from the team api url and call the
// extractPlayerContent to collect the individual data from each player from the team
func extractTeamContent(
	wg *sync.WaitGroup,
	config configuration,
	id int,
	playerCh chan Player,
	errCh chan error,
) {
	defer wg.Done()

	teamPlayers, err := collectDataFromTeam(id, config.TeamsURL())
	if err != nil {
		errCh <- err
	}

	var playerWg sync.WaitGroup
	for _, player := range teamPlayers {
		playerWg.Add(1)
		go extractPlayerContent(&playerWg, player, config.PlayerURL(), playerCh, errCh)
	}
	playerWg.Wait()
}

// extractPlayerContent responsible to extract the teams from the player.
func extractPlayerContent(
	wg *sync.WaitGroup,
	player Player,
	url string,
	playerCh chan Player,
	errCh chan error,
) {
	defer wg.Done()

	p, err := collectDataFromPlayer(player, url)
	if err != nil {
		errCh <- err
	}

	playerCh <- p
}
