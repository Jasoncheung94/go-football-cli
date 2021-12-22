package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/entity"
	"github.com/jasoncheung94/go-football-cli/pkg/client"
)

func FetchScorers(competitionId string) entity.Scorers {
	result := entity.Scorers{}
	f := &client.Filters{}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/scorers", competitionId),
		&result,
		f,
	)
	return result
}
