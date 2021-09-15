package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	input := strings.Join(flag.Args(), "")
	fmt.Print(len(strings.Split(input, " ")))
}
