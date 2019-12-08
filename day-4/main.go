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
	fmt.Printf("Part one: ")
	partOne()
	fmt.Printf("Part two: ")
	partTwo()
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

func partTwo() {
	count := 0
	for i := Min; i <= Max; i++ {
		if Two(i) {
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

func Two(num int) bool {
	d1 := num % 10
	d2 := num % 100 / 10
	d3 := num % 1000 / 100
	d4 := num % 10000 / 1000
	d5 := num % 100000 / 10000
	d6 := num % 1000000 / 100000

	twoMatching := false
	if d1 == d2 && d2 != d3 ||
		d1 != d2 && d2 == d3 && d3 != d4 ||
		d2 != d3 && d3 == d4 && d4 != d5 ||
		d3 != d4 && d4 == d5 && d5 != d6 ||
		d4 != d5 && d5 == d6 {
		twoMatching = true
	}
	inc := false
	if d6 <= d5 && d5 <= d4 && d4 <= d3 && d3 <= d2 && d2 <= d1 {
		inc = true
	}
	return twoMatching && inc
}
