package search

import (
	"github.com/keimoon/cerebro/tacit"
)

func EasterEgg(keyword string, context tacit.Context) (string, error) {
	switch keyword {
	case "make me a sandwich":
		return "Make it yourself, son", nil
	case "sudo make me a sandwich":
		return "Okay", nil
	case "cat":
		return "You are a kitty!!", nil
	case "mom":
		return "Yes?", nil
	default:
		return "", nil
	}
}
