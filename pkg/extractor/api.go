package extractor

type (
	// PlayersList represent the payload that will be collected from the API response
	PlayersList struct {
		Data struct {
			Team struct {
				Players []Player `json:"players"`
			} `json:"team"`
		} `json:"data"`
	}

	// PlayerTeams represents the payload that will be collected from the API response
	PlayerTeams struct {
		Data struct {
			Info struct {
				NationalTeam struct {
					Name string `json:"name"`
				} `json:"nationalTeam"`

				ClubTeam struct {
					Name string `json:"name"`
				} `json:"clubTeam"`
			} `json:"info"`
		} `json:"data"`
	}

	// Player individual player representation
	Player struct {
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       string `json:"age"`
		FullName  string
		Teams     []string
	}
)
