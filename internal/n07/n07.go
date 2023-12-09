package n07

import (
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Hand string
	Bid  int
}

var score = map[string]int{
	"high-hand":       1,
	"one-pair":        2,
	"two-pair":        3,
	"three-of-a-kind": 4,
	"full-house":      5,
	"four-of-a-kind":  6,
	"five-of-a-kind":  7,
}

var cardsList = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func ParseGames(input string) *[]Game {
	games := []Game{}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		data := strings.Split(l, " ")
		bid, _ := strconv.Atoi(data[1])
		games = append(games, Game{
			Hand: data[0],
			Bid:  bid,
		})
	}
	return &games
}

func GetScore(game *Game) int {
	cards := map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}

	for _, c := range game.Hand {
		cards[string(c)]++
	}

	// Five of a kind
	for _, n := range cards {
		if n == 5 {
			return score["five-of-a-kind"]
		}
	}

	// Four of a kind
	for _, n := range cards {
		if n == 4 {
			return score["four-of-a-kind"]
		}
	}

	// Full House
	for i, n := range cards {
		if n == 3 {
			for j, n := range cards {
				if i == j {
					continue
				}
				if n == 2 {
					return score["full-house"]
				}
			}
			// Three of a kind
			return score["three-of-a-kind"]
		}
	}

	// Two Pair
	for i, n := range cards {
		if n == 2 {
			for j, n := range cards {
				if i == j {
					continue
				}
				if n == 2 {
					return score["two-pair"]
				}
			}
			// One Pair
			return score["one-pair"]
		}
	}

	// High Card
	return score["high-hand"]
}

func OrderGames(games *[]Game) {
	sort.SliceStable(*games, func(i, j int) bool {
		// Sort by score
		scoreI, scoreJ := GetScore(&(*games)[i]), GetScore(&(*games)[j])

		// When the score is the same, sort by hand
		if scoreI == scoreJ {

			// Wins the first hand that has an higher card value
			for k := 0; k < len((*games)[i].Hand); k++ {
				cI, cJ := cardsList[string((*games)[i].Hand[k])], cardsList[string((*games)[j].Hand[k])]
				if cI == cJ {
					continue
				}
				return cI < cJ
			}

			// They're the same
			return false
		}

		// Wins the game with higher score
		return scoreI < scoreJ
	})
}
