package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/keimoon/cerebro/tacit"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	Data struct {
		After    string      `json:"after"`
		Before   interface{} `json:"before"`
		Children []struct {
			Data struct {
				Subreddit string `json:"subreddit"`
				//SubredditId  string        `json:"subreddit_id"`
				Thumbnail string `json:"thumbnail"`
				Title     string `json:"title"`
				Ups       int64  `json:"ups"`
				Url       string `json:"url"`
				//UserReports  []interface{} `json:"user_reports"`
				//Visited      bool          `json:"visited"`
			} `json:"data"`
			//Kind string `json:"kind"`
		} `json:"children"`
		//Modhash string `json:"modhash"`
	} `json:"data"`
	//Kind string `json:"kind"`
}

func fetchJson(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func printResult(res result) string {
	var b bytes.Buffer
	items := res.Data.Children
	for index, item := range items {
		fmt.FPrintf(b, "%d. %s - %s (%d votes)\n", index, item.Data.Title, item.Data.Url, item.Data.Ups)
	}
	return b.String()
}

func Reddit(subreddit string, context tacit.Context) (string, error) {
	var res result
	err := FetchJson("http://www.reddit.com/r/programming.json", &res)
	if err != nil {
		return "", err
	}
	return printResult(res), nil
}
