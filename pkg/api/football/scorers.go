package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/client"
	"github.com/jasoncheung94/go-football-cli/pkg/entity"
)

func FetchScorers(competitionId string) entity.Scorers {
	result := entity.Scorers{}
	f := &client.Filters{
		CacheFile: "scorers.json",
	}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/scorers", competitionId),
		&result,
		f,
	)
	return result
}
