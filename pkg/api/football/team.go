package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/client"
	"github.com/jasoncheung94/go-football-cli/pkg/entity"
)

func FetchTeams(competitionCode string) entity.TeamData {
	if competitionCode == "" {
		competitionCode = "PL"
	}
	result := entity.TeamData{}
	f := &client.Filters{
		CacheFile: "teams.json",
	}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/teams", competitionCode),
		&result,
		f,
	)
	return result
}
