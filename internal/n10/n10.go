package n10

import (
	"math"
	"regexp"
	"strings"
)

var upRE = regexp.MustCompile(`[|7F]`)
var downRE = regexp.MustCompile(`[|LJ]`)
var leftRe = regexp.MustCompile(`[\-FL]`)
var rightRe = regexp.MustCompile(`[\-J7]`)

var directionsMap = map[string]int{
	"N": 0,
	"E": 1,
	"S": 2,
	"W": 3,
}

var directionsMapReverse = map[int]string{
	0: "N",
	1: "E",
	2: "S",
	3: "W",
}

func ParseInput(data string) [][]string {
	gameMap := [][]string{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		row := []string{}
		for _, c := range line {
			row = append(row, string(c))
		}
		gameMap = append(gameMap, row)
	}
	return gameMap
}

func getVertexs(gameMap *[][]string) [][]int {

	// Visited map
	vertexs := [][]int{}

	visited := make([][]int, len(*gameMap))

	for i := range visited {
		visited[i] = make([]int, len((*gameMap)[i]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	startR, startC := 0, 0
	for i := range *gameMap {
		for j := range (*gameMap)[i] {
			if (*gameMap)[i][j] == "S" {
				startR, startC = i, j
				break
			}
		}
	}
	vertexs = append(vertexs, []int{startR, startC})
	nodes := [][]int{{startR, startC, -1}}
	for len(nodes) > 0 {

		i, j, d := nodes[0][0], nodes[0][1], nodes[0][2]
		nodes = nodes[1:]

		// Out of bound
		if i < 0 || i >= len(*gameMap) || j < 0 || j >= len((*gameMap)[i]) {
			continue
		}

		// Visited
		if visited[i][j] >= 0 {
			continue
		}

		// fmt.Println("Visited map:")
		// for i := range visited {
		// 	fmt.Printf("%v\n", (visited)[i])
		// }

		e := (*gameMap)[i][j]

		if e == "S" {
			visited[i][j] = 0

			if i-1 >= 0 && upRE.MatchString((*gameMap)[i-1][j]) {
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			} else if i+1 < len(*gameMap) && downRE.MatchString((*gameMap)[i+1][j]) {
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			} else if j+1 < len((*gameMap)[i]) && rightRe.MatchString((*gameMap)[i][j+1]) {
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			} else if j-1 >= 0 && leftRe.MatchString((*gameMap)[i][j-1]) {
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

		if e == "|" {
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
		}

		if e == "-" {
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

		if e == "7" {
			vertexs = append(vertexs, []int{i, j})
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
		}

		if e == "F" {
			vertexs = append(vertexs, []int{i, j})
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
		}

		if e == "L" {
			vertexs = append(vertexs, []int{i, j})
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
		}

		if e == "J" {
			vertexs = append(vertexs, []int{i, j})
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

	}

	vertexs = append(vertexs, []int{startR, startC})

	return vertexs
}

func GetMaximumDistance(gameMap *[][]string) int {

	// Visited map
	visited := make([][]int, len(*gameMap))

	for i := range visited {
		visited[i] = make([]int, len((*gameMap)[i]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	// Find S
	startR, startC := 0, 0
	for i := range *gameMap {
		for j := range (*gameMap)[i] {
			if (*gameMap)[i][j] == "S" {
				startR, startC = i, j
				break
			}
		}
	}

	nodes := [][]int{{startR, startC, -1}}
	for len(nodes) > 0 {

		i, j, d := nodes[0][0], nodes[0][1], nodes[0][2]
		nodes = nodes[1:]

		// Out of bound
		if i < 0 || i >= len(*gameMap) || j < 0 || j >= len((*gameMap)[i]) {
			continue
		}

		// Visited
		if visited[i][j] >= 0 {
			continue
		}

		// fmt.Println("Visited map:")
		// for i := range visited {
		// 	fmt.Printf("%v\n", (visited)[i])
		// }

		e := (*gameMap)[i][j]

		if e == "S" {
			visited[i][j] = 0

			if i-1 >= 0 && upRE.MatchString((*gameMap)[i-1][j]) {
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
			if i+1 < len(*gameMap) && downRE.MatchString((*gameMap)[i+1][j]) {
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
			if j+1 < len((*gameMap)[i]) && rightRe.MatchString((*gameMap)[i][j+1]) {
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
			if j-1 >= 0 && leftRe.MatchString((*gameMap)[i][j-1]) {
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

		if e == "|" {
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
		}

		if e == "-" {
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

		if e == "7" {
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
		}

		if e == "F" {
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i + 1, j, directionsMap["S"]})
			}
			if directionsMapReverse[d] == "N" {
				visited[i][j] = visited[i+1][j] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
		}

		if e == "L" {
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i, j + 1, directionsMap["E"]})
			}
			if directionsMapReverse[d] == "W" {
				visited[i][j] = visited[i][j+1] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
		}

		if e == "J" {
			if directionsMapReverse[d] == "E" {
				visited[i][j] = visited[i][j-1] + 1
				nodes = append(nodes, []int{i - 1, j, directionsMap["N"]})
			}
			if directionsMapReverse[d] == "S" {
				visited[i][j] = visited[i-1][j] + 1
				nodes = append(nodes, []int{i, j - 1, directionsMap["W"]})
			}
		}

	}

	max := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] > max {
				max = visited[i][j]
			}
		}
	}
	return max

}

/**
* Pick's theorem (https://en.wikipedia.org/wiki/Pick%27s_theorem)
* loopArea = interiorPointsCount + (boundaryPointsCount / 2) - 1
*
* Part 2 answer is interiorPointsCount
* transforming Pick's formula:
* interiorPointsCount = loopArea - (boundaryPointsCount / 2) + 1
*
* boundaryPointsCount is length of loop (practically part1 answer * 2)
*
* loopArea can by calculated using Shoelace formula (https://en.wikipedia.org/wiki/Shoelace_formula):
* vertices = (x1, y1) (x2, y2) (x3, y3) ...
* 2 * loopArea = x1 * y2 - y1 * x2 + x2 * y3 - x3 * y2 + ...
* loopArea = result / 2
 */
func GetNests(gameMap *[][]string, boundaryPointsCount int) int {

	loopArea := 0
	boundaryPointsCount *= 2
	vertexs := getVertexs(gameMap)

	for i := 0; i < len(vertexs)-1; i++ {
		x1, y1 := vertexs[i][1], vertexs[i][0]
		x2, y2 := vertexs[i+1][1], vertexs[i+1][0]
		loopArea += x1*y2 - y1*x2
	}

	loopArea /= 2

	return int(math.Abs(float64(loopArea))) - (boundaryPointsCount / 2) + 1
}
