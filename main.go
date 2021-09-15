package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	input := strings.Join(flag.Args(), "")
	if input == "" {
		fmt.Print(0)
		return
	}
	fmt.Print(len(strings.Split(input, " ")))
}
