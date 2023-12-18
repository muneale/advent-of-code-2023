package n14

import "strings"

func ParseInput(data string) *[][]rune {
	gameMap := [][]rune{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		gameMap = append(gameMap, row)
	}
	return &gameMap
}

func rotate(g *[][]rune) *[][]rune {
	R := len(*g)
	C := len((*g)[0])

	NG := make([][]rune, R)
	for i := 0; i < R; i++ {
		NG[i] = make([]rune, C)
	}

	for c := 0; c < C; c++ {
		for r := 0; r < R; r++ {
			NG[c][R-1-r] = (*g)[r][c]
		}
	}
	return &NG
}

func roll(g *[][]rune) *[][]rune {
	R := len(*g)
	C := len((*g)[0])

	for c := 0; c < C; c++ {
		for rn := 0; rn < R; rn++ {
			for r := 0; r < R; r++ {
				if (*g)[r][c] == 'O' && r > 0 && (*g)[r-1][c] == '.' {
					(*g)[r][c] = '.'
					(*g)[r-1][c] = 'O'
				}
			}
		}
	}

	return g
}

func score(g *[][]rune) int {
	ans := 0
	R := len(*g)
	C := len((*g)[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if (*g)[r][c] == 'O' {
				ans += len((*g)) - r
			}
		}
	}
	return ans
}

func GetLoadNorth(g *[][]rune) int {
	roll(g)
	return score(g)
}

func getSequence(values []int) (frequency, offset int, ok bool) {
	if len(values) > 10 {
		var reps []int
		for i := 0; i < len(values)-1; i++ {
			if values[i] == values[len(values)-1] {
				reps = append(reps, i)
				if len(reps) == 3 {
					if reps[2]-reps[1] == reps[1]-reps[0] && reps[2]-reps[1] > 1 {
						ok = true
						for i := 0; i < reps[1]-reps[0]; i++ {
							if values[reps[0]+i] != values[reps[1]+i] {
								ok = false
							}
						}
						if ok {
							frequency = reps[2] - reps[1]
							offset = reps[0]
						}
					}
				}
			}
		}
	}
	return frequency, offset, ok
}

func GetLoadCycle(g *[][]rune, cycles int) int {
	result := 0
	results := []int{}
	for c := 0; c < cycles; c++ {
		for i := 0; i < 4; i++ {
			g = roll(g)
			g = rotate(g)
		}
		results = append(results, score(g))

		if frequency, offset, ok := getSequence(results); ok {
			mod := (cycles - offset) % frequency
			result = results[offset+mod-1]
			break
		}

	}

	return result
}
