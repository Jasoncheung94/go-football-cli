package table

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/jasoncheung94/go-football-cli/entity"
)

func Fetch() entity.TableData {
	// http://api.football-data.org/v2/competitions

	file, err := ioutil.ReadFile("premier-league.json") // need to fix the path for this and have a unique place.
	data := entity.TableData{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("failed to process table data")
		return data
	}

	return data
	unix := time.Now().Unix()
	fmt.Println("creating file")
	f, err := os.Create(fmt.Sprintf("./%v", unix))
	if err != nil {
		panic("failed to create file to read into")
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, "http://api.football-data.org/v2/competitions", nil)
	if err != nil {
		fmt.Println("failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("failed to do request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("failed body")
	}
	fmt.Println(string(body))

	_, err = f.Write(body)
	if err != nil {
		fmt.Println("failed to write to file")
	}

	_ = f.Sync()
	return entity.TableData{}
}
