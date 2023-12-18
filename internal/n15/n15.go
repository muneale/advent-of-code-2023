package n15

import (
	"regexp"
	"strconv"
	"strings"
)

var removeRE = regexp.MustCompile(`^(.*)\-$`)
var addRE = regexp.MustCompile(`^(.*)=(.*)$`)

type Lens struct {
	label string
	value int
}

func ParseInput(data string) *[]string {
	tokens := strings.Split(data, ",")
	return &tokens
}

func hash(token string) int {
	val := 0
	for _, c := range token {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}

func HashSum(tokens *[]string) int {
	sum := 0
	for _, token := range *tokens {
		sum += hash(token)
	}
	return sum
}

func hashmap(tokens *[]string) *map[int][]Lens {
	m := map[int][]Lens{}
	for _, token := range *tokens {
		if val := removeRE.FindStringSubmatch(token); val != nil {
			k := hash(val[1])
			if _, ok := m[k]; !ok {
				continue
			}
			for i := 0; i < len(m[k]); i++ {
				if m[k][i].label == val[1] {
					m[k] = append(m[k][:i], m[k][i+1:]...)
					break
				}
			}
		} else if val := addRE.FindStringSubmatch(token); val != nil {
			k := hash(val[1])
			if _, ok := m[k]; !ok {
				m[k] = []Lens{}
			}
			value, _ := strconv.Atoi(val[2])

			// Check if value already exists
			found := false
			for i := 0; i < len(m[k]); i++ {
				if m[k][i].label == val[1] {
					m[k][i].value = value
					found = true
					break
				}
			}
			if !found {
				m[k] = append(m[k], Lens{val[1], value})
			}
		}
	}
	return &m
}

func HashmapSum(tokens *[]string) int {
	m := hashmap(tokens)
	sum := 0
	for i := 0; i < 256; i++ {
		if _, ok := (*m)[i]; !ok {
			continue
		}
		v := (*m)[i]
		for j := 0; j < len(v); j++ {
			sum += (i + 1) * (j + 1) * v[j].value
		}
	}
	return sum
}
