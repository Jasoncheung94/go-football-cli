package competitions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jasoncheung94/go-football-cli/entity"
)

func Fetch() entity.CompetitionData {
	// http://api.football-data.org/v2/competitions

	file, err := ioutil.ReadFile("competitions.json") // need to fix the path for this and have a unique place.
	data := entity.CompetitionData{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("failed to process competition data")
		return data
	}

	return data
}
