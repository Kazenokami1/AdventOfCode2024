package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day2() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day2Sample.txt"
	} else {
		fileName = "../inputfiles/Day2.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	var reports [][]int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		var levelInts []int64
		for _, level := range levels {
			levelInt, _ := strconv.ParseInt(level, 0, 0)
			levelInts = append(levelInts, levelInt)
		}
		reports = append(reports, levelInts)
	}
	var safeReports int
	for _, report := range reports {
		if checkReportSafety(report) {
			safeReports++
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", safeReports)
	safeReports = 0
	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			newReport := slices.Clone(report[0:i])
			if i != len(report)-1 {
				newReport = append(newReport, report[i+1:]...)
			}
			if checkReportSafety(newReport) {
				safeReports++
				break
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", safeReports)
}

func checkReportSafety(report []int64) bool {
	if slices.IsSorted(report) {
		return checkDifferenceSafety(report)
	} else {
		slices.Reverse(report)
		if slices.IsSorted(report) {
			return checkDifferenceSafety(report)
		} else {
			return false
		}
	}
}

func checkDifferenceSafety(report []int64) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] == report[i+1] || report[i+1]-report[i] > 3 {
			return false
		}
	}
	return true
}
