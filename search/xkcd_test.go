package search

import (
	"testing"
)

func TestXkcd(t *testing.T) {
	_, err := Xkcd("xkcd random")
	if err != nil {
		t.Fatal(err)
	}
}
