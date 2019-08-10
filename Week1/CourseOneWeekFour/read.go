package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	personList := make([]Person, 0)

	var fileName string
	fmt.Printf("Please enter a file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		person := Person{fName: lineSlice[0], lName: lineSlice[1]}
		personList = append(personList, person)
	}
	file.Close()
	for i := range personList {
		fmt.Printf("firstName: %s, lastName: %s \n", personList[i].fName,personList[i].lName)
	}

	if (err != nil) {
		fmt.Printf("err: ", err)
	}

}

type Person struct {
	fName string
	lName string
}
