package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/client"
	"github.com/jasoncheung94/go-football-cli/pkg/entity"
)

func FetchTable(competitionCode string) entity.TableData {
	if competitionCode == "" {
		competitionCode = "PL"
	}
	result := entity.TableData{}
	f := &client.Filters{
		CacheFile: "tables.json",
	}
	if competitionCode == "CL" {
		f.CacheFile = "championsleague.json"
	}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/standings", competitionCode),
		&result,
		f,
	)
	return result
}
