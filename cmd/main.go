package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/FelipeUmpierre/onefootball-test/pkg/config"
	"github.com/FelipeUmpierre/onefootball-test/pkg/extractor"
	"github.com/FelipeUmpierre/onefootball-test/pkg/teams"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.New()

	executionTime := time.Now()
	players, errCh := extractor.ExtractPlayers(conf, teams.TeamsList)

	go func() {
		for {
			select {
			case err := <-errCh:
				if err != nil {
					log.Println("an error happen:", err)
				}
			}
		}
	}()

	for i, player := range players {
		fmt.Printf("%d. %s; %s; %s\n", i+1, player.FullName, player.Age, strings.Join(player.Teams, ", "))
	}

	fmt.Println("Executed in", time.Since(executionTime))
}
