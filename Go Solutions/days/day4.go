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
	xmasCount := part1IsSoUgly(wordSearch)
	fmt.Printf("Part 1 Answer: %d\n", xmasCount)
	xmasCount = slidingWindow(wordSearch, 3, true)
	fmt.Printf("Part 2 Answer: %d\n", xmasCount)
}

func part1IsSoUgly(wordSearch []string) int {
	var xmasCount int
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
	for i := 0; i < len(wordSearch[0]); i++ {
		var line string
		for j := 0; j <= i; j++ {
			line += string(wordSearch[i-j][i-(i-j)])
		}
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	for i := len(wordSearch) - 1; i > 0; i-- {
		var line string
		for j := 0; j+i < len(wordSearch[0]); j++ {
			line += string(wordSearch[i+j][len(wordSearch[0])-1-j])
		}
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	for i := 0; i < len(wordSearch); i++ {
		var line string
		for j := 0; j <= i; j++ {
			line += string(wordSearch[i-j][len(wordSearch[0])-1-j])
		}
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	for i := len(wordSearch) - 1; i > 0; i-- {
		var line string
		for j := 0; j+i < len(wordSearch[0]); j++ {
			line += string(wordSearch[len(wordSearch)-1-j][len(wordSearch)-1-i-j])
		}
		xmasCount += strings.Count(line, "XMAS")
		xmasCount += strings.Count(line, "SAMX")
	}
	return xmasCount
}

func slidingWindow(wordSearch []string, windowSize int, findCrossing bool) int {
	var xmasCount int
	for i := 0; i <= len(wordSearch[0])-windowSize; i++ {
		for j := 0; j <= len(wordSearch)-windowSize; j++ {
			var window []string
			for k := 0; k < windowSize; k++ {
				window = append(window, wordSearch[j+k][i:windowSize+i])
			}
			if findCrossing {
				//check center letter
				if window[1][1] == 'A' {
					var line string
					var lineTwo string
					for j := 0; j < windowSize; j++ {
						line += string(window[j][j])
						lineTwo += string(window[j][windowSize-j-1])
					}
					if (line == "MAS" || line == "SAM") && (lineTwo == "MAS" || lineTwo == "SAM") {
						xmasCount++
					}
				}
			}
		}
	}
	return xmasCount
}
