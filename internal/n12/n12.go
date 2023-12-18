package n12

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Line struct {
	cl []int // combination line
	nv []int // numeric values
}

func ParseString2(s string) (l Line, err error) {
	r := strings.Split(s, " ")
	if len(r) != 2 {
		return l, errors.New("invalid input")
	}
	l.cl = make([]int, (len(r[0])+1)*5)
	l.cl[0] = 1 // adding padding
	for j := 0; j < 5; j++ {
		for i, v := range r[0] {
			switch v {
			case '.':
				l.cl[i+1+j*(len(r[0])+1)] = 1
			case '#':
				l.cl[i+1+j*(len(r[0])+1)] = 2
			case '?':
				l.cl[i+1+j*(len(r[0])+1)] = 0
			default:
				return l, errors.New("invalid input")
			}
		}
		if j < 4 {
			l.cl[len(r[0])+1+j*(len(r[0])+1)] = 0
		}
	}
	r1 := strings.Split(r[1], ",")
	l.nv = make([]int, len(r1)*5)
	var n int
	for j := 0; j < 5; j++ {
		for i, v := range r1 {
			if n, err = strconv.Atoi(v); err != nil {
				return l, errors.New("invalid input")
			}
			l.nv[i+j*len(r1)] = n
		}
	}
	return l, err
}

func ParseString(s string) (l Line, err error) {
	r := strings.Split(s, " ")
	if len(r) != 2 {
		return l, errors.New("invalid input")
	}
	l.cl = make([]int, len(r[0])+1)
	l.cl[0] = 1 // adding padding
	for i, v := range r[0] {
		switch v {
		case '.':
			l.cl[i+1] = 1
		case '#':
			l.cl[i+1] = 2
		case '?':
			l.cl[i+1] = 0
		default:
			return l, errors.New("invalid input")
		}
	}
	r1 := strings.Split(r[1], ",")
	l.nv = make([]int, len(r1))
	var n int
	for i, v := range r1 {
		if n, err = strconv.Atoi(v); err != nil {
			return l, errors.New("invalid input")
		}
		l.nv[i] = n
	}
	return l, err
}

func itFits(pattern []int, pos int, cl []int) bool {
	if len(pattern)+pos > len(cl) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if pattern[i] != cl[pos+i] && cl[pos+i] != 0 {
			return false
		}
	}
	return true
}

var cache map[string]int

func convertToKey(i int, a []int, j int) (key string) {
	out := make([]rune, len(a))
	s := fmt.Sprintf("%d,%d,", i, j)
	for idx, v := range a {
		out[idx] = rune(v)
	}
	return s + string(out)
}

func getFitting(combinationList []int, startPos int, numericValueList []int, index int) (result int) {
	// get current pattern
	pattern := make([]int, numericValueList[index]+1)
	pattern[0] = 1
	for k := 0; k < numericValueList[index]; k++ {
		pattern[k+1] = 2
	}
	// find next valid match
	for j := startPos; j < len(combinationList); j++ {
		if itFits(pattern, j, combinationList) {
			if index == len(numericValueList)-1 {
				valid := true
				for h := j + len(pattern); h < len(combinationList); h++ {
					if combinationList[h] == 2 {
						valid = false
						break
					}
				}
				if valid {
					result += 1
				}
			} else {
				if value, ok := cache[convertToKey(j, pattern, index+1)]; ok {
					result += value
				} else {
					val := getFitting(combinationList, j+len(pattern), numericValueList, index+1)
					cache[convertToKey(j, pattern, index+1)] = val
					result += val
				}
			}
		}
		for h := startPos; h < j+1; h++ {
			if combinationList[h] == 2 {
				return result
			}
		}
	}
	return result
}

func GetCombinations(s string, parseString func(s string) (l Line, err error)) (cs int, err error) {
	var l Line
	if l, err = parseString(s); err != nil {
		return cs, err
	}
	cache = make(map[string]int)
	return getFitting(l.cl, 0, l.nv, 0), err
}

func GetSumCombinations(s []string, parseString func(s string) (l Line, err error)) (sum int, err error) {
	var v int
	for _, c := range s {
		if v, err = GetCombinations(c, parseString); err != nil {
			return sum, err
		}
		sum += v
	}
	return sum, err
}
