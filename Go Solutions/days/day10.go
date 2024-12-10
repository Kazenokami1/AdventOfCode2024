package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day10Sample.txt"
	} else {
		fileName = "../inputfiles/Day10.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	heightPositions := make(map[[2]int]int64)
	var row int
	for scanner.Scan() {
		heightStrings := strings.Split(scanner.Text(), "")
		for col, height := range heightStrings {
			heightInt, _ := strconv.ParseInt(string(height), 0, 0)
			heightPositions[[2]int{row, col}] = heightInt
		}
		row++
	}
	var trailHeadSum int
	trailHeadRatings := make(map[[2]int]int)
	for position, val := range heightPositions {
		if val == 0 {
			trailHeads := returnTrailHeads(heightPositions, position, make(map[[2]int]int))
			trailHeadSum += len(trailHeads)
			for pos, val := range trailHeads {
				trailHeadRatings[pos] += val
			}
		}
	}
	var trailHeadRatingSum int
	for _, val := range trailHeadRatings {
		trailHeadRatingSum += val
	}
	fmt.Printf("Part 1 Answer: %d\n", trailHeadSum)
	fmt.Printf("Part 2 Answer: %d\n", trailHeadRatingSum)
}

func returnTrailHeads(positions map[[2]int]int64, currentPosition [2]int, numberOfTrailHeads map[[2]int]int) map[[2]int]int {
	if positions[currentPosition] == 9 {
		numberOfTrailHeads[currentPosition]++
		return numberOfTrailHeads
	}
	if val, ok := positions[[2]int{currentPosition[0] - 1, currentPosition[1]}]; ok {
		if val == positions[currentPosition]+1 {
			numberOfTrailHeads = returnTrailHeads(positions, [2]int{currentPosition[0] - 1, currentPosition[1]}, numberOfTrailHeads)
		}
	}
	if val, ok := positions[[2]int{currentPosition[0] + 1, currentPosition[1]}]; ok {
		if val == positions[currentPosition]+1 {
			numberOfTrailHeads = returnTrailHeads(positions, [2]int{currentPosition[0] + 1, currentPosition[1]}, numberOfTrailHeads)
		}
	}
	if val, ok := positions[[2]int{currentPosition[0], currentPosition[1] - 1}]; ok {
		if val == positions[currentPosition]+1 {
			numberOfTrailHeads = returnTrailHeads(positions, [2]int{currentPosition[0], currentPosition[1] - 1}, numberOfTrailHeads)
		}
	}
	if val, ok := positions[[2]int{currentPosition[0], currentPosition[1] + 1}]; ok {
		if val == positions[currentPosition]+1 {
			numberOfTrailHeads = returnTrailHeads(positions, [2]int{currentPosition[0], currentPosition[1] + 1}, numberOfTrailHeads)
		}
	}
	return numberOfTrailHeads
}
