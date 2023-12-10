package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/muneale/advent-of-code-2023/internal/n01"
	"github.com/muneale/advent-of-code-2023/internal/n02"
	"github.com/muneale/advent-of-code-2023/internal/n03"
	"github.com/muneale/advent-of-code-2023/internal/n04"
	"github.com/muneale/advent-of-code-2023/internal/n05"
	"github.com/muneale/advent-of-code-2023/internal/n06"
	"github.com/muneale/advent-of-code-2023/internal/n07"
	"github.com/muneale/advent-of-code-2023/internal/n08"
	"github.com/muneale/advent-of-code-2023/internal/n09"
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
	case "06":
		runDay06()
	case "07":
		runDay07()
	case "08":
		runDay08()
	case "09":
		runDay09()
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

func runDay06() {

	input, _ := os.ReadFile("./input/n06.txt")

	dts := n06.ParseDistanceAndTime(string(input))

	wins := 1
	for _, dt := range *dts {
		// fmt.Printf("Distance: %v | Time: %v | Wins: %v\n", dt.Distance, dt.Time, n06.WinningCombinations(dt))
		wins *= n06.WinningCombinations(dt)
	}

	fmt.Printf("Part 1: %d\n", wins)

	realDistance, realTime := "", ""
	for _, dt := range *dts {
		realDistance += strconv.Itoa(dt.Distance)
		realTime += strconv.Itoa(dt.Time)
	}

	d, _ := strconv.Atoi(realDistance)
	t, _ := strconv.Atoi(realTime)
	dt := n06.DistanceAndTime{
		Distance: d,
		Time:     t,
	}

	wins = n06.WinningCombinations(dt)

	fmt.Printf("Part 2: %d\n", wins)
}

func runDay07() {

	input, _ := os.ReadFile("./input/n07.txt")

	games := n07.ParseGames(string(input))

	n07.OrderGames(games, false)

	totalWinnings := 0
	for i, g := range *games {
		totalWinnings += g.Bid * (i + 1)
	}

	fmt.Printf("Part 1: %v\n", totalWinnings)

	gamesWithJolly := n07.ParseGames(string(input))

	n07.OrderGames(gamesWithJolly, true)

	totalWinnings = 0
	for i, g := range *gamesWithJolly {
		// fmt.Printf("Game: %v | Bid: %d | Score: %s | Rank: %d\n", g.Hand, g.Bid, n07.ReversedScore[n07.GetScoreWithJolly(&g)], i+1)
		totalWinnings += g.Bid * (i + 1)
	}

	fmt.Printf("Part 2: %v\n", totalWinnings)
}

func runDay08() {

	input, _ := os.ReadFile("./input/n08.txt")

	directions, nodes := n08.ParseInput(string(input))

	steps := n08.StepsToGoal(directions, nodes)
	fmt.Printf("Part 1: %d\n", steps)

	steps = n08.StepsToGoalMultidimensional(directions, nodes)
	fmt.Printf("Part 2: %d\n", steps)

}

func runDay09() {

	input, _ := os.ReadFile("./input/n09.txt")
	sequences := n09.ParseInput(string(input))

	total := 0
	for _, s := range sequences {
		predicted := n09.PredictValue(s, false)
		// fmt.Printf("Sequence: %v | Predicted: %d\n", s, predicted)
		total += predicted
	}

	fmt.Printf("Part 1: %d\n", total)

	total = 0
	for _, s := range sequences {
		predicted := n09.PredictValue(s, true)
		// fmt.Printf("Sequence: %v | Predicted: %d\n", s, predicted)
		total += predicted
	}

	fmt.Printf("Part 2: %d\n", total)
}
