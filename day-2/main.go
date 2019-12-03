package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("Part one:", calc(12, 2))
	partTwo()
}

func partTwo() {
	for i := 0; i <= 99; i++ {
		for j := 99; j >= 0; j-- {
			v := calc(i, j)
			if v == 19690720 {
				fmt.Println(i, j, 100*i+j)
				return
			}
		}
	}
}

func calc(noun, verb int) int {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := []int{}
	for _, i := range bytes.Split(bytes.TrimSuffix(inputBytes, []byte("\n")), []byte(",")) {
		n, _ := strconv.Atoi(string(i))
		input = append(input, n)
	}
	input[1] = noun
	input[2] = verb
	for i := 0; i < len(input)-4; i += 4 {
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
	return input[0]
}
