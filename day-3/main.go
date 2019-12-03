package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	// "strconv"
)

func main() {
	partOne()
}

func partOne() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := [2][]string{}
	for index, i := range bytes.Split(bytes.TrimSuffix(inputBytes, []byte("\n")), []byte("\n")) {
		for _, j := range bytes.Split(i, []byte(",")) {
			input[index] = append(input[index], string(j))
		}
	}
	fmt.Printf("%+v\n", input[0])
	fmt.Printf("%+v\n", input[1])
}
