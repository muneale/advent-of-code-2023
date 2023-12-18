package n13

import (
	"math"
	"strings"
)

type Mirror struct {
	Rows, Cols []int
}

func ParseInput(s string) *[]Mirror {
	mirrors := []Mirror{}

	var sb strings.Builder
	lines := strings.Split(s, "\n")
	for i, line := range lines {

		if len(line) == 0 || i == len(lines)-1 {
			mirror := parseMirror(sb.String())
			mirrors = append(mirrors, *mirror)
			sb.Reset()
			continue
		}

		sb.WriteString(line + "\n")
	}

	return &mirrors
}

func parseMirror(s string) *Mirror {
	mirror := Mirror{}
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]

	for i := 0; i < len(lines); i++ {
		rowCount := 0
		for j, c := range lines[i] {
			if c == '#' {
				rowCount += int(math.Pow(2, float64(len(lines[i])-1-j)))
			}
		}
		mirror.Rows = append(mirror.Rows, rowCount)
	}

	for i := 0; i < len(lines[0]); i++ {
		colCount := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				colCount += int(math.Pow(2, float64(len(lines)-1-j)))
			}
		}
		mirror.Cols = append(mirror.Cols, colCount)
	}

	return &mirror
}

func SummarizePattern(mirrors *[]Mirror, withSmudges bool) int {

	row, col := 0, 0
	for _, mirror := range *mirrors {

		// Check symmetric rows
		found := false
		for i := 0; i < len(mirror.Rows)-1 && !found; i++ {
			if withSmudges {
				if isReflectedWithSmudge(i, i+1, &mirror.Rows, false) {
					if !isReflected(i, i+1, &mirror.Rows) {
						row += i + 1
						found = true
					}
				}
			} else {
				if isReflected(i, i+1, &mirror.Rows) {
					row += i + 1
					found = true
				}
			}
		}

		// Check symmetric cols
		for i := 0; i < len(mirror.Cols)-1 && !found; i++ {
			if withSmudges {
				if isReflectedWithSmudge(i, i+1, &mirror.Cols, false) {
					if !isReflected(i, i+1, &mirror.Cols) {
						col += i + 1
						found = true
					}
				}
			} else {
				if isReflected(i, i+1, &mirror.Cols) {
					col += i + 1
					found = true
				}
			}
		}
	}

	return col + row*100
}

func isReflected(i, j int, nums *[]int) bool {
	if i < 0 || j >= len(*nums) {
		return true
	}

	if (*nums)[i] == (*nums)[j] {
		return isReflected(i-1, j+1, nums)
	}

	return false
}

func isReflectedWithSmudge(i, j int, nums *[]int, alreadySmudged bool) bool {
	if i < 0 || j >= len(*nums) {
		return true
	}

	if (*nums)[i] == (*nums)[j] {
		return isReflectedWithSmudge(i-1, j+1, nums, alreadySmudged)
	}

	if alreadySmudged {
		return false
	}

	// Check if the XOR is a power of 2
	xor := (*nums)[i] ^ (*nums)[j]
	isPowerOfTwo := xor&(xor-1) == 0
	if isPowerOfTwo {
		return isReflectedWithSmudge(i-1, j+1, nums, true)
	}
	return false
}
