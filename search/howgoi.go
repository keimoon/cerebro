package search

import (
	"github.com/keimoon/cerebro/tacit"
	"github.com/mattn/howgoi"
)

func Howgoi(keyword string, context tacit.Context) (string, error) {
	answers, err := howgoi.Query(keyword)
	if err != nil {
                return "", err
        }
	if len(answers) == 0 {
                return "", nil
        }
	anwser := answers[0]
	return anwser.Code + "\n" + anwser.Link, nil
}
