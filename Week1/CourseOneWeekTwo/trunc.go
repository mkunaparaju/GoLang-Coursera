package main

import "fmt"

func main(){
	//fmt.Printf("Inside Main Function \n\n")
	fmt.Printf("Please enter a floating point number here: \n")

	inputNumber := 0.0
	num, err := fmt.Scan(&inputNumber)

	fmt.Printf("number %f \n\n", inputNumber)
	fmt.Printf("number of inputs: %d and errror: %s \n\n", num, err)

	numConvertedToInteger := int32(inputNumber)
	fmt.Printf("the converted number: %d \n", numConvertedToInteger)
}

