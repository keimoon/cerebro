package search

import (
	"bytes"
	"fmt"
	"github.com/keimoon/cerebro/tacit"
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

func printResult(res result) string {
	var b bytes.Buffer
	items := res.Data.Children
	for index, item := range items {
		fmt.Fprintf(&b, "%d. %s - %s (%d votes)\n", index, item.Data.Title, item.Data.Url, item.Data.Ups)
	}
	return b.String()
}

func Reddit(subreddit string, context tacit.Context) (string, error) {
	var res result
	err := fetchJson("http://www.reddit.com/r/programming.json", &res)
	if err != nil {
		return "", err
	}
	return printResult(res), nil
}
