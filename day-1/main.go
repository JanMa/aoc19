package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	var fuel int
	for _, i := range bytes.Split(inputBytes, []byte("\n")) {
		mass, _ := strconv.Atoi(string(i))
		tmpFuel := requiredFuel(mass)
		tmp := tmpFuel
		for requiredFuel(tmp) > 0 {
			tmptmp := requiredFuel(tmp)
			tmpFuel += tmptmp
			tmp = tmptmp
		}
		fuel += tmpFuel
	}
	fmt.Println(fuel)
}

func partOne() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	var fuel int
	for _, i := range bytes.Split(inputBytes, []byte("\n")) {
		mass, _ := strconv.Atoi(string(i))
		fuel += requiredFuel(mass)
	}
	fmt.Println(fuel)
}

func requiredFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel > 0 {
		return fuel
	}
	return 0
}
