package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Please enter upto 10 space seperated numbers, then press enter\n")

	consoleReader := bufio.NewScanner(os.Stdin)
	consoleReader.Scan()
	inString := strings.TrimSpace(consoleReader.Text())
	splitArray := strings.Split(inString, " ")

	fmt.Printf("input String: %s\n", splitArray)

	if len(splitArray) > 10 {
		fmt.Print("You have entered more than 10 integers. Please re-execute with less than 10 integers \n")
		return
	}

	integerArr := createIntSlice(splitArray)

	BubbleSort(integerArr)

	fmt.Print(integerArr)
}

func createIntSlice(inArr []string) []int {
	var r []int
	for _, v := range inArr {
		i, _ := strconv.Atoi(v)
		r = append(r, i)
	}
	return r
}

func BubbleSort(inArr []int) {
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < len(inArr)-1; i++ {
			if inArr[i] > inArr[i+1] {
				Swap(inArr, i)
				isSorted = false
			}
		}

	}
}

func Swap(inArr []int, i int) {
	temp := inArr[i]
	inArr[i] = inArr[i+1]
	inArr[i+1] = temp
}
