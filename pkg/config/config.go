package config

import "os"

// Configuration holds the data available in the env variables
type Configuration struct {
	teamsURL  string
	playerURL string
}

// New collects the env variables and add to the configuration struct
func New() *Configuration {
	return &Configuration{
		teamsURL:  os.Getenv("TEAMS_URL"),
		playerURL: os.Getenv("PLAYER_URL"),
	}
}

// TeamsURL returns the url for the teams data
func (c Configuration) TeamsURL() string { return c.teamsURL }

// PlayerURL returns the url for the player data
func (c Configuration) PlayerURL() string { return c.playerURL }
