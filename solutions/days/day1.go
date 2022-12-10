package days

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Day1 struct {}

func parseInput(input string) []string {
	return strings.Split(input, ", ")
}

func parseInstruction(instruction string) (byte, int) {
	direction := instruction[0]
	steps, err := strconv.Atoi(instruction[1:])
	if (err != nil) { panic(err) }
	return direction, steps
}

func changeFacing(facing int, direction byte) int {
	if (direction == 'L') { facing = (facing - 1) % 4 }
	if (direction == 'R') { facing = (facing + 1) % 4 }
	if (facing < 0) { facing = 4 + facing }
	return facing
}

func walk(direction byte, steps int, facing int, coordinates [2]int) ([2]int, int) {
	if (direction == 'L') { facing = (facing - 1) % 4 }
	if (direction == 'R') { facing = (facing + 1) % 4 }
	if (facing < 0) { facing = 4 + facing }
	switch facing {
		case 0: coordinates[1] -= steps
		case 1: coordinates[0] += steps
		case 2: coordinates[1] += steps
		case 3: coordinates[0] -= steps
	}
	return coordinates, facing
}

func walkStep(direction byte, facing int, coordinates [2]int) [2]int {
	switch facing {
		case 0: coordinates[1] -= 1
		case 1: coordinates[0] += 1
		case 2: coordinates[1] += 1
		case 3: coordinates[0] -= 1
	}
	return coordinates
}

func (Day1) Puzzle1(input string) string {
	instructions := parseInput(input)
	coordinates := [2]int{0,0}
	facing := 0
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		direction, steps := parseInstruction(instruction)
		coordinates, facing = walk(direction, steps, facing, coordinates)
	}
	if (coordinates[0] < 0) { coordinates[0] = -coordinates[0] }
	if (coordinates[1] < 0) { coordinates[1] = -coordinates[1] }
	return strconv.Itoa(coordinates[0] + coordinates[1])
}

func (Day1) Puzzle2(input string) string {
	instructions := parseInput(input)
	visitedCoordinates := make([]string, 0)
	visitedCoordinates = append(visitedCoordinates, "0,0")
	coordinates := [2]int{0,0}
	facing := 0
	out:
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		direction, steps := parseInstruction(instruction)
		facing = changeFacing(facing, direction)
		for i := 0; i < steps; i++ {
			coordinates = walkStep(direction, facing, coordinates)
			coordinatesKey := strings.Join([]string{strconv.Itoa(coordinates[0]), strconv.Itoa(coordinates[1])}, ",")
			if (slices.Contains(visitedCoordinates, coordinatesKey)) { break out }
			visitedCoordinates = append(visitedCoordinates, coordinatesKey)
		}
	}
	if (coordinates[0] < 0) { coordinates[0] = -coordinates[0] }
	if (coordinates[1] < 0) { coordinates[1] = -coordinates[1] }
	return strconv.Itoa(coordinates[0] + coordinates[1])
}
