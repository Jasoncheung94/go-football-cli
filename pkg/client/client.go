package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

func RequestData(url string, bind interface{}) {
	if url == "" {
		fmt.Println("url is required")
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("failed to create request")
		return
	}

	req.Header.Add("x-auth-token", viper.Get("apikey").(string))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("failed to send request")
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		fmt.Println("failed to decode response to bind", err)
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
		fmt.Println("failed to create request")
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
		fmt.Println("failed to write to file")
	}

	_ = f.Sync()
}
