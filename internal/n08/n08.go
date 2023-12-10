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
var endsWithARE = regexp.MustCompile(`.*A$`)
var endsWithZRE = regexp.MustCompile(`.*Z$`)

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

func StepsToGoalMultidimensional(directions *[]int, nodes *map[string][]string) int {

	currPos := []string{}
	for k := range *nodes {
		if endsWithARE.MatchString(k) {
			currPos = append(currPos, k)
		}
	}

	totalSteps := []int{}
	for _, c := range currPos {
		steps := 0
		for {
			currDir := steps % len(*directions)

			// Update position
			c = (*nodes)[c][(*directions)[currDir]]

			// Increase the steps
			steps++

			// If the position is the goal, return the steps
			if endsWithZRE.MatchString(c) {
				totalSteps = append(totalSteps, steps)
				break
			}
		}
	}

	return lcm(totalSteps...)
}

// Greatest common divisor (GCD) via Euclidean algorithm
func gdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Find Least Common Multiple (LCM) via GCD
func lcm(integers ...int) int {
	a, b := integers[0], integers[1]
	result := a * b / gdc(a, b)

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
