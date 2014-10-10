package search

import (
	"bytes"
	"fmt"
	"github.com/keimoon/cerebro/tacit"
	"math/rand"
	"regexp"
	"strconv"
)

type xkcdResult struct {
	// Alt        string `json:"alt"`
	Day        int64  `json:"day,string"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      int64  `json:"month,string"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       int64  `json:"year,string"`
}

func xkcdPrint(res *xkcdResult) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "#%d. %s - %s", res.Num, res.Title, res.Img)
	return b.String()
}

var reXkcd = regexp.MustCompile(`xkcd ([0-9]+|random)`)

func Xkcd(question string, context tacit.Context) (string, error) {
	parts := reXkcd.FindStringSubmatch(question)
	// fmt.Printf("%#v", parts)
	if len(parts) == 0 {
		return "", nil
	}
	var num string
	if len(parts) == 1 {
		num = "latest"
	} else {
		num = parts[1] // 'random' or a number
	}

	var Url string
	if num != "random" {
		Url = "http://xkcd.com/" + num + "/info.0.json"
	} else {
		Url = "http://xkcd.com/info.0.json"
	}

	var res xkcdResult
	err := fetchJson(Url, &res)
	if err != nil {
		return "", err
	}
	if num == "random" {
		i := rand.Intn(res.Num-1) + 1
		Url = "http://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		err = fetchJson(Url, &res)
		if err != nil {
			return "", err
		}
	}

	answer := xkcdPrint(&res)
	// fmt.Print(answer)
	return answer, nil
}
