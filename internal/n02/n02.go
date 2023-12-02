package n02

import (
	"regexp"
	"strconv"
	"strings"
)

type Cubes struct {
	Blue  int
	Green int
	Red   int
}

type Game struct {
	id           int
	combinations []Cubes
}

var gameRegex = regexp.MustCompile("^Game (.*?): (.*)$")

func parseGame(data string) Game {

	matches := gameRegex.FindStringSubmatch(data)
	id, _ := strconv.Atoi(matches[1])

	unparsedGames := strings.Split(matches[2], ";")
	game := Game{
		id:           id,
		combinations: []Cubes{},
	}
	for _, unparsedGame := range unparsedGames {
		unparsedCubes := strings.Split(unparsedGame, ",")
		cubes := Cubes{
			Blue:  0,
			Green: 0,
			Red:   0,
		}
		for _, unparsedCube := range unparsedCubes {
			vals := strings.Split(strings.TrimSpace(unparsedCube), " ")
			switch vals[1] {
			case "blue":
				cubes.Blue, _ = strconv.Atoi(vals[0])
			case "green":
				cubes.Green, _ = strconv.Atoi(vals[0])
			case "red":
				cubes.Red, _ = strconv.Atoi(vals[0])
			}
		}
		game.combinations = append(game.combinations, cubes)
	}

	return game
}

func ParseGames(data string) []Game {

	games := []Game{}

	for _, unparsedGame := range strings.Split(data, "\n") {
		games = append(games, parseGame(unparsedGame))
	}

	return games
}

func GetSumOfPossibleCombinations(maxCubes Cubes, games []Game) int {
	sum := 0

	for _, game := range games {
		validGame := true
		for _, combination := range game.combinations {
			if combination.Blue > maxCubes.Blue || combination.Green > maxCubes.Green || combination.Red > maxCubes.Red {
				validGame = false
				break
			}
		}
		if validGame {
			sum += game.id
		}
	}

	return sum
}

func GetSumOfPowerMinimumCombinations(games []Game) int {
	sum := 0

	for _, game := range games {

		// fmt.Printf("Game %d: %v\n", game.id, game.combinations)

		minCubes := Cubes{
			Blue:  0,
			Green: 0,
			Red:   0,
		}

		for i := 0; i < len(game.combinations); i++ {
			if game.combinations[i].Blue > minCubes.Blue {
				minCubes.Blue = game.combinations[i].Blue
			}
			if game.combinations[i].Green > minCubes.Green {
				minCubes.Green = game.combinations[i].Green
			}
			if game.combinations[i].Red > minCubes.Red {
				minCubes.Red = game.combinations[i].Red
			}
		}

		// fmt.Printf("Game %d: B %d G %d R %d\n", game.id, minCubes.Blue, minCubes.Green, minCubes.Red)

		sum += minCubes.Blue * minCubes.Green * minCubes.Red
	}

	return sum
}
