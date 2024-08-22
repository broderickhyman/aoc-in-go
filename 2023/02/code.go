package main

import (
	"strconv"
	"strings"

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
func run(isPart2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if isPart2 {
		return part2(input)
	}
	return part1(input)
}

func part1(input string) int {
	gameIds := 0
	data := strings.Split(input, "\n")
	for _, game := range data {
		gameData := strings.Split(game, ":")
		allGrabData := gameData[1]
		gameId, _ := strconv.Atoi(strings.Replace(gameData[0], "Game ", "", 1))
		maxMap := createMap(allGrabData)

		if maxMap["red"] <= 12 && maxMap["green"] <= 13 && maxMap["blue"] <= 14 {
			gameIds += gameId
		}
	}

	return gameIds
}

func part2(input string) int {
	totalPower := 0
	data := strings.Split(input, "\n")
	for _, game := range data {
		gameData := strings.Split(game, ":")
		allGrabData := gameData[1]
		maxMap := createMap(allGrabData)

		cubePower := maxMap["red"] * maxMap["green"] * maxMap["blue"]
		totalPower += cubePower
	}

	return totalPower
}

func createMap(allGrabData string) map[string]int {
	allGrabs := strings.Split(allGrabData, ";")

	maxMap := make(map[string]int)
	for _, pickData := range allGrabs {
		colorCombos := strings.Split(pickData, ",")
		trimmedCombos := make([]string, len(colorCombos))
		for i, colorCombo := range colorCombos {
			trimmedCombos[i] = strings.TrimSpace(colorCombo)
		}
		for _, dieCombo := range trimmedCombos {
			dieData := strings.Split(dieCombo, " ")
			num, _ := strconv.Atoi(dieData[0])
			color := dieData[1]
			currentMax := maxMap[color]
			if currentMax < num {
				maxMap[color] = num
			}
		}
	}
	return maxMap
}
