package main

import (
	"fmt"
	"strconv"
)

const (
	Min = 359282
	Max = 820401
)

func main() {
	partOne()
}

func partOne() {
	count := 0
	for i := Min; i <= Max; i++ {
		if Match(i) {
			count++
		}
	}
	fmt.Println(count)
}

func Match(num int) bool {
	if num < Min || num > Max {
		return false
	}
	str := strconv.FormatInt(int64(num), 10)
	if len(str) != 6 {
		return false
	}
	adjacent := false
	for i := 0; i < 5; i++ {
		if str[i] == str[i+1] {
			adjacent = true
		}
		a, _ := strconv.Atoi(string(str[i]))
		b, _ := strconv.Atoi(string(str[i+1]))
		if a > b {
			return false
		}
	}
	if !adjacent {
		return false
	}

	return true
}
