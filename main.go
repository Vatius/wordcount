package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewReader(os.Stdin)
	input, _ = scanner.ReadString('\n')
	input = strings.Replace(input, "\n", "", 1)
	fmt.Print(len(strings.Split(input, " ")))
}
