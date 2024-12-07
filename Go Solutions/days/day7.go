package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
I like this solution but it does take a bit (~10 seconds) to run due to the number of loops it takes to convert to different bases
and pad enough 0s to the base x representation of the int
*/
func Day7() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day7Sample.txt"
	} else {
		fileName = "../inputfiles/Day7.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var calibrations [][]int64
	for scanner.Scan() {
		var calibration []int64
		intStrings := strings.Split(scanner.Text(), " ")
		for _, value := range intStrings {
			intValue, _ := strconv.ParseInt(strings.Trim(value, ":"), 0, 0)
			calibration = append(calibration, intValue)
		}
		calibrations = append(calibrations, calibration)
	}
	var totalSumCalibrations int64
	for _, calibration := range calibrations {
		numberOfOperators := len(calibration) - 2
		for i := 0; i < int(math.Pow(2, float64(len(calibration)-2))); i++ {
			currentCalibration := calibration[1]
			baseString := convertToBase(i, 2)
			for len(baseString) < numberOfOperators {
				baseString = "0" + baseString
			}
			for j := 0; j < numberOfOperators; j++ {
				switch baseString[j] {
				case '0':
					currentCalibration += calibration[j+2]
				case '1':
					currentCalibration *= calibration[j+2]
				}
			}
			if currentCalibration == calibration[0] {
				totalSumCalibrations += calibration[0]
				break
			}
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", totalSumCalibrations)
	totalSumCalibrations = 0
	for _, calibration := range calibrations {
		numberOfOperators := len(calibration) - 2
		for i := 0; i < int(math.Pow(3, float64(len(calibration)-2))); i++ {
			currentCalibration := calibration[1]
			baseString := convertToBase(i, 3)
			for len(baseString) < numberOfOperators {
				baseString = "0" + baseString
			}
			for j := 0; j < numberOfOperators; j++ {
				switch baseString[j] {
				case '0':
					currentCalibration += calibration[j+2]
				case '1':
					currentCalibration *= calibration[j+2]
				case '2':
					currentCalibration, _ = strconv.ParseInt(strconv.Itoa(int(currentCalibration))+strconv.Itoa(int(calibration[j+2])), 0, 0)
				}
			}
			if currentCalibration == calibration[0] {
				totalSumCalibrations += calibration[0]
				break
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", totalSumCalibrations)
}

func convertToBase(numberToConvert int, base int) string {
	var convertedNumber string
	if numberToConvert == 0 {
		return "0"
	}
	for numberToConvert > 0 {
		remainder := numberToConvert % base
		numberToConvert /= base
		convertedNumber = strconv.Itoa(remainder) + convertedNumber
	}
	return convertedNumber
}
