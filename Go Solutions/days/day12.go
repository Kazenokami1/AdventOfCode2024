package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
	As a note, Part 2 of this will return the wrong cost if a group fully contained a group of plots that also fully contained another group of plots.
	A fix for this would be to move the currently recursive function to a separate function that basically didn't check for an internal group.
	In other words, copy the labyrinthRobotFunction, and cut lines 121-133 from that second function.
	The first call (76) would still call labyrinthRobotFunction, the 2nd call (127) would call the new function instead.
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
		//Check to see if block in front of us is in group
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
		//Check to see if we have a block to our right
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
