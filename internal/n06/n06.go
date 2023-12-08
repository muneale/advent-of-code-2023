package n06

import (
	"regexp"
	"strconv"
	"strings"
)

var numRe = regexp.MustCompile(`([0-9]+)`)

type DistanceAndTime struct {
	Distance int
	Time     int
}

func ParseDistanceAndTime(data string) *[]DistanceAndTime {
	dts := []DistanceAndTime{}

	lines := strings.Split(data, "\n")

	unparsedTimes, unparsedDistances := lines[0], lines[1]

	times := numRe.FindAllString(unparsedTimes, -1)
	distances := numRe.FindAllString(unparsedDistances, -1)

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		dts = append(dts, DistanceAndTime{
			Distance: distance,
			Time:     time,
		})
	}

	return &dts
}

func WinningCombinations(dt DistanceAndTime) int {
	wins := -1
	// wins := 0

	// fmt.Printf("Distance: %d | Time: %d\n", dt.Distance, dt.Time)
	for i := 0; i <= dt.Time; i++ {
		if wins > -1 {
			break
		}
		distance := i * (dt.Time - i)
		// fmt.Printf("Speed: %d | Remaining Time: %d | Target Distance %d |  Distance: %d | Wins: %v\n", i, dt.Time-i, dt.Distance, distance, distance > dt.Distance)
		if distance > dt.Distance {
			// wins = dt.Time - i
			if dt.Time%2 == 0 {
				wins = dt.Time - 2*i + 1
			} else {
				wins = (dt.Time + 1) - 2*i
			}
		}
	}

	return wins
}
