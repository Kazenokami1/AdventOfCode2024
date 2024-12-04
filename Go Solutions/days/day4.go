package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day4() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day4Sample.txt"
	} else {
		fileName = "../inputfiles/Day4.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	var wordSearch []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordSearch = append(wordSearch, scanner.Text())
	}
	xmasCount := slidingWindow(wordSearch, 4, false)
	for _, line := range wordSearch {
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	for i := 0; i < len(wordSearch[0]); i++ {
		var line string
		for j := 0; j < len(wordSearch); j++ {
			line += string(wordSearch[j][i])
		}
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	fmt.Printf("Part 1 Answer: %d\n", xmasCount)
	xmasCount = slidingWindow(wordSearch, 3, true)
	fmt.Printf("Part 2 Answer: %d\n", xmasCount)
}

func slidingWindow(wordSearch []string, windowSize int, part2 bool) int {
	var xmasCount int
	for i := 0; i <= len(wordSearch[0])-windowSize; i++ {
		for j := 0; j <= len(wordSearch)-windowSize; j++ {
			var window []string
			for k := 0; k < windowSize; k++ {
				window = append(window, wordSearch[j+k][i:windowSize+i])
			}
			if part2 {
				//check center letter
				if window[1][1] == 'A' {
					var line string
					var lineTwo string
					for m := 0; m < windowSize; m++ {
						line += string(window[m][m])
						lineTwo += string(window[m][windowSize-m-1])
					}
					if (line == "MAS" || line == "SAM") && (lineTwo == "MAS" || lineTwo == "SAM") {
						xmasCount++
					}
				}
			} else {
				var diagonal string
				var diagonalTwo string
				for m := 0; m < windowSize; m++ {
					diagonal += string(window[m][m])
					diagonalTwo += string(window[windowSize-m-1][m])
				}
				if diagonal == "XMAS" || diagonal == "SAMX" {
					xmasCount++
				}
				if diagonalTwo == "XMAS" || diagonalTwo == "SAMX" {
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}
