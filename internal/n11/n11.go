package n11

import (
	"math"
	"slices"
	"strings"
	"sync"
)

type Galaxy struct {
	x int
	y int
}

func ParseInput(input string) *[][]rune {

	lines := strings.Split(input, "\n")
	universe := make([][]rune, len(lines))

	for i, line := range lines {
		for _, c := range line {
			universe[i] = append(universe[i], c)
		}
	}

	// return expandUniverse(&universe)
	return &universe
}

func expandUniverse(universe *[][]rune) *[][]rune {
	expandedUniverse := [][]rune{}

	// Find the empty rows, cols
	emptyRows, emptyCols := []int{}, []int{}

	for i, row := range *universe {
		empty := true
		for _, c := range row {
			if c == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, i)
		}
	}

	for i := 0; i < len((*universe)[0]); i++ {
		empty := true
		for _, row := range *universe {
			if row[i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}

	for i := 0; i < len(*universe); i++ {
		newRow := []rune{}
		for j := 0; j < len((*universe)[i]); j++ {
			if slices.Contains(emptyCols, j) {
				newRow = append(newRow, '.')
			}
			newRow = append(newRow, (*universe)[i][j])
		}
		if slices.Contains(emptyRows, i) {
			emptyRow := make([]rune, len(newRow))
			for j := 0; j < len(newRow); j++ {
				emptyRow[j] = '.'
			}
			expandedUniverse = append(expandedUniverse, emptyRow)
		}
		expandedUniverse = append(expandedUniverse, newRow)
	}

	// for _, row := range *universe {
	// 	expand := true
	// 	for _, c := range row {
	// 		if c == '#' {
	// 			expand = false
	// 			break
	// 		}
	// 	}
	// 	if expand {
	// 		expandedUniverse = append(expandedUniverse, expandedRow)
	// 	}
	// 	expandedUniverse = append(expandedUniverse, row)
	// }

	// // If the col does not contain an hashtag, add a new row
	// expandedCol := make([]rune, len(expandedUniverse[0]))
	// for i := 0; i < len((*universe)[0]); i++ {
	// 	expandedCol[i] = '.'
	// }

	// for i := 0; i < len((*universe)[0]); i++ {
	// 	expand := true
	// 	for j := 0; j < len(*universe); j++ {
	// 		if (*universe)[j][i] == '#' {
	// 			expand = false
	// 			break
	// 		}

	// 	}
	// 	if expand {
	// 		expandedUniverse = append(expandedUniverse, expandedCol)
	// 	}
	// 	expandedUniverse[i] = append(expandedUniverse[i], expandedCol...)
	// }

	return &expandedUniverse
}

func getEmptyRows(universe *[][]rune) []int {

	emptyRows := []int{}

	for i, row := range *universe {
		empty := true
		for _, c := range row {
			if c == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, i)
		}
	}

	return emptyRows
}

func getEmptyCols(universe *[][]rune) []int {

	emptyCols := []int{}

	for i := 0; i < len((*universe)[0]); i++ {
		empty := true
		for _, row := range *universe {
			if row[i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}

	return emptyCols
}

func getGalaxies(universe *[][]rune) []Galaxy {
	galaxies := []Galaxy{}
	for i := 0; i < len(*universe); i++ {
		for j := 0; j < len((*universe)[i]); j++ {
			if (*universe)[i][j] == '#' {
				galaxies = append(galaxies, Galaxy{i, j})
			}
		}
	}
	return galaxies
}

func getMinumDistanceGalaxies(a, b Galaxy, emptyRows, emptyCols *[]int, factor int) int {

	ax, bx := a.x, b.x
	ay, by := a.y, b.y

	cAx, cBx := 0, 0
	for _, row := range *emptyRows {
		if ax > row {
			cAx += factor - 1
		}
		if bx > row {
			cBx += factor - 1
		}
	}
	ax += cAx
	bx += cBx

	cAy, cBy := 0, 0
	for _, col := range *emptyCols {
		if ay > col {
			cAy += factor - 1
		}
		if by > col {
			cBy += factor - 1
		}
	}
	ay += cAy
	by += cBy

	l1, l2 := int(math.Abs(float64(ax-bx))), int(math.Abs(float64(ay-by)))

	if l1 <= l2 {
		return 2*l1 + (l2 - l1)
	}
	return 2*l2 + (l1 - l2)

}

func SumShortestPath(universe *[][]rune, factor int) int {

	emptyRows := getEmptyRows(universe)
	emptyCols := getEmptyCols(universe)

	galaxies := getGalaxies(universe)

	sum := 0
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			wg.Add(1)
			go func(a, b Galaxy, emptyRows, emptyCols *[]int, factor int, s *int) {
				defer wg.Done()
				defer mutex.Unlock()
				d := getMinumDistanceGalaxies(a, b, emptyRows, emptyCols, factor)
				mutex.Lock()
				*s += d
			}(galaxies[i], galaxies[j], &emptyRows, &emptyCols, factor, &sum)
		}
	}

	wg.Wait()

	return sum
}
