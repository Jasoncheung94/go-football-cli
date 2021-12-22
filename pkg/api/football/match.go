package football

import (
	"github.com/jasoncheung94/go-football-cli/entity"
	"github.com/jasoncheung94/go-football-cli/pkg/client"
)

func FetchMatches(competitionIDs []string) entity.MatchData {
	result := entity.MatchData{}
	f := &client.Filters{
		Competitions: competitionIDs,
	}
	client.RequestData(
		"http://api.football-data.org/v2/matches",
		&result,
		f,
	)
	return result
}
