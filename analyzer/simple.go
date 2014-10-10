package analyzer

import (
	"github.com/keimoon/cerebro/tacit"
	"regexp"
	"strings"
)

var reXkcd = regexp.MustCompile(`xkcd ([0-9]+|random)`)

func Simple(question string, context tacit.Context) (string, error) {
	keywords := strings.ToLower(strings.TrimSpace(question))
	switch keywords {
	case "make me a sandwich", "sudo make me a sandwich", "cat", "mom":
		context["class"] = "easter egg"
	default:
		context["class"] = "programming"
	}
	if reXkcd.MatchString(question) {
		context["class"] = "xkcd"
	}
	return keywords, nil
}
