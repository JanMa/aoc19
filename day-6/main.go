package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	// "strconv"
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
	cnt := 0
	for _, o := range objects {
		c := countOrbits(o)
		// fmt.Println(o.name, c)
		cnt += c
	}
	fmt.Printf("Total orbits: %d\n", cnt)
}

func countOrbits(o *obj) int {
	if o.orbit == nil {
		// fmt.Printf("%s\n", o.name)
		return 0
	}
	// fmt.Printf("%s -> ", o.name)
	return countOrbits(o.orbit) + 1
}
