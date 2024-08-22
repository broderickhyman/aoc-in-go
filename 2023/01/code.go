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
	if part2 {
		return part2Func(input)
	}

	return part1Func(input)
}

func part1Func(input string) int {
	result := 0
	for _, line := range strings.Split(input, "\n") {

		length := len(line)
		currentString := ""
		for i := 0; i < length; i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				currentString += string(char)
				break
			}
		}
		for i := length - 1; i >= 0; i-- {
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
	return result
}

func part2Func(input string) int {
	wordToValue := make(map[string]int)
	wordToValue["one"] = 1
	wordToValue["two"] = 2
	wordToValue["three"] = 3
	wordToValue["four"] = 4
	wordToValue["five"] = 5
	wordToValue["six"] = 6
	wordToValue["seven"] = 7
	wordToValue["eight"] = 8
	wordToValue["nine"] = 9

	result := 0
	for _, line := range strings.Split(input, "\n") {

		currentString := ""
		firstDigit := findFirstDigitIndex(line)
		firstWord := findFirstWordIndex(wordToValue, line)
		if firstDigit.Index >= 0 && (firstDigit.Index < firstWord.Index || firstWord.Index == -1) {
			currentString += firstDigit.Value
		} else if firstWord.Index >= 0 && (firstWord.Index < firstDigit.Index || firstDigit.Index == -1) {
			currentString += firstWord.Value
		}
		lastDigit := findLastDigitIndex(line)
		lastWord := findLastWordIndex(wordToValue, line)
		if lastDigit.Index < 9999 && (lastDigit.Index > lastWord.Index || lastWord.Index == 9999) {
			currentString += lastDigit.Value
		} else if lastDigit.Index < 9999 && (lastWord.Index > lastDigit.Index || lastDigit.Index == 9999) {
			currentString += lastWord.Value
		}

		num, _ := strconv.Atoi(currentString)
		result += num
	}
	return result
}

func findFirstDigitIndex(input string) IndexString {
	length := len(input)
	for i := 0; i < length; i++ {
		char := rune(input[i])
		if unicode.IsDigit(char) {
			return IndexString{i, string(char)}
		}
	}
	return IndexString{-1, ""}
}

func findLastDigitIndex(input string) IndexString {
	length := len(input)
	for i := length - 1; i >= 0; i-- {
		char := rune(input[i])
		if unicode.IsDigit(char) {
			return IndexString{i, string(char)}
		}
	}
	return IndexString{-1, ""}
}

func findFirstWordIndex(wordToValue map[string]int, input string) IndexString {
	minValue := IndexString{9999, ""}
	for k := range wordToValue {
		index := strings.Index(input, k)
		if index >= 0 {
			s := strconv.Itoa(wordToValue[k])
			if index < minValue.Index {
				minValue = IndexString{index, s}
			}
		}
	}

	return minValue
}

func findLastWordIndex(wordToValue map[string]int, input string) IndexString {
	maxValue := IndexString{-1, ""}
	for k := range wordToValue {
		index := strings.LastIndex(input, k)
		if index >= 0 {
			s := strconv.Itoa(wordToValue[k])
			if index > maxValue.Index {
				maxValue = IndexString{index, s}
			}
		}
	}

	return maxValue
}

type IndexString struct {
	Index int
	Value string
}
