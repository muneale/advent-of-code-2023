package n08

import (
	"fmt"
	"regexp"
	"strings"
)

var directions = map[rune]int{
	'L': 0,
	'R': 1,
}

var directionsRE = regexp.MustCompile(`[\(\),]`)

func ParseInput(input string) (*[]int, *map[string][]string) {
	lines := strings.Split(input, "\n")

	// Parse left/right
	directionsSlice := []int{}
	for _, d := range lines[0] {
		directionsSlice = append(directionsSlice, directions[d])
	}

	// Parse nodes
	nodes := map[string][]string{}
	for i := 2; i < len(lines); i++ {
		data := strings.Split(lines[i], " = ")
		idx := strings.TrimSpace(data[0])
		values := strings.Split(directionsRE.ReplaceAllString(data[1], ""), " ")
		if _, ok := nodes[idx]; ok {
			fmt.Printf("Duplicate node: %s\n", idx)
		}
		nodes[idx] = values
	}

	return &directionsSlice, &nodes
}

func StepsToGoal(directions *[]int, nodes *map[string][]string) int {

	currPos := "AAA"
	steps := 0

	for {
		currDir := steps % len(*directions)

		// Update position
		currPos = (*nodes)[currPos][(*directions)[currDir]]

		// Increase the steps
		steps++

		// If the position is the goal, return the steps
		if currPos == "ZZZ" {
			return steps
		}

	}
}
