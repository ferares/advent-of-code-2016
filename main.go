package main

import (
	"aoc2016/solutions"
	"aoc2016/solutions/days"
	"fmt"
	"io"
	"os"
	"strings"
)

func readInput(day string) string {
	file, err := os.Open("./input/" + day + ".txt")
	if (err != nil) { panic(err)}
	defer file.Close()
	data, err := io.ReadAll(file)
	if (err != nil) { panic(err)}
	return string(data)
}

func createDayFromTemplate(day string) {
	fileName := "./solutions/days/day" + day + ".go"
	_, err := os.Stat(fileName)
	if (err == nil) { panic("File '" + fileName + "' already exists!") }
	in, err := os.Open("./template.go.txt")
	if (err != nil) { panic(err) }
	defer in.Close()
	out, err := os.Create(fileName)
	if (err != nil) { panic(err) }
	defer out.Close()
	content, err := io.ReadAll(in)
	if (err != nil) { panic(err) }
	newContent := strings.ReplaceAll(string(content), "Day#", "Day" + day)
	_, err = out.WriteString(newContent)
	if (err != nil) { panic(err) }
	err = out.Sync()
	if (err != nil) { panic(err) }
}

func runSolution(daySolution solutions.DaySolution, puzzle string, input string) {
	solution := ""
	if (puzzle == "1") {
		solution = daySolution.Puzzle1(input)
	} else {
		solution = daySolution.Puzzle2(input)
	}
	fmt.Println((solution))
}

func main() {
	args := os.Args[1:]
	if (len(args) == 2) {
		day := args[0]
		puzzle := args[1]
		input := readInput(day)
		var daySolution solutions.DaySolution
		switch day {
		case "1":
			daySolution = days.Day1{}
		}
		runSolution(daySolution, puzzle, input)
	} else if (len(args) == 1) {
		day := args[0]
		createDayFromTemplate(day)
	}
}