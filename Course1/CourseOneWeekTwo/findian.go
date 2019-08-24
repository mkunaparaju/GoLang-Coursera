package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Starting findian class \n")

	var inputString string

	fmt.Printf("Please enter text: \n")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString = strings.ToLower(scanner.Text())
	}
	len := len(inputString)

	start := inputString[0] == 'i'
	end := inputString[len-1] == 'n'
	middle := strings.Contains(inputString, "a")

	if start && middle && end {
		fmt.Printf("Found")
		return
	}

	fmt.Printf("Not Found")
}

//I d skd a efju N
