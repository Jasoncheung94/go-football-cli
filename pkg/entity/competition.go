package entity

import (
	"fmt"
	"time"
)

type CompetitionData struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competitions []struct {
		ID   int `json:"id"`
		Area struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			CountryCode string      `json:"countryCode"`
			EnsignURL   interface{} `json:"ensignUrl"`
		} `json:"area"`
		Name          string      `json:"name"`
		Code          string      `json:"code"`
		EmblemURL     interface{} `json:"emblemUrl"`
		Plan          string      `json:"plan"`
		CurrentSeason struct {
			ID              int         `json:"id"`
			StartDate       string      `json:"startDate"`
			EndDate         string      `json:"endDate"`
			CurrentMatchday int         `json:"currentMatchday"`
			Winner          interface{} `json:"winner"`
		} `json:"currentSeason"`
		NumberOfAvailableSeasons int       `json:"numberOfAvailableSeasons"`
		LastUpdated              time.Time `json:"lastUpdated"`
	} `json:"competitions"`
}

func (c CompetitionData) Display() {
	fmt.Println("Total Competitions:", c.Count)
	fmt.Printf("%4s |%-30s|%-5s %s\n", "ID", "Name", "Code", "Location")
	for _, competition := range c.Competitions {
		fmt.Printf("%-5d|%-30s|%-5s %s \n",
			competition.ID, competition.Name, competition.Code, competition.Area.Name)
	}
}
