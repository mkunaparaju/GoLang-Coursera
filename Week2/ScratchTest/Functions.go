package main

import "fmt"

type MyInt int

func (mi MyInt) Double() int {
	// MyInt is the receiver type.
	//Mi Is the variable that refers to the particular reciver object this Double invoked on
	return int(mi * 2)
}

func main() {
	V := MyInt(3)           // v is the particular object of MyInt
	fmt.Println(V.Double()) // we invoke the method double associtated with the type V.
}
