package entity

import (
	"fmt"
	"time"
)

type MatchData struct {
	Count   int `json:"count"`
	Filters struct {
		DateFrom string `json:"dateFrom"`
		DateTo   string `json:"dateTo"`
	} `json:"filters"`
	Matches []struct {
		ID          int `json:"id"`
		Competition struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"competition"`
		Season struct {
			ID              int      `json:"id"`
			StartDate       string   `json:"startDate"`
			EndDate         string   `json:"endDate"`
			CurrentMatchday int      `json:"currentMatchday"`
			AvailableStages []string `json:"availableStages"`
		} `json:"season"`
		UtcDate     time.Time   `json:"utcDate"`
		Status      string      `json:"status"`
		Attendance  interface{} `json:"attendance"`
		Matchday    int         `json:"matchday"`
		Stage       string      `json:"stage"`
		Group       string      `json:"group"`
		LastUpdated time.Time   `json:"lastUpdated"`
		HomeTeam    struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Coach struct {
				ID             int    `json:"id"`
				Name           string `json:"name"`
				CountryOfBirth string `json:"countryOfBirth"`
				Nationality    string `json:"nationality"`
			} `json:"coach"`
			Captain struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"captain"`
			Lineup []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Position    string `json:"position"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"lineup"`
			Bench []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Position    string `json:"position"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"bench"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Coach struct {
				ID             int    `json:"id"`
				Name           string `json:"name"`
				CountryOfBirth string `json:"countryOfBirth"`
				Nationality    string `json:"nationality"`
			} `json:"coach"`
			Captain struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"captain"`
			Lineup []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Position    string `json:"position"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"lineup"`
			Bench []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Position    string `json:"position"`
				ShirtNumber int    `json:"shirtNumber"`
			} `json:"bench"`
		} `json:"awayTeam"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"halfTime"`
			ExtraTime struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"extraTime"`
			Penalties struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"penalties"`
		} `json:"score"`
		Goals []struct {
			Minute int    `json:"minute"`
			Type   string `json:"type"`
			Team   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			Scorer struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"scorer"`
			Assist struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"assist"`
		} `json:"goals"`
		Bookings []struct {
			Minute int `json:"minute"`
			Team   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			Player struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"player"`
			Card string `json:"card"`
		} `json:"bookings"`
		Substitutions []struct {
			Minute int `json:"minute"`
			Team   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			PlayerOut struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"playerOut"`
			PlayerIn struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"playerIn"`
		} `json:"substitutions"`
		Referees []struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Nationality interface{} `json:"nationality"`
		} `json:"referees"`
	} `json:"matches"`
}

func (m MatchData) Display() {
	renderedCompetition := ""
	for _, match := range m.Matches {
		if renderedCompetition != match.Competition.Name {
			fmt.Println("Competition:", match.Competition.Name)
			fmt.Printf("\n%-20s |%-30s      %-30s\n", "Datetime", "Home Team", "Away Team")
			renderedCompetition = match.Competition.Name
		}
		fmt.Printf("%-20s |%-30s vs   %-30s \n", match.UtcDate.Format("Jan 02 15:04 PM"), match.HomeTeam.Name, match.AwayTeam.Name)
	}
}
