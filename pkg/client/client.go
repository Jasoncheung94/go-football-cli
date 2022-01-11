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
	CacheFile    string
}

func RequestData(url string, bind interface{}, f *Filters) {
	if url == "" {
		panic("url is required")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("failed to create request")
	}

	if val, ok := viper.Get("apikey").(string); ok {
		req.Header.Add("x-auth-token", val)
	} else {
		RequestFromFile(f.CacheFile, &bind)
		// Return early the with data.
		return
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
		panic("failed to send request")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		panic(err)
	}
}

func RequestToFile(url string, filter *Filters) {
	if url == "" {
		panic("url is required")
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
	if val, ok := viper.Get("apikey").(string); ok {
		req.Header.Add("x-auth-token", val)
	}

	queryParams := req.URL.Query()
	if filter != nil {
		switch {
		case len(filter.Competitions) > 0:
			queryParams.Add("competitions", strings.Join(filter.Competitions, ","))
			queryParams.Add("status", "SCHEDULED")
			// For dates the API requires year month day.
			queryParams.Add("dateFrom", time.Now().Format("2006-01-02"))
			queryParams.Add("dateTo", time.Now().Add((time.Hour*24)*7).Format("2006-01-02"))
		}
	}
	req.URL.RawQuery = queryParams.Encode()

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

func RequestFromFile(filename string, bind interface{}) {
	file, err := ioutil.ReadFile("./data/" + filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &bind)
	if err != nil {
		panic("failed to bind data")
	}
}
