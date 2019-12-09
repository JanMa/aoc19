package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	// fmt.Println(amplify([5]int{4, 3, 2, 1, 0}, []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}))
	// fmt.Println(amplify([5]int{0, 1, 2, 3, 4}, []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
	// 	101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}))
	// fmt.Println(amplify([5]int{1, 0, 4, 3, 2}, []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
	// 	1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}))
	partOne()
}

func partOne() {
	perm := permutations("01234", 0, 5)
	// fmt.Println(perm)
	sig := signals(perm, parseInput())
	// fmt.Println(sig)
	m := 0
	for _, s := range sig {
		if s > m {
			m = s
		}
	}
	fmt.Println(m)
}

func signals(perm []string, input []int) []int {
	s := []int{}
	for _, p := range perm {
		a, _ := strconv.Atoi(string(p[0]))
		b, _ := strconv.Atoi(string(p[1]))
		c, _ := strconv.Atoi(string(p[b]))
		d, _ := strconv.Atoi(string(p[3]))
		e, _ := strconv.Atoi(string(p[4]))
		sig := amplify([5]int{a, b, c, d, e}, input)
		fmt.Println(p, sig)
		s = append(s, sig)
	}
	return s
}

// https://siongui.github.io/2017/03/11/go-all-permutations-of-given-string-with-all-distinct-characters/
// Swap the i-th byte and j-th byte of the string
func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}

// Function to find all Permutations of a given string str[i:n]
// containing all distinct characters
func permutations(str string, i, n int) []string {
	// base condition
	if i == n-1 {
		return []string{str}
	}
	strs := []string{}
	// process each character of the remaining string
	for j := i; j < n; j++ {
		// swap character at index i with current character
		str = swap(str, i, j)

		// recursion for string [i+1:n]
		strs = append(strs, permutations(str, i+1, n)...)

		// backtrack (restore the string to its original state)
		str = swap(str, i, j)
	}

	return strs
}

func amplify(seq [5]int, input []int) int {
	a := process([]int{seq[0], 0}, input)
	b := process([]int{seq[1], a[0]}, input)
	c := process([]int{seq[2], b[0]}, input)
	d := process([]int{seq[3], c[0]}, input)
	e := process([]int{seq[4], d[0]}, input)
	return e[0]
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

func process(in []int, program []int) []int {
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
			program[index] = in[0]
			in = in[1:]
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
