package search

import (
	"github.com/keimoon/cerebro/tacit"
)

func SimpleChooser(context tacit.Context) tacit.SearchEngine {
	switch context["class"] {
	case "easter egg":
		return EasterEgg
	default:
		return Howgoi
	}
}
