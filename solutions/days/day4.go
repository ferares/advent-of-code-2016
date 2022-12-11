package days

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day4 struct {}

func sortMap(letters map[rune]int) []rune {
	keys := make([]rune, 0)
	for key := range letters { keys = append(keys, key) }
	sort.SliceStable(keys, func(i, j int) bool {
		keyA := keys[i]
		keyB := keys[j]
		letterA := letters[keyA]
		letterB := letters[keyB]
		if (letterA == letterB) { return keyA < keyB }
		return letterA > letterB
	})
	return keys
}

func getRealRooms(input string) (int, map[int]string) {
	realRoomSum := 0
	realRooms := make(map[int]string, 0)
	rooms := strings.Split(input, "\n")
	re := regexp.MustCompile(`(.*)-(\d+)\[(.*)\]`)
	for _, room := range rooms {
		data := re.FindStringSubmatch(room)
		letters := make(map[rune]int, 0)
		name := strings.ReplaceAll(data[1], "-", " ")
		for _, letter := range name {
			if (letter == ' ') { continue }
			letters[letter] += 1
		}
		id, _ := strconv.Atoi(data[2])
		chk := data[3]
		keys := sortMap(letters)
		real := true
		for i, letter := range chk {
			if (letter != keys[i]) { 
				real = false
				break
			}
		}
		if (real) {
			realRoomSum += id
			realRooms[id] = name
		}
	}
	return realRoomSum, realRooms
}

func (Day4) Puzzle1(input string) string {
	realRoomSum, _ := getRealRooms(input)
	return strconv.Itoa(realRoomSum)
}

func (Day4) Puzzle2(input string) string {
	_, rooms := getRealRooms(input)
	target := 0
	for id, room := range rooms {
		realName := ""
		for _, letter := range room {
			if (letter == ' ') {
				realName += " "
				continue
			}
			decodedLetter := string((((int(letter) + id) - 97) % 26) + 97)
			realName += decodedLetter
		}
		if (realName == "northpole object storage") {
			target = id
		}
	}
	return strconv.Itoa(target)
}
