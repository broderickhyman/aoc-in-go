package main

import (
	"fmt"
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

	result := 0
	for _, line := range strings.Split(input, "\n"){

		length := len(line)
		currentString := ""
		for i:= 0; i < length; i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				currentString += string(char)
				break
			}
		}
		for i:= length - 1; i >= 0 ; i-- {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				currentString += string(char)
				break
			}
		}
		num, err := strconv.Atoi(currentString)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		result += num
	}
	// solve part 1 here
	return result
}
