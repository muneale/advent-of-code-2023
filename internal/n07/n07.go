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

var ReversedScore = map[int]string{
	1: "high-hand",
	2: "one-pair",
	3: "two-pair",
	4: "three-of-a-kind",
	5: "full-house",
	6: "four-of-a-kind",
	7: "five-of-a-kind",
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

var cardsListWithJolly = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func ParseGames(input string) *[]Game {
	games := []Game{}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		data := strings.Split(l, " ")
		bid, _ := strconv.Atoi(strings.TrimSpace(data[1]))
		games = append(games, Game{
			Hand: data[0],
			Bid:  bid,
		})
	}
	return &games
}

func GetScore(game *Game, jolly bool) int {
	cards := map[string]int{}

	for _, c := range game.Hand {
		if _, ok := cards[string(c)]; !ok {
			cards[string(c)] = 0
		}
		cards[string(c)]++
	}

	jokers := 0
	if jolly {
		if val, ok := cards["J"]; ok {
			jokers = val
			delete(cards, "J")
		}
	}

	combos := []int{}
	for _, n := range cards {
		combos = append(combos, n)
	}
	sort.Slice(combos, func(i, j int) bool {
		return combos[i] > combos[j]
	})

	if jokers == 5 {
		return score["five-of-a-kind"]
	}

	combos[0] += jokers
	if combos[0] == 5 {
		return score["five-of-a-kind"]
	}

	if combos[0] == 4 {
		return score["four-of-a-kind"]
	}

	if combos[0] == 3 {
		if combos[1] == 2 {
			return score["full-house"]
		}
		return score["three-of-a-kind"]
	}

	if combos[0] == 2 {
		if combos[1] == 2 {
			return score["two-pair"]
		}
		return score["one-pair"]
	}

	return score["high-hand"]
}

func OrderGames(games *[]Game, jolly bool) {
	sort.SliceStable(*games, func(i, j int) bool {
		// Sort by score
		scoreI, scoreJ := GetScore(&(*games)[i], jolly), GetScore(&(*games)[j], jolly)

		// When the score is the same, sort by hand
		if scoreI == scoreJ {

			for k := 0; k < len((*games)[i].Hand); k++ {
				cI, cJ := 0, 0
				if jolly {
					cI, cJ = cardsListWithJolly[string((*games)[i].Hand[k])], cardsListWithJolly[string((*games)[j].Hand[k])]

				} else {
					cI, cJ = cardsList[string((*games)[i].Hand[k])], cardsList[string((*games)[j].Hand[k])]
				}
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
