package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day6Sample.txt"
	} else {
		fileName = "../inputfiles/Day6.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	var areaMap []string
	visitedPositions := make(map[[2]int]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		areaMap = append(areaMap, scanner.Text())
	}
	obstaclePositions := make(map[[2]int]struct{})
	var guardStartingPosition [2]int
	var guardStartingDirection byte
	var guardPosition [2]int
	var guardDirection byte
	for i, row := range areaMap {
		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				obstaclePositions[[2]int{i, j}] = struct{}{}
			} else if row[j] != '.' {
				guardStartingPosition = [2]int{i, j}
				guardPosition = [2]int{i, j}
				guardStartingDirection = row[j]
				guardDirection = row[j]
			}
		}
	}
	for guardPosition != [2]int{-1, -1} {
		visitedPositions[guardPosition] = struct{}{}
		guardPosition, guardDirection = moveGuard(guardPosition, guardDirection, obstaclePositions, areaMap)
	}
	fmt.Printf("Part 1 Answer: %d\n", len(visitedPositions))
	var newObstaclePositions int
	for visitedPosition := range visitedPositions {
		guardPosition = [2]int{guardStartingPosition[0], guardStartingPosition[1]}
		guardDirection = guardStartingDirection
		if visitedPosition != guardStartingPosition {
			previousTurnDirection := guardDirection
			previousTurns := make(map[[2]int]struct{})
			obstaclePositions[visitedPosition] = struct{}{}
			var spacesMoved int
			for guardPosition != [2]int{-1, -1} {
				guardPosition, guardDirection = moveGuard(guardPosition, guardDirection, obstaclePositions, areaMap)
				if guardDirection != previousTurnDirection {
					if _, ok := previousTurns[guardPosition]; !ok {
						previousTurns[guardPosition] = struct{}{}
						previousTurnDirection = guardDirection
						spacesMoved = 0
					} else if spacesMoved > 0 {
						newObstaclePositions++
						guardPosition = [2]int{-1, -1}
					} else {
						previousTurnDirection = guardDirection
						spacesMoved++
					}
				}
			}
			delete(obstaclePositions, visitedPosition)
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", newObstaclePositions)
}

func moveGuard(guardPosition [2]int, guardDirection byte, obstaclePositions map[[2]int]struct{}, areaMap []string) ([2]int, byte) {
	maxRow := len(areaMap)
	maxCol := len(areaMap[0])
	if guardDirection == '^' {
		if _, ok := obstaclePositions[[2]int{guardPosition[0] - 1, guardPosition[1]}]; ok {
			return guardPosition, '>'
		} else if guardPosition[0]-1 < 0 {
			return [2]int{-1, -1}, guardDirection
		} else {
			return [2]int{guardPosition[0] - 1, guardPosition[1]}, guardDirection
		}
	} else if guardDirection == '<' {
		if _, ok := obstaclePositions[[2]int{guardPosition[0], guardPosition[1] - 1}]; ok {
			return guardPosition, '^'
		} else if guardPosition[1]-1 < 0 {
			return [2]int{-1, -1}, guardDirection
		} else {
			return [2]int{guardPosition[0], guardPosition[1] - 1}, guardDirection
		}
	} else if guardDirection == '>' {
		if _, ok := obstaclePositions[[2]int{guardPosition[0], guardPosition[1] + 1}]; ok {
			return guardPosition, 'v'
		} else if guardPosition[1]+1 >= maxCol {
			return [2]int{-1, -1}, guardDirection
		} else {
			return [2]int{guardPosition[0], guardPosition[1] + 1}, guardDirection
		}
	} else {
		if _, ok := obstaclePositions[[2]int{guardPosition[0] + 1, guardPosition[1]}]; ok {
			return guardPosition, '<'
		} else if guardPosition[0]+1 >= maxRow {
			return [2]int{-1, -1}, guardDirection
		} else {
			return [2]int{guardPosition[0] + 1, guardPosition[1]}, guardDirection
		}
	}
}
