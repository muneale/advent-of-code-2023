package n18

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point struct {
	x, y       int
	dir, color string
}

func Run() {
	wd, _ := os.Getwd()
	filepath := filepath.Join(wd, "input/n18.txt")

	input := getLines(filepath)

	points := getPoints(&input, false)

	area := getArea(points)
	perimeter := getPerimeter(points)
	internalPoints := getInternalPoints(perimeter, area)

	fmt.Printf("Part 1: %v\n", perimeter+internalPoints)

	pointsByColor := getPoints(&input, true)
	area = getArea(pointsByColor)
	perimeter = getPerimeter(pointsByColor)
	internalPoints = getInternalPoints(perimeter, area)

	fmt.Printf("Part 2: %v\n", perimeter+internalPoints)

}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func getPoints(lines *[]string, byColor bool) *[]Point {

	points := make([]Point, len(*lines)+1)

	i, j := 0, 0
	points[0] = Point{i, j, "", ""}
	for k, line := range *lines {
		data := strings.Split(line, " ")
		n := 0
		dir := data[0]
		if byColor {
			hexNum := data[2]
			dir = hexNum[len(hexNum)-2 : len(hexNum)-1]
			hexNum = hexNum[2 : len(hexNum)-2]
			parsedN, _ := strconv.ParseInt(hexNum, 16, 32)
			n = int(parsedN)
		} else {
			n, _ = strconv.Atoi(data[1])
		}

		if dir == "R" || dir == "0" {
			j += n
		} else if dir == "L" || dir == "2" {
			j -= n
		} else if dir == "U" || dir == "3" {
			i -= n
		} else if dir == "D" || dir == "1" {
			i += n
		}

		points[k+1] = Point{i, j, "", line}
	}

	return &points
}

func getArea(points *[]Point) int {
	sum := 0

	for i := 0; i < len(*points)-1; i++ {

		p1 := (*points)[i]
		p2 := (*points)[i+1]

		sum += p1.x*p2.y - p1.y*p2.x
	}

	area := math.Abs(float64(sum)) / 2

	return int(area)
}

func getPerimeter(points *[]Point) int {
	p := 0

	for i := 0; i < len(*points)-1; i++ {

		p1 := (*points)[i]
		p2 := (*points)[i+1]

		p += int(math.Abs(float64(p2.x-p1.x))) +
			int(math.Abs(float64(p2.y-p1.y)))
	}

	return p
}

func getInternalPoints(perimeter int, area int) int {
	return area - perimeter/2 + 1
}
