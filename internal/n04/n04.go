package n04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
}

// Syntax: Card N: W1 W2 W3 W4 W5 | N1 N2 N3 N4 N5 N6 N7 N8
// E.g.: Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
var cardRe = regexp.MustCompile(`Card (.*): (.*) \| (.*)`)

func NewCard(data string) *Card {

	matches := cardRe.FindStringSubmatch(data)

	id, _ := strconv.Atoi(strings.TrimSpace(matches[1]))

	winningNumbersStr := strings.Split(strings.TrimSpace(matches[2]), " ")
	winningNumbers := []int{}
	for _, ws := range winningNumbersStr {
		if strings.TrimSpace(ws) == "" {
			continue
		}
		w, _ := strconv.Atoi(ws)
		winningNumbers = append(winningNumbers, w)
	}

	numbersStr := strings.Split(strings.TrimSpace(matches[3]), " ")
	numbers := []int{}
	for _, ws := range numbersStr {
		if strings.TrimSpace(ws) == "" {
			continue
		}
		w, _ := strconv.Atoi(ws)
		numbers = append(numbers, w)
	}

	card := Card{
		Id:             id,
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
	}

	return &card
}

func (c *Card) Print() {
	fmt.Printf("Card %d: %v | %v\n", c.Id, c.WinningNumbers, c.Numbers)
}

func (c *Card) Matches() int {
	matches := 0
	for _, w := range c.WinningNumbers {
		for _, n := range c.Numbers {
			if n == w {
				matches++
				break
			}
		}
	}
	return matches
}

func (c *Card) Points() int {

	matches := c.Matches()

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches)-1))
}

func GetCards(data string) []*Card {

	cards := []*Card{}

	for _, line := range strings.Split(data, "\n") {
		cards = append(cards, NewCard(line))
	}

	return cards
}

func CountCardCopies(cards []*Card) int {

	mapMatch := map[int]int{}
	copies := len(cards)
	cardCopies := cards
	for len(cardCopies) > 0 {
		c := cardCopies[0]
		// c.Print()
		cardCopies = cardCopies[1:]
		val, ok := mapMatch[c.Id]
		if !ok {
			val = c.Matches()
			mapMatch[c.Id] = val
		}
		// fmt.Printf("Card %d matches %d\n", c.Id, val)
		copies += val
		if val > 0 {
			for i := c.Id; i < c.Id+val; i++ {
				cardCopies = append(cardCopies, cards[i])
			}
		}
	}
	return copies
}
