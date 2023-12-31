package n16

import (
	"strings"
)

func ParseInput(data string) *[][]rune {

	g := [][]rune{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		g = append(g, row)
	}
	return &g
}

func countTiles(g *[][]int) int {
	tiles := 0
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[0]); j++ {
			if (*g)[i][j] > 0 {
				tiles++
			}
		}
	}
	return tiles
}

func CountEnergizedTiles(g *[][]rune) int {

	en := computeEnergizedGrid(g)

	return countTiles(en)
}

func computeEnergizedGrid(g *[][]rune) *[][]int {

	R := len(*g)
	C := len((*g)[0])

	ng := make([][]int, R)
	for i := 0; i < R; i++ {
		ng[i] = make([]int, C)
	}

	computeNext(0, 0, 'E', g, &ng, []int{0, 0, 0})
	return &ng
}

func ComputeMaxEnergizedGrid(g *[][]rune) int {

	R := len(*g)
	C := len((*g)[0])
	max := 0

	ng := make([][]int, R)
	for i := 0; i < R; i++ {
		ng[i] = make([]int, C)
	}

	// Start from left
	for i := 0; i < R; i++ {
		computeNext(i, 0, 'E', g, &ng, []int{0, 0, 0})
		curr := countTiles(&ng)
		if curr > max {
			max = curr
		}
		ng = make([][]int, R)
		for i := 0; i < R; i++ {
			ng[i] = make([]int, C)
		}
	}

	// Start from right
	for i := 0; i < R; i++ {
		computeNext(i, C-1, 'W', g, &ng, []int{0, 0, 0})
		curr := countTiles(&ng)
		if curr > max {
			max = curr
		}
		ng = make([][]int, R)
		for i := 0; i < R; i++ {
			ng[i] = make([]int, C)
		}
	}

	// Start from top
	for i := 0; i < C; i++ {
		computeNext(0, i, 'S', g, &ng, []int{0, 0, 0})
		curr := countTiles(&ng)
		if curr > max {
			max = curr
		}
		ng = make([][]int, R)
		for i := 0; i < R; i++ {
			ng[i] = make([]int, C)
		}
	}

	// Start from bottom
	for i := 0; i < C; i++ {
		computeNext(R-1, i, 'N', g, &ng, []int{0, 0, 0})
		curr := countTiles(&ng)
		if curr > max {
			max = curr
		}
		ng = make([][]int, R)
		for i := 0; i < R; i++ {
			ng[i] = make([]int, C)
		}
	}

	return max
}

func computeNext(i, j int, d rune, g *[][]rune, ng *[][]int, lasts []int) {
	maxPass := 100
	if i < 0 || i >= len(*g) || j < 0 || j >= len((*g)[0]) {
		return
	}

	if (*ng)[i][j] > maxPass && lasts[0] > maxPass && lasts[1] > maxPass && lasts[2] > maxPass {
		return
	}

	(*ng)[i][j]++
	lasts = lasts[1:]
	lasts = append(lasts, (*ng)[i][j])

	// for k := 0; k < len(*ng); k++ {
	// 	for l := 0; l < len((*ng)[k]); l++ {
	// 		if k == i && l == j {
	// 			fmt.Printf("%s ", "X")
	// 		} else {
	// 			if (*ng)[k][l] > 0 {
	// 				fmt.Printf("%s ", "#")
	// 			} else {
	// 				fmt.Printf("%s ", ".")
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
	// fmt.Printf("\n")

	if (*g)[i][j] == '.' {
		if d == 'N' {
			computeNext(i-1, j, d, g, ng, lasts)
		} else if d == 'S' {
			computeNext(i+1, j, d, g, ng, lasts)
		} else if d == 'E' {
			computeNext(i, j+1, d, g, ng, lasts)
		} else if d == 'W' {
			computeNext(i, j-1, d, g, ng, lasts)
		}
	} else if (*g)[i][j] == '\\' {
		if d == 'N' {
			computeNext(i, j-1, 'W', g, ng, lasts)
		} else if d == 'S' {
			computeNext(i, j+1, 'E', g, ng, lasts)
		} else if d == 'E' {
			computeNext(i+1, j, 'S', g, ng, lasts)
		} else if d == 'W' {
			computeNext(i-1, j, 'N', g, ng, lasts)
		}
	} else if (*g)[i][j] == '/' {
		if d == 'N' {
			computeNext(i, j+1, 'E', g, ng, lasts)
		} else if d == 'S' {
			computeNext(i, j-1, 'W', g, ng, lasts)
		} else if d == 'E' {
			computeNext(i-1, j, 'N', g, ng, lasts)
		} else if d == 'W' {
			computeNext(i+1, j, 'S', g, ng, lasts)
		}
	} else if (*g)[i][j] == '|' {
		if d == 'N' {
			computeNext(i-1, j, d, g, ng, lasts)
		} else if d == 'S' {
			computeNext(i+1, j, d, g, ng, lasts)
		} else {
			computeNext(i-1, j, 'N', g, ng, lasts)
			computeNext(i+1, j, 'S', g, ng, lasts)
		}
	} else if (*g)[i][j] == '-' {
		if d == 'E' {
			computeNext(i, j+1, d, g, ng, lasts)
		} else if d == 'W' {
			computeNext(i, j-1, d, g, ng, lasts)
		} else {
			computeNext(i, j-1, 'W', g, ng, lasts)
			computeNext(i, j+1, 'E', g, ng, lasts)
		}
	}
}
