package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day16() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day16Sample.txt"
	} else {
		fileName = "../inputfiles/Day16.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var maze []string
	var rudolph *reindeer
	mazeSquares := make(map[[2]int]*mazeSquare)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	for i := 0; i < len(maze); i++ {
		for j, square := range maze[i] {
			if square == 'S' {
				rudolph = &reindeer{position: [2]int{i, j}, facing: "EAST"}
			}
			if square != '#' {
				mazeSquares[[2]int{i, j}] = &mazeSquare{position: [2]int{i, j}, squareType: string(square), score: math.Inf(1)}
			}
		}
	}
	for _, square := range mazeSquares {
		if square.squareType == "S" {
			square.score = 0
		}
		if nextSquare, ok := mazeSquares[[2]int{square.position[0] + 1, square.position[1]}]; ok {
			square.exits = append(square.exits, nextSquare)
		}
		if nextSquare, ok := mazeSquares[[2]int{square.position[0] - 1, square.position[1]}]; ok {
			square.exits = append(square.exits, nextSquare)
		}
		if nextSquare, ok := mazeSquares[[2]int{square.position[0], square.position[1] + 1}]; ok {
			square.exits = append(square.exits, nextSquare)
		}
		if nextSquare, ok := mazeSquares[[2]int{square.position[0], square.position[1] - 1}]; ok {
			square.exits = append(square.exits, nextSquare)
		}
	}
	leastScore := moveRudolphThroughMaze(mazeSquares, rudolph, rudolph.position, "EAST", math.Inf(1))
	fmt.Printf("Part 1 Answer: %d\n", int(leastScore))
}

func moveRudolphThroughMaze(mazeSquares map[[2]int]*mazeSquare, rudolph *reindeer, lastPosition [2]int, lastFacing string, leastScore float64) float64 {
	currentSquare := mazeSquares[rudolph.position]
	rudolph.currentPath = append(rudolph.currentPath, currentSquare.position)
	if currentSquare.squareType == "E" && currentSquare.score < leastScore {
		return currentSquare.score
	}
	for _, exit := range currentSquare.exits {
		switch rudolph.facing {
		case "EAST":
			if exit.position == [2]int{rudolph.position[0], rudolph.position[1] + 1} && currentSquare.score+1 < exit.score {
				exit.score = currentSquare.score + 1
				rudolph.position = exit.position
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "EAST", leastScore)
			} else if exit.position == [2]int{rudolph.position[0] - 1, rudolph.position[1]} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "NORTH"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "EAST", leastScore)
			} else if exit.position == [2]int{rudolph.position[0] + 1, rudolph.position[1]} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "SOUTH"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "EAST", leastScore)
			}
		case "WEST":
			if exit.position == [2]int{rudolph.position[0], rudolph.position[1] - 1} && currentSquare.score+1 < exit.score {
				exit.score = currentSquare.score + 1
				rudolph.position = exit.position
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "WEST", leastScore)
			} else if exit.position == [2]int{rudolph.position[0] - 1, rudolph.position[1]} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "NORTH"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "WEST", leastScore)
			} else if exit.position == [2]int{rudolph.position[0] + 1, rudolph.position[1]} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "SOUTH"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "WEST", leastScore)
			}
		case "NORTH":
			if exit.position == [2]int{rudolph.position[0] - 1, rudolph.position[1]} && currentSquare.score+1 < exit.score {
				exit.score = currentSquare.score + 1
				rudolph.position = exit.position
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "NORTH", leastScore)
			} else if exit.position == [2]int{rudolph.position[0], rudolph.position[1] - 1} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "WEST"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "NORTH", leastScore)
			} else if exit.position == [2]int{rudolph.position[0], rudolph.position[1] + 1} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "EAST"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "NORTH", leastScore)
			}
		case "SOUTH":
			if exit.position == [2]int{rudolph.position[0] + 1, rudolph.position[1]} && currentSquare.score+1 < exit.score {
				exit.score = currentSquare.score + 1
				rudolph.position = exit.position
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "SOUTH", leastScore)
			} else if exit.position == [2]int{rudolph.position[0], rudolph.position[1] - 1} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "WEST"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "SOUTH", leastScore)
			} else if exit.position == [2]int{rudolph.position[0], rudolph.position[1] + 1} && currentSquare.score+1001 < exit.score {
				exit.score = currentSquare.score + 1001
				rudolph.position = exit.position
				rudolph.facing = "EAST"
				leastScore = moveRudolphThroughMaze(mazeSquares, rudolph, currentSquare.position, "SOUTH", leastScore)
			}
		}
	}
	rudolph.position = lastPosition
	rudolph.facing = lastFacing
	return leastScore
}
