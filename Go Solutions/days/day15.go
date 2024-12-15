package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Day15() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day15Sample.txt"
	} else {
		fileName = "../inputfiles/Day15.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var robotDirections bool
	var robotMovements string
	var row int
	var robotInitialPosition [2]int
	caveMapPart1 := make(map[[2]int]string)
	caveMapPart2 := make(map[[2]int]string)
	for scanner.Scan() {
		var part2ColumnAdd int
		if scanner.Text() == "" {
			robotDirections = true
		} else if robotDirections {
			robotMovements += scanner.Text()
		} else {
			for col, char := range scanner.Text() {
				caveMapPart1[[2]int{row, col}] = string(char)
				caveMapPart2[[2]int{row, part2ColumnAdd}] = string(char)
				part2ColumnAdd++
				if char == '@' {
					robotInitialPosition = [2]int{row, col}
					caveMapPart2[[2]int{row, part2ColumnAdd}] = "."
				} else if char == '#' {
					caveMapPart2[[2]int{row, part2ColumnAdd}] = "#"
				} else if char == '.' {
					caveMapPart2[[2]int{row, part2ColumnAdd}] = "."
				} else if char == 'O' {
					caveMapPart2[[2]int{row, part2ColumnAdd - 1}] = "["
					caveMapPart2[[2]int{row, part2ColumnAdd}] = "]"
				}
				part2ColumnAdd++
			}
			row++
		}
	}
	gps := solvePart1(caveMapPart1, robotMovements, robotInitialPosition)
	robotInitialPosition = [2]int{robotInitialPosition[0], robotInitialPosition[1] * 2}
	fmt.Printf("Part 1 Answer: %d\n", gps)
	gps = solvePart2(caveMapPart2, robotMovements, robotInitialPosition)
	fmt.Printf("Part 2 Answer: %d\n", gps)
}

func solvePart1(caveMap map[[2]int]string, robotMovements string, robotPosition [2]int) int {
	for _, robotMovement := range robotMovements {
		if robotMovement == '^' {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0] - spotsToCheck, robotPosition[1]}]
			for nextSpot != "#" {
				if nextSpot == "O" {
					spotsToCheck++
					nextSpot = caveMap[[2]int{robotPosition[0] - spotsToCheck, robotPosition[1]}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0] - i, robotPosition[1]}], caveMap[[2]int{robotPosition[0] - i + 1, robotPosition[1]}] = caveMap[[2]int{robotPosition[0] - i + 1, robotPosition[1]}], caveMap[[2]int{robotPosition[0] - i, robotPosition[1]}]
					}
					robotPosition[0]--
					break
				}
			}
		} else if robotMovement == '>' {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0], robotPosition[1] + spotsToCheck}]
			for nextSpot != "#" {
				if nextSpot == "O" {
					spotsToCheck++
					nextSpot = caveMap[[2]int{robotPosition[0], robotPosition[1] + spotsToCheck}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0], robotPosition[1] + i}], caveMap[[2]int{robotPosition[0], robotPosition[1] + i - 1}] = caveMap[[2]int{robotPosition[0], robotPosition[1] + i - 1}], caveMap[[2]int{robotPosition[0], robotPosition[1] + i}]
					}
					robotPosition[1]++
					break
				}
			}
		} else if robotMovement == '<' {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0], robotPosition[1] - spotsToCheck}]
			for nextSpot != "#" {
				if nextSpot == "O" {
					spotsToCheck++
					nextSpot = caveMap[[2]int{robotPosition[0], robotPosition[1] - spotsToCheck}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0], robotPosition[1] - i}], caveMap[[2]int{robotPosition[0], robotPosition[1] - i + 1}] = caveMap[[2]int{robotPosition[0], robotPosition[1] - i + 1}], caveMap[[2]int{robotPosition[0], robotPosition[1] - i}]
					}
					robotPosition[1]--
					break
				}
			}
		} else {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0] + spotsToCheck, robotPosition[1]}]
			for nextSpot != "#" {
				if nextSpot == "O" {
					spotsToCheck++
					nextSpot = caveMap[[2]int{robotPosition[0] + spotsToCheck, robotPosition[1]}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0] + i, robotPosition[1]}], caveMap[[2]int{robotPosition[0] + i - 1, robotPosition[1]}] = caveMap[[2]int{robotPosition[0] + i - 1, robotPosition[1]}], caveMap[[2]int{robotPosition[0] + i, robotPosition[1]}]
					}
					robotPosition[0]++
					break
				}
			}
		}
	}
	var gps int
	for k, v := range caveMap {
		if v == "O" {
			gps += 100*k[0] + k[1]
		}
	}
	return gps
}

func solvePart2(caveMap map[[2]int]string, robotMovements string, robotPosition [2]int) int {
	for _, robotMovement := range robotMovements {
		if robotMovement == '^' {
			nextSpots := caveMap[[2]int{robotPosition[0] - 1, robotPosition[1]}]
			spotsToMove := [][2]int{robotPosition}
			for !strings.Contains(nextSpots, "#") {
				if !strings.Contains(nextSpots, "[") && !strings.Contains(nextSpots, "]") {
					slices.Reverse(spotsToMove)
					for _, spot := range spotsToMove {
						caveMap[spot], caveMap[[2]int{spot[0] - 1, spot[1]}] = caveMap[[2]int{spot[0] - 1, spot[1]}], caveMap[spot]
					}
					robotPosition[0]--
					break
				} else {
					nextSpots = ""
					for _, spot := range spotsToMove {
						if !slices.Contains(spotsToMove, [2]int{spot[0] - 1, spot[1]}) {
							if caveMap[[2]int{spot[0] - 1, spot[1]}] == "[" {
								nextSpots += "[]"
								spotsToMove = append(spotsToMove, [][2]int{{spot[0] - 1, spot[1]}, {spot[0] - 1, spot[1] + 1}}...)
							} else if caveMap[[2]int{spot[0] - 1, spot[1]}] == "]" {
								nextSpots += "[]"
								spotsToMove = append(spotsToMove, [][2]int{{spot[0] - 1, spot[1]}, {spot[0] - 1, spot[1] - 1}}...)
							} else if caveMap[[2]int{spot[0] - 1, spot[1]}] == "." {
								nextSpots += "."
							} else {
								nextSpots = "#"
							}
						}
					}
				}
			}
		} else if robotMovement == '>' {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0], robotPosition[1] + spotsToCheck}]
			for nextSpot != "#" {
				if nextSpot == "[" {
					spotsToCheck += 2
					nextSpot = caveMap[[2]int{robotPosition[0], robotPosition[1] + spotsToCheck}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0], robotPosition[1] + i}], caveMap[[2]int{robotPosition[0], robotPosition[1] + i - 1}] = caveMap[[2]int{robotPosition[0], robotPosition[1] + i - 1}], caveMap[[2]int{robotPosition[0], robotPosition[1] + i}]
					}
					robotPosition[1]++
					break
				}
			}
		} else if robotMovement == '<' {
			spotsToCheck := 1
			nextSpot := caveMap[[2]int{robotPosition[0], robotPosition[1] - spotsToCheck}]
			for nextSpot != "#" {
				if nextSpot == "]" {
					spotsToCheck += 2
					nextSpot = caveMap[[2]int{robotPosition[0], robotPosition[1] - spotsToCheck}]
				} else if nextSpot == "." {
					for i := spotsToCheck; i > 0; i-- {
						caveMap[[2]int{robotPosition[0], robotPosition[1] - i}], caveMap[[2]int{robotPosition[0], robotPosition[1] - i + 1}] = caveMap[[2]int{robotPosition[0], robotPosition[1] - i + 1}], caveMap[[2]int{robotPosition[0], robotPosition[1] - i}]
					}
					robotPosition[1]--
					break
				}
			}
		} else {
			nextSpots := caveMap[[2]int{robotPosition[0] + 1, robotPosition[1]}]
			spotsToMove := [][2]int{robotPosition}
			for !strings.Contains(nextSpots, "#") {
				if !strings.Contains(nextSpots, "[") && !strings.Contains(nextSpots, "]") {
					slices.Reverse(spotsToMove)
					for _, spot := range spotsToMove {
						caveMap[spot], caveMap[[2]int{spot[0] + 1, spot[1]}] = caveMap[[2]int{spot[0] + 1, spot[1]}], caveMap[spot]
					}
					robotPosition[0]++
					break
				} else {
					nextSpots = ""
					for _, spot := range spotsToMove {
						if !slices.Contains(spotsToMove, [2]int{spot[0] + 1, spot[1]}) {
							if caveMap[[2]int{spot[0] + 1, spot[1]}] == "[" {
								nextSpots += "[]"
								spotsToMove = append(spotsToMove, [][2]int{{spot[0] + 1, spot[1]}, {spot[0] + 1, spot[1] + 1}}...)
							} else if caveMap[[2]int{spot[0] + 1, spot[1]}] == "]" {
								nextSpots += "[]"
								spotsToMove = append(spotsToMove, [][2]int{{spot[0] + 1, spot[1]}, {spot[0] + 1, spot[1] - 1}}...)
							} else if caveMap[[2]int{spot[0] + 1, spot[1]}] == "." {
								nextSpots += "."
							} else {
								nextSpots = "#"
							}
						}
					}
				}
			}
		}
	}
	var gps int
	for k, v := range caveMap {
		if v == "[" {
			gps += 100*k[0] + k[1]
		}
	}
	return gps
}
