package exec

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Result struct {
	Responses []Response `json:"responses"`
}

// Response is return value to stdout
type Response struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
}

func crawler(url string, ch chan Response) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)

	statusCode := http.StatusNotFound

	if err != nil {
		log.Println(err)
	} else {
		statusCode = resp.StatusCode
		defer resp.Body.Close()
	}

	var data = Response{}
	data.StatusCode = statusCode
	data.URL = url
	ch <- data
}

func Exec() {
	ch := make(chan Response)
	list, err := readLines(viper.GetString("recordFileName"))
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range list {
		go crawler(u, ch)
	}

	data := make([]Response, 0)
	for range list {
		data = append(data, <-ch)
	}
	var res = Result{}
	res.Responses = data
	outputJSON, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, string(outputJSON))
}

func readLines(filePath string) ([]string, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Open file error")
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
