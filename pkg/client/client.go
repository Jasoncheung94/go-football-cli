package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

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
