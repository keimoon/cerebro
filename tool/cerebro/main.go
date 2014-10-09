package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/keimoon/cerebro"
)

func main() {
	question := strings.Join(os.Args[1:], " ")
	answer, err := cerebro.Cerebro.Ask(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}
