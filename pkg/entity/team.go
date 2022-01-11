package entity

import (
	"fmt"
	"time"
)

type TeamData struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday interface{} `json:"currentMatchday"`
		AvailableStages []string    `json:"availableStages"`
	} `json:"season"`
	Teams []struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		ShortName   string    `json:"shortName"`
		Tla         string    `json:"tla"`
		CrestURL    string    `json:"crestUrl"`
		Address     string    `json:"address"`
		Phone       string    `json:"phone"`
		Website     string    `json:"website"`
		Email       string    `json:"email"`
		Founded     int       `json:"founded"`
		ClubColors  string    `json:"clubColors"`
		Venue       string    `json:"venue"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"teams"`
}

func (t TeamData) Display() {
	fmt.Println("Competition:", t.Competition.Name)
	fmt.Printf("%-4s|%-30s\n", "ID", "Name")
	for _, team := range t.Teams {
		fmt.Printf("%-4d|%-30s|\n", team.ID, team.Name)
	}
}

func (t TeamData) DisplayWithPterm() {
	t.Display()
	// TODO - Add proper formatting
}
