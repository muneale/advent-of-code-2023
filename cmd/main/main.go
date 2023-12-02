package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/muneale/advent-of-code-2023/internal/n01"
	"github.com/muneale/advent-of-code-2023/internal/n02"
)

func main() {

	day := os.Args[1]

	switch day {
	case "01":
		runDay01()
	case "02":
		runDay02()
	default:
		fmt.Printf("Invalid day: %s\n", day)
	}

}

func runDay01() {
	input, _ := os.ReadFile("./input/n01.txt")
	data := strings.Split(string(input), "\n")

	fmt.Printf("Part 1: %d\n", n01.GetTotalCalibrationValue(data))

	for i := range data {
		n01.TranslateValue(&data[i])
	}

	fmt.Printf("Part 2: %d\n", n01.GetTotalCalibrationValue(data))
}

func runDay02() {

	input, _ := os.ReadFile("./input/n02.txt")

	games := n02.ParseGames(string(input))

	maxCubes := n02.Cubes{
		Blue:  14,
		Green: 13,
		Red:   12,
	}

	sum := n02.GetSumOfPossibleCombinations(maxCubes, games)

	fmt.Printf("Part 1: %d\n", sum)

	sum = n02.GetSumOfPowerMinimumCombinations(games)

	fmt.Printf("Part 2: %d\n", sum)
}
