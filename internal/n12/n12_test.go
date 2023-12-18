package n12_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muneale/advent-of-code-2023/internal/n12"
)

func TestGetCombinations(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected int
		parser   func(s string) (l n12.Line, err error)
	}{
		{
			desc:     "part1: expected ???.### 1,1,3 - 1 arrangement",
			input:    "???.### 1,1,3",
			parser:   n12.ParseString,
			expected: 1,
		},
		{
			desc:     "part2: expected ???.### 1,1,3 - 1 arrangement",
			input:    "???.### 1,1,3",
			parser:   n12.ParseString2,
			expected: 1,
		},
		{
			desc:     "part1: expected .??..??...?##. 1,1,3 - 4 arrangements",
			input:    ".??..??...?##. 1,1,3",
			parser:   n12.ParseString,
			expected: 4,
		},
		{
			desc:     "part2: expected .??..??...?##. 1,1,3 - 16384 arrangements",
			input:    ".??..??...?##. 1,1,3",
			parser:   n12.ParseString2,
			expected: 16384,
		},
		{
			desc:     "part1: expected ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement",
			input:    "?#?#?#?#?#?#?#? 1,3,1,6",
			parser:   n12.ParseString,
			expected: 1,
		},
		{
			desc:     "part2: expected ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement",
			input:    "?#?#?#?#?#?#?#? 1,3,1,6",
			parser:   n12.ParseString2,
			expected: 1,
		},
		{
			desc:     "part1: expected ????.#...#... 4,1,1 - 1 arrangement",
			input:    "????.#...#... 4,1,1",
			parser:   n12.ParseString,
			expected: 1,
		},
		{
			desc:     "part2: expected ????.#...#... 4,1,1 - 16 arrangement",
			input:    "????.#...#... 4,1,1",
			parser:   n12.ParseString2,
			expected: 16,
		},
		{
			desc:     "part1: expected ????.######..#####. 1,6,5 - 4 arrangements",
			input:    "????.######..#####. 1,6,5",
			parser:   n12.ParseString,
			expected: 4,
		},
		{
			desc:     "part2: expected ????.######..#####. 1,6,5 - 2500 arrangements",
			input:    "????.######..#####. 1,6,5",
			parser:   n12.ParseString2,
			expected: 2500,
		},
		{
			desc:     "part1: expected ?###???????? 3,2,1 - 10 arrangements",
			input:    "?###???????? 3,2,1",
			parser:   n12.ParseString,
			expected: 10,
		},
		{
			desc:     "part2: expected ?###???????? 3,2,1 - 506250 arrangements",
			input:    "?###???????? 3,2,1",
			parser:   n12.ParseString2,
			expected: 506250,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := n12.GetCombinations(tc.input, tc.parser)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetSumCombinations(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		parser   func(string) (n12.Line, error)
		expected int
	}{
		{
			desc: "expected 21 = 1 + 4 + 1 + 1 + 4 + 10",
			input: []string{
				"???.### 1,1,3",
				".??..??...?##. 1,1,3",
				"?#?#?#?#?#?#?#? 1,3,1,6",
				"????.#...#... 4,1,1",
				"????.######..#####. 1,6,5",
				"?###???????? 3,2,1",
			},
			parser:   n12.ParseString,
			expected: 21,
		},
		{
			desc: "expected 525152 = 1 + 16384 + 1 + 16 + 2500 + 506250",
			input: []string{
				"???.### 1,1,3",
				".??..??...?##. 1,1,3",
				"?#?#?#?#?#?#?#? 1,3,1,6",
				"????.#...#... 4,1,1",
				"????.######..#####. 1,6,5",
				"?###???????? 3,2,1",
			},
			parser:   n12.ParseString2,
			expected: 525152,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := n12.GetSumCombinations(tc.input, tc.parser)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
