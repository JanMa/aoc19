package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("Part one:")
	fmt.Println(process(1, parseInput()))
	// fmt.Println("TEST:")
	// fmt.Println(process(8, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}))
	// fmt.Println(process(8, []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}))
	// fmt.Println(process(8, []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}))
	// fmt.Println(process(8, []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}))
	// fmt.Println(process(0, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}))
	// fmt.Println(process(1, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}))
	// fmt.Println(process(0, []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}))
	// fmt.Println(process(1, []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}))
	// fmt.Println(process(1, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}))
	// fmt.Println(process(8, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}))
	// fmt.Println(process(9, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}))
	fmt.Println("Part two:")
	fmt.Println(process(5, parseInput()))
}

func parseInput() []int {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := []int{}
	for _, i := range bytes.Split(bytes.TrimSuffix(inputBytes, []byte("\n")), []byte(",")) {
		n, _ := strconv.Atoi(string(i))
		input = append(input, n)
	}
	return input
}

func process(in int, program []int) []int {
	out := []int{}
	i := 0
	for {
		mode, op := parseMode(program[i])
		mS := func(pos int) int {
			switch mode[pos] {
			case 1:
				return program[i+pos+1]
			default:
				return program[program[i+pos+1]]
			}
		}
		switch op {
		case 1:
			index := program[i+3]
			program[index] = mS(0) + mS(1)
			i += 4
		case 2:
			index := program[i+3]
			program[index] = mS(0) * mS(1)
			i += 4
		case 3:
			index := program[i+1]
			program[index] = in
			i += 2
		case 4:
			out = append(out, mS(0))
			i += 2
		case 5:
			if mS(0) != 0 {
				i = mS(1)
			} else {
				i += 3
			}
		case 6:
			if mS(0) == 0 {
				i = mS(1)
			} else {
				i += 3
			}
		case 7:
			if mS(0) < mS(1) {
				program[program[i+3]] = 1
			} else {
				program[program[i+3]] = 0
			}
			i += 4
		case 8:
			if mS(0) == mS(1) {
				program[program[i+3]] = 1
			} else {
				program[program[i+3]] = 0
			}
			i += 4
		case 99:
			return out
		default:
			fmt.Println("error:", program[i])
			return out
		}
	}
}

func parseMode(input int) (mode []int, op int) {
	d1 := input % 10
	d2 := input % 100 / 10
	d3 := input % 1000 / 100
	d4 := input % 10000 / 1000
	d5 := input % 100000 / 10000
	d6 := input % 1000000 / 100000

	return []int{d3, d4, d5, d6}, d1 + d2*10
}
