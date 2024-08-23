package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return runPart1(input)
}

func runPart1(input string) int {
	offsets := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	data := strings.Split(input, "\n")
	rowCount := len(data)
	columnCount := len(data[0])
	schematic := make([][]rune, columnCount)
	for i := 0; i < columnCount; i++ {
		schematic[i] = make([]rune, rowCount)
	}
	for rowIndex, row := range data {
		for columnIndex, column := range row {
			schematic[columnIndex][rowIndex] = column
		}
	}

	sumOfPartNumbers := 0
	for r := 0; r < rowCount; r++ {
		nextToSymbol := false
		numberString := ""
		numberEnd := false
		for c := 0; c < columnCount; c++ {
			currentChar := schematic[c][r]
			isDigit := unicode.IsDigit(currentChar)
			if isDigit {
				numberString += string(currentChar)
				if !nextToSymbol {
					for _, offset := range offsets {
						col := c + offset[0]
						row := r + offset[1]
						if col == -1 || row == -1 || col >= columnCount || row >= rowCount {
							continue
						}
						check := schematic[col][row]
						if !unicode.IsDigit(check) && !unicode.IsLetter(check) && check != '.' {
							nextToSymbol = true
							break
						}
					}
				}
			} else {
				numberEnd = true
			}
			// Handle the end of the line
			if isDigit && c == columnCount-1 {
				numberEnd = true
			}
			if numberEnd {
				if nextToSymbol {
					partNumber, _ := strconv.Atoi(numberString)
					sumOfPartNumbers += partNumber
					nextToSymbol = false
				}
				numberString = ""
				numberEnd = false
			}
		}
	}

	return sumOfPartNumbers
}
