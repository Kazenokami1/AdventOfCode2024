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

func Day1() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day1Sample.txt"
	} else {
		fileName = "inputfiles/Day1.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var firstList []float64
	var secondList []float64
	for scanner.Scan() {
		lists := strings.Split(scanner.Text(), "   ")
		listOneInt, _ := strconv.ParseFloat(lists[0], 0)
		listTwoInt, _ := strconv.ParseFloat(lists[1], 0)
		firstList = append(firstList, listOneInt)
		secondList = append(secondList, listTwoInt)
	}
	slices.Sort(firstList)
	slices.Sort(secondList)
	var totalDifference float64
	for i := 0; i < len(firstList); i++ {
		totalDifference += math.Abs(firstList[i] - secondList[i])
	}
	fmt.Printf("Part 1 Answer: %d\n", int(totalDifference))
	var similarityScore float64
	for _, value := range firstList {
		var totalSimilar float64
		for _, valueTwo := range secondList {
			if value == valueTwo {
				totalSimilar++
			}
		}
		similarityScore += value * totalSimilar
	}
	fmt.Printf("Part 2 Answer: %d\n", int(similarityScore))
}
