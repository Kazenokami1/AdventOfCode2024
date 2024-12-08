package days

import (
	"bufio"
	"fmt"
	"os"
)

/*
I like this solution but it does take a bit (~10 seconds) to run due to the number of loops it takes to convert to different bases
and pad enough 0s to the base x representation of the int
*/
func Day8() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day8Sample.txt"
	} else {
		fileName = "../inputfiles/Day8.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	var mapRows []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mapRows = append(mapRows, scanner.Text())
	}
	antiNodes2 := make(map[[2]int]struct{})
	frequencyMap := make(map[rune][][2]int)
	for row, mapRow := range mapRows {
		for col, letter := range mapRow {
			if letter != '.' {
				frequencyMap[letter] = append(frequencyMap[letter], [2]int{row, col})
				antiNodes2[[2]int{row, col}] = struct{}{}
			}
		}
	}
	maxRow := len(mapRows)
	maxCol := len(mapRows[0])
	antiNodes := make(map[[2]int]struct{})
	for _, frequency := range frequencyMap {
		for i := 0; i < len(frequency)-1; i++ {
			for j := i + 1; j < len(frequency); j++ {
				slope := [2]int{frequency[i][0] - frequency[j][0], frequency[i][1] - frequency[j][1]}
				point1 := [2]int{frequency[i][0] + slope[0], frequency[i][1] + slope[1]}
				point2 := [2]int{frequency[j][0] - slope[0], frequency[j][1] - slope[1]}
				if !(point1[0] < 0 || point1[0] >= maxRow || point1[1] < 0 || point1[1] >= maxCol) {
					antiNodes[point1] = struct{}{}
					antiNodes2[point1] = struct{}{}
					point1 = [2]int{point1[0] + slope[0], point1[1] + slope[1]}
					for point1[0] >= 0 && point1[1] >= 0 && point1[0] < maxRow && point1[1] < maxCol {
						antiNodes2[point1] = struct{}{}
						point1 = [2]int{point1[0] + slope[0], point1[1] + slope[1]}
					}
				}
				if !(point2[0] < 0 || point2[0] >= maxRow || point2[1] < 0 || point2[1] >= maxCol) {
					antiNodes[point2] = struct{}{}
					antiNodes2[point2] = struct{}{}
					point2 = [2]int{point2[0] - slope[0], point2[1] - slope[1]}
					for point2[0] >= 0 && point2[1] >= 0 && point2[0] < maxRow && point2[1] < maxCol {
						antiNodes2[point2] = struct{}{}
						point2 = [2]int{point2[0] - slope[0], point2[1] - slope[1]}
					}
				}
			}
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", len(antiNodes))
	fmt.Printf("Part 2 Answer: %d\n", len(antiNodes2))
}
