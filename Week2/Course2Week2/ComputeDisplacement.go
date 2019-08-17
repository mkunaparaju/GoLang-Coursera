package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter Acceleration, Inital Velocity, Inital displacement, seperated by comma : ")

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println("Enter time for which you want to calculate displacement: ")

	var timeInput string
	if scanner.Scan() {
		timeInput = scanner.Text()
	}

	splitSplice := strings.Split(input, ",")
	accelaration, err := strconv.ParseFloat(strings.TrimSpace(splitSplice[0]), 64)
	initVelocity, err1 := strconv.ParseFloat(strings.TrimSpace(splitSplice[1]), 64)
	initDisplacement, err2 := strconv.ParseFloat(strings.TrimSpace(splitSplice[2]), 64)
	time, err3 := strconv.ParseFloat(strings.TrimSpace(timeInput), 64)
	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		fmt.Print("Errored out")
		return
	}

	fmt.Printf("Accelaration %f, Initial Velocity %f, Inital Displacement %f, Time to calc Displacement: %f ", accelaration, initVelocity, initDisplacement, time)
	calcDisplacement := GenDisplaceFn(accelaration, initVelocity, initDisplacement)

	fmt.Printf("Calculated dispacement for given time is : %f", calcDisplacement(time))

}

func GenDisplaceFn(acc, vel, disp float64) func(float64) float64 {
	calculateDisplacement := func(time float64) float64 {
		return (0.5 * acc * math.Pow(time, 2)) + vel*time + disp
	}
	return calculateDisplacement
}
