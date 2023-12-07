package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/muneale/advent-of-code-2023/internal/n01"
	"github.com/muneale/advent-of-code-2023/internal/n02"
	"github.com/muneale/advent-of-code-2023/internal/n03"
	"github.com/muneale/advent-of-code-2023/internal/n04"
	"github.com/muneale/advent-of-code-2023/internal/n05"
)

func main() {

	day := os.Args[1]

	switch day {
	case "01":
		runDay01()
	case "02":
		runDay02()
	case "03":
		runDay03()
	case "04":
		runDay04()
	case "05":
		runDay05()
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

func runDay03() {
	input, _ := os.ReadFile("./input/n03.txt")

	m := n03.GetMatrix(string(input))
	sum := m.GetEngineSchematicSum()
	fmt.Printf("Part 1: %d\n", sum)

	m = n03.GetMatrix(string(input))
	sum = m.GetSchamticGearRatioSum()
	fmt.Printf("Part 2: %d\n", sum)
}

func runDay04() {

	input, _ := os.ReadFile("./input/n04.txt")

	c := n04.GetCards(string(input))

	sum := 0
	for i := range c {
		// c[i].Print()
		sum += c[i].Points()
	}

	fmt.Printf("Part 1: %v\n", sum)

	copies := n04.CountCardCopies(c)

	fmt.Printf("Part 2: %v\n", copies)
}

func runDay05() {

	input, _ := os.ReadFile("./input/n05.txt")

	almanac := n05.NewAlmanac(string(input), false)
	min := almanac.GetMinimumLocation()
	fmt.Printf("Part 1: %v\n", min)

	almanac = n05.NewAlmanac(string(input), true)
	min = almanac.GetMinimumLocation()
	fmt.Printf("Part 2: %v\n", min)
}
