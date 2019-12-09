package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	// "strconv"
	// "math"
	"strings"
)

func main() {
	partOne(parseInput())
	// partOne([]string{
	// 	"COM)B",
	// 	"B)C",
	// 	"C)D",
	// 	"D)E",
	// 	"E)F",
	// 	"B)G",
	// 	"G)H",
	// 	"D)I",
	// 	"E)J",
	// 	"J)K",
	// 	"K)L",
	// })
	partTwo(parseInput())
	// partTwo([]string{
	// 	"COM)B",
	// 	"B)C",
	// 	"C)D",
	// 	"D)E",
	// 	"E)F",
	// 	"B)G",
	// 	"G)H",
	// 	"D)I",
	// 	"E)J",
	// 	"J)K",
	// 	"K)L",
	// 	"K)YOU",
	// 	"I)SAN",
	// })
}

func parseInput() []string {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := []string{}
	for _, i := range bytes.Split(bytes.TrimSuffix(inputBytes, []byte("\n")), []byte("\n")) {
		input = append(input, string(i))
	}
	return input
}

type obj struct {
	name  string
	orbit *obj
}

func partOne(input []string) {
	objects := generateOrbitMap(input)
	cnt := 0
	for _, o := range objects {
		c := countOrbits(o)
		// fmt.Println(o.name, c)
		cnt += c
	}
	fmt.Printf("Total orbits: %d\n", cnt)
}

func partTwo(input []string) {
	objects := generateOrbitMap(input)
	interSec := findIntersection(objects["YOU"], objects["SAN"])
	interSec.orbit = nil
	me := countOrbits(objects["YOU"]) - 1
	santa := countOrbits(objects["SAN"]) - 1
	fmt.Printf("Total transfers: %d\n", me+santa)
}

func generateOrbitMap(input []string) map[string]*obj {
	objects := make(map[string]*obj)
	for _, l := range input {
		pair := strings.Split(l, ")")
		if _, ok := objects[pair[0]]; !ok {
			objects[pair[0]] = &obj{
				name:  pair[0],
				orbit: nil,
			}
		}
		if _, ok := objects[pair[1]]; !ok {
			objects[pair[1]] = &obj{
				name:  pair[1],
				orbit: objects[pair[0]],
			}
		} else {
			objects[pair[1]].orbit = objects[pair[0]]
		}
	}
	return objects
}

func countOrbits(o *obj) int {
	if o.orbit == nil {
		// fmt.Printf("%s\n", o.name)
		return 0
	}
	// fmt.Printf("%s -> ", o.name)
	return countOrbits(o.orbit) + 1
}

func findIntersection(a, b *obj) *obj {
	first := a
	pathA := make(map[string]*obj)
	for first.orbit != nil {
		pathA[first.name] = first
		first = first.orbit
	}
	second := b
	for second.orbit != nil {
		if _, ok := pathA[second.name]; ok {
			return second
		}
		second = second.orbit
	}
	return nil
}
