package n19

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var ruleRE = regexp.MustCompile(`^(.+)\{(.*)\}$`)
var partRE = regexp.MustCompile(`[x|m|a|s]=(\d+)`)
var conditionRE = regexp.MustCompile(`(x|m|a|s)(<|>)(\d+)`)

type part struct {
	x, m, a, s int
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func Run() {
	wd, _ := os.Getwd()
	filepath := filepath.Join(wd, "input/n19.txt")

	input := getLines(filepath)

	rules, parts := getRulesAndParts(&input)

	sum := 0
	for _, part := range *parts {
		if isPartValid(part, rules) {
			sum += part.x + part.m + part.a + part.s
		}
	}

	fmt.Printf("Part 1: %d\n", sum)

}

func getRulesAndParts(input *[]string) (*map[string]string, *[]part) {

	rules := map[string]string{}
	parts := []part{}

	parseRules := true
	for _, line := range *input {
		if line == "" {
			parseRules = false
			continue
		}

		if parseRules {
			matches := ruleRE.FindStringSubmatch(line)
			key := matches[1]
			if _, ok := rules[key]; !ok {
				rules[key] = matches[2]
			} else {
				fmt.Printf("Duplicate rule for key %s\n", key)
			}
			continue
		} else {
			matches := partRE.FindAllStringSubmatch(line, -1)
			x, _ := strconv.Atoi(matches[0][1])
			m, _ := strconv.Atoi(matches[1][1])
			a, _ := strconv.Atoi(matches[2][1])
			s, _ := strconv.Atoi(matches[3][1])
			parts = append(parts, part{
				x: x,
				m: m,
				a: a,
				s: s,
			})
		}
	}

	return &rules, &parts
}

func getNextRule(rule string, part part) string {
	rules := strings.Split(rule, ",")
	for i := 0; i < len(rules)-1; i++ {
		d := strings.Split(rules[i], ":")
		condition := d[0]
		nextRule := d[1]
		matches := conditionRE.FindStringSubmatch(condition)
		val, _ := strconv.Atoi(matches[3])
		if matches[2] == ">" {
			if matches[1] == "x" {
				if part.x > val {
					return nextRule
				}
			}
			if matches[1] == "m" {
				if part.m > val {
					return nextRule
				}
			}
			if matches[1] == "a" {
				if part.a > val {
					return nextRule
				}
			}
			if matches[1] == "s" {
				if part.s > val {
					return nextRule
				}
			}
		} else {
			if matches[1] == "x" {
				if part.x < val {
					return nextRule
				}
			}
			if matches[1] == "m" {
				if part.m < val {
					return nextRule
				}
			}
			if matches[1] == "a" {
				if part.a < val {
					return nextRule
				}
			}
			if matches[1] == "s" {
				if part.s < val {
					return nextRule
				}
			}
		}
	}
	return rules[len(rules)-1]
}

func isPartValid(part part, rules *map[string]string) bool {
	nextRule := "in"

	for nextRule != "A" && nextRule != "R" {
		rule := (*rules)[nextRule]
		nextRule = getNextRule(rule, part)
	}

	if nextRule == "A" {
		return true
	} else {
		return false
	}
}
