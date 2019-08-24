package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	fmt.Printf("Inside the main loop. Initialiing the array \n")
	arr := make([]int, 0)

	input := ""

	for {
		fmt.Println("Provide one input integer: ")
		fmt.Scan(&input)
		fmt.Println("You have provided the input as : ", input)

		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Provided non integer. EXITING...")
			break;
		}
		arr = append(arr, i)
		sort.Ints(arr)
		fmt.Println("Sorted Array : ", arr )
	}
}
