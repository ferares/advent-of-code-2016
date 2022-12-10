package days

import (
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct {}

func checkColumn(column []int) int {
	possible := 0
	for i := 0; i < len(column); i = i + 3 {
		side1 := column[i]
		side2 := column[i + 1]
		side3 := column[i + 2]
		if (side1 + side2 > side3) && (side1 + side3 > side2) && (side2 + side3 > side1) {
			possible++
		}
	}
	return possible
}

func (Day3) Puzzle1(input string) string {
	triangles := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s*(?P<side1>\d+)\s*(?P<side2>\d+)\s*(?P<side3>\d+)`)
	possible := 0
	for _, triangle := range triangles {
		sides := re.FindStringSubmatch(triangle)
		side1, err := strconv.Atoi(strings.TrimSpace(sides[1]))
		if (err != nil) { panic(err) }
		side2, err := strconv.Atoi(strings.TrimSpace(sides[2]))
		if (err != nil) { panic(err) }
		side3, err := strconv.Atoi(strings.TrimSpace(sides[3]))
		if (err != nil) { panic(err) }
		if (side1 + side2 > side3) && (side1 + side3 > side2) && (side2 + side3 > side1) {
			possible++
		}
	}
	return strconv.Itoa(possible)
}

func (Day3) Puzzle2(input string) string {
	triangles := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s*(?P<side1>\d+)\s*(?P<side2>\d+)\s*(?P<side3>\d+)`)
	column1 := make([]int, 0)
	column2 := make([]int, 0)
	column3 := make([]int, 0)
	possible := 0
	for _, triangle := range triangles {
		sides := re.FindStringSubmatch(triangle)
		side1, err := strconv.Atoi(strings.TrimSpace(sides[1]))
		if (err != nil) { panic(err) }
		side2, err := strconv.Atoi(strings.TrimSpace(sides[2]))
		if (err != nil) { panic(err) }
		side3, err := strconv.Atoi(strings.TrimSpace(sides[3]))
		if (err != nil) { panic(err) }
		column1 = append(column1, side1)
		column2 = append(column2, side2)
		column3 = append(column3, side3)
	}
	possible += checkColumn(column1)
	possible += checkColumn(column2)
	possible += checkColumn(column3)
	return strconv.Itoa(possible)
}
