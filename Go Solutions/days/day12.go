package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

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
	plotGroups := make(map[int][]*farmPlot)
	var group int
	for _, plot := range farmMap {
		if !slices.Contains(checkedPlots, plot) {
			groupedPlots := findFarmGroups(plot, []*farmPlot{plot})
			checkedPlots = append(checkedPlots, groupedPlots...)
			var fences int
			for _, groupedPlot := range groupedPlots {
				fences += groupedPlot.fences
			}
			totalCost += fences * len(groupedPlots)
			plotGroups[group] = groupedPlots
			group++
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", totalCost)
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
