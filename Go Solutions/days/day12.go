package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
	As a note, Part 2 of this will return the wrong cost if a group fully contained more than one group within the same contained grouping of plots.
	The fix for this would be to basically flip the robot so it runs a path around the inside of any hole that is found the same way it runs the
	path around the outside right now.

	In other words, pick a square and start inside the area, then check on your left for the wall instead of the right, and turn left when you move into
	a space where the wall is no longer on your left, and turn right if you would run into a wall. You also need to take into account that you need to check
	left first as you could have a 1 space opening on your left and have a wall directly in front of you (similar to how it is done that I check after I move).
*/

func Day12() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day12Sample.txt"
	} else {
		fileName = "../inputfiles/Day12.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	farmMap := make(map[[2]int]*farmPlot)
	var row int
	for scanner.Scan() {
		farmRow := strings.Split(scanner.Text(), "")
		for col, letter := range farmRow {
			farmMap[[2]int{row, col}] = &farmPlot{position: [2]int{row, col}, letter: letter}
		}
		row++
	}
	for _, plot := range farmMap {
		if neighbor, ok := farmMap[[2]int{plot.position[0] - 1, plot.position[1]}]; ok {
			plot.neighbors = append(plot.neighbors, neighbor)
		}
		if neighbor, ok := farmMap[[2]int{plot.position[0] + 1, plot.position[1]}]; ok {
			plot.neighbors = append(plot.neighbors, neighbor)
		}
		if neighbor, ok := farmMap[[2]int{plot.position[0], plot.position[1] - 1}]; ok {
			plot.neighbors = append(plot.neighbors, neighbor)
		}
		if neighbor, ok := farmMap[[2]int{plot.position[0], plot.position[1] + 1}]; ok {
			plot.neighbors = append(plot.neighbors, neighbor)
		}
		plot.determineFences()
	}
	var checkedPlots []*farmPlot
	var totalCost int
	plotGroups := make(map[[2]int][]*farmPlot)
	var group int
	for _, plot := range farmMap {
		if !slices.Contains(checkedPlots, plot) {
			minGroupRow := row
			groupedPlots := findFarmGroups(plot, []*farmPlot{plot})
			checkedPlots = append(checkedPlots, groupedPlots...)
			var fences int
			for _, groupedPlot := range groupedPlots {
				fences += groupedPlot.fences
				if groupedPlot.position[0] < minGroupRow {
					minGroupRow = groupedPlot.position[0]
				}
			}
			totalCost += fences * len(groupedPlots)
			plotGroups[[2]int{group, minGroupRow}] = groupedPlots
			group++
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", totalCost)
	totalCost = 0
	for key, group := range plotGroups {
		area, turns := labyrinthRobotFunction(group, plotGroups, farmMap, key[1])
		totalCost += area * turns
	}
	fmt.Printf("Part 2 Answer: %d\n", totalCost)
}

func findFarmGroups(plot *farmPlot, groupedPlots []*farmPlot) []*farmPlot {
	for _, neighbor := range plot.neighbors {
		if neighbor.letter == plot.letter && !slices.Contains(groupedPlots, neighbor) {
			groupedPlots = append(groupedPlots, neighbor)
			groupedPlots = findFarmGroups(neighbor, groupedPlots)
		}
	}
	return groupedPlots
}

func labyrinthRobotFunction(group []*farmPlot, plotGroups map[[2]int][]*farmPlot, farm map[[2]int]*farmPlot, minRow int) (int, int) {
	var robotStartingPosition [2]int
	var turns int
	var neighborsVisited []*farmPlot
	labyrinthRobot := robot{facing: 4, moveDirections: [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}}
	for _, plot := range group {
		if plot.position[0] == minRow {
			robotStartingPosition = [2]int{plot.position[0] - 1, plot.position[1]}
			labyrinthRobot.position = [2]int{plot.position[0] - 1, plot.position[1]}
			break
		}
	}
	var returnedToStartingPosition bool
	for !returnedToStartingPosition {
		if slices.Contains(group, farm[[2]int{labyrinthRobot.position[0] + labyrinthRobot.moveDirections[labyrinthRobot.facing%4][0], labyrinthRobot.position[1] + labyrinthRobot.moveDirections[labyrinthRobot.facing%4][1]}]) {
			labyrinthRobot.facing--
			turns++
		} else {
			labyrinthRobot.moveRobot()
			neighborsVisited = append(neighborsVisited, farm[labyrinthRobot.position])
			if !slices.Contains(group, farm[[2]int{labyrinthRobot.position[0] + labyrinthRobot.moveDirections[(labyrinthRobot.facing+1)%4][0], labyrinthRobot.position[1] + labyrinthRobot.moveDirections[(labyrinthRobot.facing+1)%4][1]}]) {
				labyrinthRobot.facing++
				turns++
			}
		}
		returnedToStartingPosition = (labyrinthRobot.position == robotStartingPosition)
	}
	for _, plot := range group {
		for _, neighbor := range plot.neighbors {
			if !slices.Contains(neighborsVisited, neighbor) && !slices.Contains(group, neighbor) {
				for key, internalGroup := range plotGroups {
					if slices.Contains(internalGroup, neighbor) {
						neighborsVisited = append(neighborsVisited, internalGroup...)
						_, turnsToAdd := labyrinthRobotFunction(internalGroup, plotGroups, farm, key[1])
						turns += turnsToAdd
					}
				}
			}
		}
	}
	return len(group), turns
}
