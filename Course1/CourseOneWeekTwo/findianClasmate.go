package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a string")

	consoleReader := bufio.NewReader(os.Stdin)
	str, _ := consoleReader.ReadString('\n')

	formattedStr := strings.ToLower(strings.TrimSpace(str))

	if strings.HasPrefix(formattedStr, "i") && strings.HasSuffix(formattedStr, "n") && strings.Contains(formattedStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
