package competitions

import (
	"github.com/jasoncheung94/go-football-cli/entity"
	"github.com/jasoncheung94/go-football-cli/pkg/client"
)

func Fetch() entity.CompetitionData {
	result := entity.CompetitionData{}
	client.RequestData("http://api.football-data.org/v2/competitions", &result)
	return result
}
