package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	For Part 2 I looked and found that my robots formed general lines between rows 40 and 72
	every 103 seconds starting at second 78. So I just but a breakpoint in my code after it printed those rows,
	checked to see if it was a tree, then hit continue if it wasn't until I found my tree. The tree is easily identifiable
	but I dislike part 2 today because there's no condition of what a Christmas tree actually is or how you should find it
	without just putting human eyes on it.

	To do the same with yours, find out when your robots basically repeat a large amount of them like mine in a top/bottom row,
	figure out how many seconds you need to subtract because they didn't start there, and change the printing accordingly.
*/

func Day14() {
	var fileName string
	var mapHeight int
	var mapWidth int
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day14Sample.txt"
		mapHeight = 7
		mapWidth = 11
	} else {
		fileName = "../inputfiles/Day14.txt"
		mapHeight = 103
		mapWidth = 101
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var robotsPart1 []*robot
	var robotsPart2 []*robot
	for scanner.Scan() {
		robotStrings := strings.Split(scanner.Text(), "=")
		robotPosX, _ := strconv.Atoi(robotStrings[1][0:strings.Index(robotStrings[1], ",")])
		robotPosY, _ := strconv.Atoi(robotStrings[1][strings.Index(robotStrings[1], ",")+1 : strings.Index(robotStrings[1], " ")])
		robotVelX, _ := strconv.Atoi(robotStrings[2][0:strings.Index(robotStrings[2], ",")])
		robotVelY, _ := strconv.Atoi(robotStrings[2][strings.Index(robotStrings[2], ",")+1:])
		robotsPart1 = append(robotsPart1, &robot{position: [2]int{robotPosX, robotPosY}, velocity: [2]int{robotVelX, robotVelY}})
		robotsPart2 = append(robotsPart2, &robot{position: [2]int{robotPosX, robotPosY}, velocity: [2]int{robotVelX, robotVelY}})
	}
	quadrants := make(map[int][]*robot)
	for _, r := range robotsPart1 {
		r.calcPosition(100, mapWidth, mapHeight)
		if r.position[0] < mapWidth/2 && r.position[1] < mapHeight/2 {
			quadrants[1] = append(quadrants[1], r)
		} else if r.position[0] > mapWidth/2 && r.position[1] < mapHeight/2 {
			quadrants[2] = append(quadrants[2], r)
		} else if r.position[0] < mapWidth/2 && r.position[1] > mapHeight/2 {
			quadrants[3] = append(quadrants[3], r)
		} else if r.position[0] > mapWidth/2 && r.position[1] > mapHeight/2 {
			quadrants[4] = append(quadrants[4], r)
		}
	}
	safetyFactor := 1
	for _, quadrant := range quadrants {
		safetyFactor *= len(quadrant)
	}
	fmt.Printf("Part 1 Answer: %d\n", safetyFactor)
	var seconds int
	roomMap := make(map[[2]int]string)
	var previousSeconds int
	for {
		seconds++
		if seconds%103-78 == 0 {
			for row := 0; row < mapHeight; row++ {
				for col := 0; col < mapWidth; col++ {
					roomMap[[2]int{col, row}] = " "
				}
			}
			for _, r := range robotsPart2 {
				r.calcPosition(seconds-previousSeconds, mapWidth, mapHeight)
				roomMap[r.position] = "*"
			}
			for row := 0; row < mapHeight; row++ {
				for col := 0; col < mapWidth; col++ {
					if row > 39 && row < 73 {
						fmt.Print(roomMap[[2]int{col, row}])
					}
				}
				if row > 39 && row < 73 {
					fmt.Println()
				}
			}
			previousSeconds = seconds
			fmt.Printf("Number of Seconds Passed: %d\n", seconds)
		}
	}
}
