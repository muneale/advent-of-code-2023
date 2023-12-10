package n09

import (
	"strconv"
	"strings"
)

func ParseInput(data string) [][]int {
	lines := strings.Split(data, "\n")
	resultStr := [][]string{}
	for i := range lines {
		resultStr = append(resultStr, strings.Split(strings.TrimSpace(lines[i]), " "))
	}

	result := [][]int{}
	for _, r := range resultStr {
		res := []int{}
		for _, v := range r {
			i, _ := strconv.Atoi(v)
			res = append(res, i)
		}
		result = append(result, res)
	}
	return result
}

func PredictValue(data []int, prev bool) int {

	if len(data) == 0 {
		return 0
	}

	allZeros := true
	for i := range data {
		if data[i] != 0 {
			allZeros = false
			break
		}
	}
	if allZeros {
		return 0
	}

	// Compute the new sequence
	newSequence := []int{}
	for i := 1; i < len(data); i++ {
		newSequence = append(newSequence, data[i]-data[i-1])
	}

	if prev {
		return data[0] - PredictValue(newSequence, true)
	}
	return data[len(data)-1] + PredictValue(newSequence, false)
}
