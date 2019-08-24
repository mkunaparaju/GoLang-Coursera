package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var addressMap map[string]string
	addressMap = make(map[string]string)

	var name string
	fmt.Printf("Please enter a name: ")
	fmt.Scan(&name)

	var address string
	fmt.Printf("Please enter corresponding address: ")
	fmt.Scan(&address)

	addressMap[name] = address

	mapByteArr, err := json.Marshal(addressMap)

	fmt.Println("The JSON object as follows")
	fmt.Println(string(mapByteArr))

	if err != nil {
		fmt.Println(err)
	}

}
