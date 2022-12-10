package days

import "strings"

type Day2 struct {}

func decodeDigit(coordinates [2]int) string {
	digit := ""
	if (coordinates[1] == 0) {
		switch coordinates[0] {
			case 0: digit = "1"
			case 1: digit = "2"
			case 2: digit = "3"
		}
	} else if (coordinates[1] == 1) {
		switch coordinates[0] {
			case 0: digit = "4"
			case 1: digit = "5"
			case 2: digit = "6"
		}
	} else if (coordinates[1] == 2) {
		switch coordinates[0] {
			case 0: digit = "7"
			case 1: digit = "8"
			case 2: digit = "9"
		}
	}
	return digit
}

func decodeDigit2(coordinates [2]int) string {
	digit := ""
	if (coordinates[1] == 0) {
		switch coordinates[0] {
			case 2: digit = "1"
		}
	} else if (coordinates[1] == 1) {
		switch coordinates[0] {
			case 1: digit = "2"
			case 2: digit = "3"
			case 3: digit = "4"
		}
	} else if (coordinates[1] == 2) {
		switch coordinates[0] {
			case 0: digit = "5"
			case 1: digit = "6"
			case 2: digit = "7"
			case 3: digit = "8"
			case 4: digit = "9"
		}
	} else if (coordinates[1] == 3) {
		switch coordinates[0] {
			case 1: digit = "A"
			case 2: digit = "B"
			case 3: digit = "C"
		}
	} else if (coordinates[1] == 4) {
		switch coordinates[0] {
			case 2: digit = "D"
		}
	}
	return digit
}

func (Day2) Puzzle1(input string) string {
	lines := strings.Split(input, "\n")
	coordinates := [2]int{1,1}
	digits := ""
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			instruction := line[j]
			newCoordinates := coordinates
			switch instruction {
				case 'U': newCoordinates[1] -= 1
				case 'D': newCoordinates[1] += 1
				case 'L': newCoordinates[0] -= 1
				case 'R': newCoordinates[0] += 1
			}
			digit := decodeDigit(newCoordinates)
			if (digit != "") {
				coordinates = newCoordinates
			}
		}
		digits += decodeDigit(coordinates)
	}
	return digits
}

func (Day2) Puzzle2(input string) string {
	lines := strings.Split(input, "\n")
	coordinates := [2]int{0,2}
	digits := ""
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			instruction := line[j]
			newCoordinates := coordinates
			switch instruction {
				case 'U': newCoordinates[1] -= 1
				case 'D': newCoordinates[1] += 1
				case 'L': newCoordinates[0] -= 1
				case 'R': newCoordinates[0] += 1
			}
			digit := decodeDigit2(newCoordinates)
			if (digit != "") {
				coordinates = newCoordinates
			}
		}
		digits += decodeDigit2(coordinates)
	}
	return digits
}
