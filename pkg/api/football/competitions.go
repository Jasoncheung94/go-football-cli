package football

import (
	"github.com/jasoncheung94/go-football-cli/pkg/client"
	"github.com/jasoncheung94/go-football-cli/pkg/entity"
)

func FetchCompetitions() entity.CompetitionData {
	result := entity.CompetitionData{}
	f := &client.Filters{
		CacheFile: "competitions.json",
	}
	client.RequestData("http://api.football-data.org/v2/competitions", &result, f)
	return result
}
