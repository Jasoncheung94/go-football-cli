package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/entity"
	"github.com/jasoncheung94/go-football-cli/pkg/client"
)

func FetchMatches(competitionCode string) entity.MatchData {
	if competitionCode == "" {
		competitionCode = "PL"
	}
	result := entity.MatchData{}
	f := &client.Filters{
		Competitions: "2021",
	}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/matches", competitionCode),
		&result,
		f,
	)
	return result
}
