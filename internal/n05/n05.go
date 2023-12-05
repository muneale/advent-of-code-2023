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

type Almanac struct {
	seeds                 []int
	seedToSoilMap         []Map
	soilToFertilizerMap   []Map
	fertilizerToWaterMap  []Map
	waterToLightMap       []Map
	lightToTemperature    []Map
	temperatureToHumidity []Map
	humidityToLocation    []Map
}

func NewAlmanac(data string) *Almanac {
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
			seeds := []int{}
			for _, seedStr := range seedsStr {
				seed, _ := strconv.Atoi(seedStr)
				seeds = append(seeds, seed)
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

func getValueFromMaps(maps *[]Map, key int) int {

	value := key

	for _, m := range *maps {
		if key < m.source || key > m.source+m.length {
			continue
		}
		offset := key - m.source
		value = m.destination + offset
	}

	return value
}

func (m *Almanac) GetMinimumLocation() int {

	min := math.MaxInt

	soils := []int{}

	for _, seed := range m.seeds {
		value := getValueFromMaps(&m.seedToSoilMap, seed)
		soils = append(soils, value)
	}

	fertilizers := []int{}
	for _, soil := range soils {
		value := getValueFromMaps(&m.soilToFertilizerMap, soil)
		fertilizers = append(fertilizers, value)
	}

	waters := []int{}
	for _, fertilizer := range fertilizers {
		value := getValueFromMaps(&m.fertilizerToWaterMap, fertilizer)
		waters = append(waters, value)
	}

	lights := []int{}
	for _, water := range waters {
		value := getValueFromMaps(&m.waterToLightMap, water)
		lights = append(lights, value)
	}

	temperatures := []int{}
	for _, light := range lights {
		value := getValueFromMaps(&m.lightToTemperature, light)
		temperatures = append(temperatures, value)
	}

	humidities := []int{}
	for _, temperature := range temperatures {
		value := getValueFromMaps(&m.temperatureToHumidity, temperature)
		humidities = append(humidities, value)
	}

	locations := []int{}
	for _, humidity := range humidities {
		value := getValueFromMaps(&m.humidityToLocation, humidity)
		locations = append(locations, value)
	}

	for _, location := range locations {
		if location < min {
			min = location
		}
	}

	return min
}
