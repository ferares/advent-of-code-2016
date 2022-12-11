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
	template, err := os.Open("./template.go.txt")
	if (err != nil) { panic(err) }
	defer template.Close()
	out, err := os.Create(fileName)
	if (err != nil) { panic(err) }
	defer out.Close()
	content, err := io.ReadAll(template)
	if (err != nil) { panic(err) }
	newContent := strings.ReplaceAll(string(content), "Day#", "Day" + day)
	_, err = out.WriteString(newContent)
	if (err != nil) { panic(err) }
	err = out.Sync()
	if (err != nil) { panic(err) }
	main, err := os.OpenFile("./main.go", os.O_RDWR, os.ModeAppend)
	if (err != nil) { panic(err) }
	defer main.Close()
	content, err = io.ReadAll(main)
	if (err != nil) { panic(err) }
	err = main.Truncate(0)
	if (err != nil) { panic(err) }
	_, err = main.Seek(0, 0)
	if (err != nil) { panic(err) }
	comment := "// case \"#\": daySolution = days.Day#{}"
	newLine := strings.ReplaceAll(comment, "// ", "")
	newLine = strings.ReplaceAll(newLine, "#", day) + "\n\t\t" + comment
	newContent = strings.ReplaceAll(string(content), comment, newLine)
	_, err = main.WriteString(newContent)
	if (err != nil) { panic(err) }
	_, err = os.Create("./input/" + day + ".txt")
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

func getDaySolution(day string) solutions.DaySolution {
	var daySolution solutions.DaySolution
	switch day {
		case "1": daySolution = days.Day1{}
		case "2": daySolution = days.Day2{}
		case "3": daySolution = days.Day3{}
		case "4": daySolution = days.Day4{}
		// case "#": daySolution = days.Day#{}
	}
	return daySolution
}

func main() {
	args := os.Args[1:]
	if (len(args) == 2) {
		day := args[0]
		puzzle := args[1]
		input := readInput(day)
		daySolution := getDaySolution(day)
		runSolution(daySolution, puzzle, input)
	} else if (len(args) == 1) {
		day := args[0]
		createDayFromTemplate(day)
	}
}