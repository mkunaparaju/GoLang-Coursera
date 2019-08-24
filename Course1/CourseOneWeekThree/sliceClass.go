package main

import (
	"fmt"
	"sort"
)
import "strings"
import "strconv"

func main () {

	var numbers []int

	for {
		var str string
		fmt.Printf("please enter an integer or X to exit: ")
		fmt.Scan(&str)

		if strings.Compare(strings.ToLower(str), "x") == 0 {
			break
		}

		var number, err = strconv.Atoi(str)

		if err == nil {

			var temp = make([]int, len(numbers) + 1, cap(numbers) + 1)
			copy(temp, numbers)
			numbers = temp
			numbers[len(numbers) - 1] = number
			sort.Ints(numbers)

			fmt.Println("the numbers are: ", numbers)
		} else {
			fmt.Println(err)
		}
	}
}

