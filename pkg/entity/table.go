package entity

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

type TableData struct {
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
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Standings []struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				CrestURL string `json:"crestUrl"`
			} `json:"team"`
			PlayedGames    int         `json:"playedGames"`
			Form           interface{} `json:"form"`
			Won            int         `json:"won"`
			Draw           int         `json:"draw"`
			Lost           int         `json:"lost"`
			Points         int         `json:"points"`
			GoalsFor       int         `json:"goalsFor"`
			GoalsAgainst   int         `json:"goalsAgainst"`
			GoalDifference int         `json:"goalDifference"`
		} `json:"table"`
	} `json:"standings"`
}

func (t TableData) Display() {
	fmt.Println(t.Competition.Name, "-", t.Season.StartDate[:4], t.Season.EndDate[:4])
	for _, standing := range t.Standings {
		fmt.Printf("%-4s|%-30s|%-4s|%-4s|%-4s|%-4s|%-4s|%-4s|%-4s|%-4s \n", "Pos", "Club", "MP", "W", "D", "L", "GF", "GA", "GD", "Pts")
		for _, team := range standing.Table {
			fmt.Printf("%-4d|%-30s|%-4d|%-4d|%-4d|%-4d|%-4d|%-4d|%-4d|%-4d \n",
				team.Position, team.Team.Name, team.PlayedGames, team.Won, team.Draw, team.Lost,
				team.GoalsFor, team.GoalsAgainst, team.GoalDifference, team.Points)
		}
	}
}

func (t TableData) DisplayWithPterm() {
	pterm.Println() // Blank line
	data := pterm.TableData{}
	for _, standing := range t.Standings {
		if len(standing.Table) == 0 {
			continue
		}
		if standing.Group != nil {
			data = append(data, []string{standing.Group.(string)})
		}
		data = append(data, []string{"Pos", "Club", "MP", "W", "D", "L", "GF", "GA", "GD", "Pts"})
		for _, team := range standing.Table {
			data = append(data, []string{
				strconv.FormatInt(int64(team.Position), 10),
				team.Team.Name,
				strconv.FormatInt(int64(team.PlayedGames), 10),
				strconv.FormatInt(int64(team.Draw), 10),
				strconv.FormatInt(int64(team.Won), 10),
				strconv.FormatInt(int64(team.Lost), 10),
				strconv.FormatInt(int64(team.GoalsFor), 10),
				strconv.FormatInt(int64(team.GoalsAgainst), 10),
				strconv.FormatInt(int64(team.GoalDifference), 10),
				strconv.FormatInt(int64(team.Points), 10),
			})
		}
		data = append(data, []string{"\n"})
	}
	pterm.DefaultTable.WithHasHeader().WithData(data).WithLeftAlignment().Render()
}
