package n05

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var mapRe = regexp.MustCompile(`^.* map$`)

type Map struct {
	destination int
	source      int
	length      int
}

type Interval struct {
	Min, Max int
}

func (i *Interval) Overlap(j *Interval) (inside []Interval, outside []Interval) {
	inside, outside = []Interval{}, []Interval{}

	if j.Max < i.Min || j.Min > i.Max {
		outside = append(outside, Interval{Min: j.Min, Max: j.Max})
		return inside, outside
	}

	if j.Min < i.Min {
		outside = append(outside, Interval{Min: j.Min, Max: i.Min - 1})
		if j.Max <= i.Max {
			inside = append(inside, Interval{Min: i.Min, Max: j.Max})
		} else {
			inside = append(inside, Interval{Min: i.Min, Max: i.Max})
			outside = append(outside, Interval{Min: i.Max + 1, Max: j.Max})
		}
	} else if j.Min >= i.Min && j.Min <= i.Max {
		if j.Max <= i.Max {
			inside = append(inside, Interval{Min: j.Min, Max: j.Max})
		} else {
			inside = append(inside, Interval{Min: j.Min, Max: i.Max})
			outside = append(outside, Interval{Min: i.Max + 1, Max: j.Max})
		}
	}
	return inside, outside
}

type Almanac struct {
	seeds                 [][]int
	seedToSoilMap         []Map
	soilToFertilizerMap   []Map
	fertilizerToWaterMap  []Map
	waterToLightMap       []Map
	lightToTemperature    []Map
	temperatureToHumidity []Map
	humidityToLocation    []Map
}

func NewAlmanac(data string, rangeMode bool) *Almanac {
	almanac := &Almanac{}
	lines := strings.Split(data, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		// fmt.Printf("Line %d: %s\n", i, line)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		keyValue := strings.Split(line, ":")
		key := strings.TrimSpace(keyValue[0])

		// fmt.Printf("Key: %s\n", key)

		if key == "seeds" {
			seedsStr := strings.Split(strings.TrimSpace(keyValue[1]), " ")
			seeds := [][]int{}
			if rangeMode {
				for k := 0; k < len(seedsStr); k += 2 {
					seed, _ := strconv.Atoi(seedsStr[k])
					seedRange, _ := strconv.Atoi(seedsStr[k+1])
					seeds = append(seeds, []int{seed, seedRange})
				}
			} else {
				for _, seedStr := range seedsStr {
					seed, _ := strconv.Atoi(seedStr)
					seeds = append(seeds, []int{seed})
				}
			}
			almanac.seeds = seeds
		} else if mapRe.MatchString(key) {
			j := 1
			for i+j < len(lines) && strings.TrimSpace(lines[i+j]) != "" {
				// fmt.Printf("Line %d: %s\n", i+j, lines[i+j])
				valuesStr := strings.Split(strings.TrimSpace(lines[i+j]), " ")
				destination, _ := strconv.Atoi(valuesStr[0])
				source, _ := strconv.Atoi(valuesStr[1])
				length, _ := strconv.Atoi(valuesStr[2])
				m := Map{
					destination: destination,
					source:      source,
					length:      length,
				}
				switch key {
				case "seed-to-soil map":
					almanac.seedToSoilMap = append(almanac.seedToSoilMap, m)
				case "soil-to-fertilizer map":
					almanac.soilToFertilizerMap = append(almanac.soilToFertilizerMap, m)
				case "fertilizer-to-water map":
					almanac.fertilizerToWaterMap = append(almanac.fertilizerToWaterMap, m)
				case "water-to-light map":
					almanac.waterToLightMap = append(almanac.waterToLightMap, m)
				case "light-to-temperature map":
					almanac.lightToTemperature = append(almanac.lightToTemperature, m)
				case "temperature-to-humidity map":
					almanac.temperatureToHumidity = append(almanac.temperatureToHumidity, m)
				case "humidity-to-location map":
					almanac.humidityToLocation = append(almanac.humidityToLocation, m)
				}
				j++
			}
			i = i + j
		} else {
			fmt.Printf("Invalid key: %s\n", key)
		}
	}

	return almanac
}

func getValuesFromMaps(maps *[]Map, key int) []int {

	values := []int{}

	for _, m := range *maps {
		if key < m.source || key > m.source+m.length {
			continue
		}
		offset := key - m.source
		values = append(values, m.destination+offset)
	}

	if len(values) == 0 {
		values = append(values, key)
	}

	return values
}

func getRangesFromMaps(maps *[]Map, min int, max int) [][]int {

	intervals := [][]int{}
	currInt := Interval{min, max}
	atLeastOne := false
	for _, m := range *maps {
		sourceInt := Interval{m.source, m.source + m.length}
		inside, _ := sourceInt.Overlap(&currInt)
		if len(inside) > 0 {
			atLeastOne = true
			offset := m.destination - m.source
			intervals = append(intervals, []int{inside[0].Min, inside[0].Max, offset})
		}
		// fmt.Printf("Current: %v, Source: %v, Inside: %v, Outside: %v\n", currInt, sourceInt, inside, outside)
	}

	if !atLeastOne {
		return [][]int{{min, max}}
	}

	minVal, maxVal := intervals[0][0], intervals[0][1]
	values := [][]int{{intervals[0][0] + intervals[0][2], intervals[0][1] + intervals[0][2]}}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < minVal {
			minVal = intervals[i][0]
		}
		if intervals[i][1] > maxVal {
			maxVal = intervals[i][1]
		}
		values = append(values, []int{intervals[i][0] + intervals[i][2], intervals[i][1] + intervals[i][2]})
	}

	if minVal > min {
		values = append(values, []int{min, minVal})
	}
	if maxVal < max {
		values = append(values, []int{maxVal, max})
	}

	return values
}

func processSingleSeed(a *Almanac, seed int) int {

	min := math.MaxInt

	soils := getValuesFromMaps(&a.seedToSoilMap, seed)

	for _, s := range soils {
		fertilizers := getValuesFromMaps(&a.soilToFertilizerMap, s)

		for _, f := range fertilizers {
			waters := getValuesFromMaps(&a.fertilizerToWaterMap, f)

			for _, w := range waters {
				lights := getValuesFromMaps(&a.waterToLightMap, w)

				for _, l := range lights {
					temperatures := getValuesFromMaps(&a.lightToTemperature, l)

					for _, t := range temperatures {

						humidities := getValuesFromMaps(&a.temperatureToHumidity, t)

						for _, h := range humidities {

							locations := getValuesFromMaps(&a.humidityToLocation, h)

							for _, loc := range locations {

								// fmt.Printf("Seed: %d, Soil: %d, Fertilizer: %d, Water: %d, Light: %d, Temperature: %d, Humidity: %d, Location: %d\n", seed, s, f, w, l, t, h, l)

								if loc < min {
									min = loc
								}
							}
						}
					}
				}
			}
		}
	}

	return min
}

func (m *Almanac) GetMinimumLocation() int {

	min := math.MaxInt

	for _, seed := range m.seeds {

		if len(seed) == 1 {
			m := processSingleSeed(m, seed[0])
			if m < min {
				min = m
			}

		} else {
			m := processRangeSeed(m, seed[0], seed[1])
			// fmt.Printf("Seed: %d, Min: %d\n", seed[0], m)
			if m < min {
				min = m
			}
		}
	}

	return min
}

func processRangeSeed(a *Almanac, seed int, seedRange int) int {

	min := math.MaxInt

	soils := getRangesFromMaps(&a.seedToSoilMap, seed, seed+seedRange)

	// fmt.Printf("Soils: %v\n", soils)

	for _, s := range soils {
		fertilizers := getRangesFromMaps(&a.soilToFertilizerMap, s[0], s[1])

		// fmt.Printf("Fertilizers: %v\n", fertilizers)

		for _, f := range fertilizers {
			waters := getRangesFromMaps(&a.fertilizerToWaterMap, f[0], f[1])

			for _, w := range waters {
				lights := getRangesFromMaps(&a.waterToLightMap, w[0], w[1])

				for _, l := range lights {
					temperatures := getRangesFromMaps(&a.lightToTemperature, l[0], l[1])

					for _, t := range temperatures {

						humidities := getRangesFromMaps(&a.temperatureToHumidity, t[0], t[1])

						for _, h := range humidities {

							locations := getRangesFromMaps(&a.humidityToLocation, h[0], h[1])

							for _, loc := range locations {

								// fmt.Printf("Seed: %d, Soil: %d, Fertilizer: %d, Water: %d, Light: %d, Temperature: %d, Humidity: %d, Location: %d\n", seed, s, f, w, l, t, h, l)
								// fmt.Printf("Loc: %d\n", loc[0])
								if loc[0] < min {
									min = loc[0]
								}
							}
						}
					}
				}
			}
		}
	}

	return min
}
