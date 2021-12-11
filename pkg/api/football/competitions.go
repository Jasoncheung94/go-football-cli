package football

import (
	"github.com/jasoncheung94/go-football-cli/entity"
	"github.com/jasoncheung94/go-football-cli/pkg/client"
)

func FetchCompetitions() entity.CompetitionData {
	result := entity.CompetitionData{}
	client.RequestData("http://api.football-data.org/v2/competitions", &result, nil)
	return result
}
