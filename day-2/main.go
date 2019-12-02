package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	partOne()
}

func partOne() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := []int{}
	for _, i := range bytes.Split(bytes.TrimSuffix(inputBytes, []byte("\n")), []byte(",")) {
		n, _ := strconv.Atoi(string(i))
		input = append(input, n)
	}
	input[1] = 12
	input[2] = 2
	for i := 0; i < len(input)-4; i += 4 {
		fmt.Println(input[i], input[i+1], input[i+2], input[i+3])
		switch input[i] {
		case 1:
			index := input[i+3]
			a := input[i+1]
			b := input[i+2]
			input[index] = input[a] + input[b]
		case 2:
			index := input[i+3]
			a := input[i+1]
			b := input[i+2]
			input[index] = input[a] * input[b]
		case 99:
			break
		default:
			fmt.Println("error:", input[i])
		}
	}
	fmt.Println(input[0])
}
