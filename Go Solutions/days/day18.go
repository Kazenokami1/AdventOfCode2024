package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
	Part 2 does take a bit as it is pretty much brute force. For my solution it takes ~36 seconds. A better way would be to see
	if there is no break between #s from one side to another. Check each # for it's neighbors (including diagonals), and if it's another
	#, keep checking. If you go from one side to another and 0,0 is contained on one side, and size,size is contained on the other, it is blocked.
*/

func Day18() {
	var fileName string
	var tetrisSize int
	var dropsToPlace int
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day18Sample.txt"
		tetrisSize = 6
		dropsToPlace = 12
	} else {
		fileName = "../inputfiles/Day18.txt"
		tetrisSize = 70
		dropsToPlace = 1024
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	tetrisMaze := make(map[[2]int]*mazeSquare)
	var droppingBlocks []*mazeSquare
	for x := 0; x <= tetrisSize; x++ {
		for y := 0; y <= tetrisSize; y++ {
			tetrisMaze[[2]int{x, y}] = &mazeSquare{position: [2]int{x, y}, score: math.Inf(1), squareType: "."}
		}
	}
	tetrisMaze[[2]int{0, 0}].score = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		xCoord, _ := strconv.Atoi(coords[0])
		yCoord, _ := strconv.Atoi(coords[1])
		droppingBlocks = append(droppingBlocks, tetrisMaze[[2]int{xCoord, yCoord}])
	}
	for i := 0; i < dropsToPlace; i++ {
		tetrisMaze[droppingBlocks[i].position].squareType = "#"
	}
	for _, square := range tetrisMaze {
		if exit, ok := tetrisMaze[[2]int{square.position[0] + 1, square.position[1]}]; ok {
			square.exits = append(square.exits, exit)
		}
		if exit, ok := tetrisMaze[[2]int{square.position[0] - 1, square.position[1]}]; ok {
			square.exits = append(square.exits, exit)
		}
		if exit, ok := tetrisMaze[[2]int{square.position[0], square.position[1] + 1}]; ok {
			square.exits = append(square.exits, exit)
		}
		if exit, ok := tetrisMaze[[2]int{square.position[0], square.position[1] - 1}]; ok {
			square.exits = append(square.exits, exit)
		}
	}
	leastPath := moveThroughMaze(tetrisMaze, tetrisMaze[[2]int{0, 0}], tetrisSize, math.Inf(1))
	fmt.Printf("Part 1 Answer: %d\n", int(leastPath))
	for i := dropsToPlace; i < len(droppingBlocks); i++ {
		tetrisMaze[droppingBlocks[i].position].squareType = "#"
		pathBlocked, _ := checkBlockage(tetrisMaze, tetrisMaze[[2]int{0, 0}], tetrisSize, []*mazeSquare{})
		if pathBlocked {
			fmt.Printf("Part 2 Answer: %d,%d\n", droppingBlocks[i].position[0], droppingBlocks[i].position[1])
			break
		}
	}
}

func moveThroughMaze(tetrisMaze map[[2]int]*mazeSquare, currentPosition *mazeSquare, size int, lowestScore float64) float64 {
	if currentPosition == tetrisMaze[[2]int{size, size}] && currentPosition.score < lowestScore {
		return currentPosition.score
	}
	for _, exit := range currentPosition.exits {
		if currentPosition.score+1 < exit.score && exit.squareType == "." {
			exit.score = currentPosition.score + 1
			lowestScore = moveThroughMaze(tetrisMaze, exit, size, lowestScore)
		}
	}
	return lowestScore
}

func checkBlockage(tetrisMaze map[[2]int]*mazeSquare, currentPosition *mazeSquare, size int, visited []*mazeSquare) (bool, []*mazeSquare) {
	var pathBlocked bool
	visited = append(visited, currentPosition)
	for _, exit := range currentPosition.exits {
		if exit.squareType == "." && !slices.Contains(visited, exit) {
			_, visited = checkBlockage(tetrisMaze, exit, size, visited)
		}
	}
	pathBlocked = !slices.Contains(visited, tetrisMaze[[2]int{size, size}])
	return pathBlocked, visited
}
