package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Filters struct {
	DateFrom     time.Time
	DateTo       time.Time
	Stage        string
	Status       string
	MatchDay     int
	Group        string
	Season       string //year
	Competitions []string
}

func RequestData(url string, bind interface{}, f *Filters) {
	if url == "" {
		fmt.Println("url is required")
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("failed to create request")
		return
	}

	if val, ok := viper.Get("apikey").(string); ok {
		req.Header.Add("x-auth-token", val)
	} else {
		fmt.Println("TODO read from cache or dummy data.")
	}

	queryParams := req.URL.Query()
	if f != nil {
		switch {
		case len(f.Competitions) > 0:
			queryParams.Add("competitions", strings.Join(f.Competitions, ","))
			queryParams.Add("status", "SCHEDULED")
			// For dates the API requires year month day.
			queryParams.Add("dateFrom", time.Now().Format("2006-01-02"))
			queryParams.Add("dateTo", time.Now().Add((time.Hour*24)*7).Format("2006-01-02"))
		}
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("failed to send request")
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		fmt.Println("failed to decode response to bind", err)
		return
	}
}

func RequestToFile(url string) {
	if url == "" {
		fmt.Println("url is required")
		return
	}
	unix := time.Now().Unix()
	f, err := os.Create(fmt.Sprintf("./%v", unix))
	if err != nil {
		panic("failed to create file to read into")
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("failed to create request")
	}

	req.Header.Set("Content-type", "application/json")
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
		panic("failed to write to file")
	}

	_ = f.Sync()
}
