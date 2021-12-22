package entity

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

type Scorers struct {
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
		ID              int      `json:"id"`
		StartDate       string   `json:"startDate"`
		EndDate         string   `json:"endDate"`
		CurrentMatchday int      `json:"currentMatchday"`
		AvailableStages []string `json:"availableStages"`
	} `json:"season"`
	Scorers []struct {
		Player struct {
			ID             int       `json:"id"`
			FirstName      string    `json:"firstName"`
			LastName       string    `json:"lastName"`
			DateOfBirth    string    `json:"dateOfBirth"`
			CountryOfBirth string    `json:"countryOfBirth"`
			Nationality    string    `json:"nationality"`
			Position       string    `json:"position"`
			LastUpdated    time.Time `json:"lastUpdated"`
		} `json:"player"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		NumberOfGoals int `json:"numberOfGoals"`
	} `json:"scorers"`
}

func (s Scorers) Display() {
	fmt.Println("Competition: ", s.Competition.Name)
	fmt.Printf("%-40s | %s\n", "Name", "Goals")
	for _, scorer := range s.Scorers {
		fullName := scorer.Player.FirstName + " " + scorer.Player.LastName
		fmt.Printf("%-40s %4d\n", fullName, scorer.NumberOfGoals)
	}
}

func (s Scorers) DisplayWithPterm() {
	data := pterm.TableData{[]string{"Competition:", s.Competition.Name}}
	for _, scorer := range s.Scorers {
		fullName := scorer.Player.FirstName + " " + scorer.Player.LastName
		data = append(data, []string{fullName, strconv.FormatInt(int64(scorer.NumberOfGoals), 10)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(data).WithLeftAlignment().Render()

}
