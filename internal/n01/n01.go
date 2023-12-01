package n01

import (
	"fmt"
	"strconv"
	"strings"
)

func extractDigits(val string) (digits []int) {
	for _, char := range strings.Split(val, "") {
		if err, i := isInt(char); err {
			digits = append(digits, i)
		}
	}
	return
}

func writtenNumbertoNumeral(number string) string {
	translation := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	return translation[number]
}

func isInt(char string) (bool, int) {
	i, err := strconv.Atoi(char)
	return err == nil, i
}

func createCalibrationValue(data []int) (value int) {
	return (data[0] * 10) + data[len(data)-1]
}

func TranslateValue(val *string) {
	writtenNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, number := range writtenNumbers {
		if strings.Contains(*val, number) {
			*val = strings.ReplaceAll(*val, number, fmt.Sprintf("%s%s%s", number, writtenNumbertoNumeral(number), number))
		}
	}
	for _, number := range writtenNumbers {
		if strings.Contains(*val, number) {
			*val = strings.ReplaceAll(*val, number, "")
		}
	}
}

func GetTotalCalibrationValue(data []string) int {
	var cleanedData [][]int

	for _, value := range data {
		extracted := extractDigits(value)
		if len(extracted) > 0 {
			cleanedData = append(cleanedData, extracted)
		}
	}

	addedCalibrationValues := 0
	for _, digits := range cleanedData {
		addedCalibrationValues += createCalibrationValue(digits)
	}
	return addedCalibrationValues
}
