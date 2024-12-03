package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day3() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day3Sample.txt"
	} else {
		fileName = "../inputfiles/Day3.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var computerInstructions string
	var computerInstructions2 string
	for scanner.Scan() {
		computerInstructions += scanner.Text()
		computerInstructions2 += scanner.Text()
	}
	var multipliedTotal int64
	for {
		index := strings.Index(computerInstructions, "mul(") + 4
		if index == 3 {
			break
		}
		computerInstructions = strings.Replace(computerInstructions, computerInstructions[0:index], "", 1)
		index = strings.Index(computerInstructions, ",")
		index2 := strings.Index(computerInstructions, ")")
		if index == -1 || index2 == -1 {
			break
		}
		possibleNumber, err := strconv.ParseInt(computerInstructions[0:index], 0, 0)
		if err != nil {
		} else {
			possibleNumber2, err := strconv.ParseInt(computerInstructions[index+1:index2], 0, 0)
			if err != nil {
			} else {
				multipliedTotal += possibleNumber * possibleNumber2
			}
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", multipliedTotal)
	multipliedTotal = 0
	multEnabled := true
	for {
		if multEnabled {
			index := strings.Index(computerInstructions2, "mul(") + 4
			dontIndex := strings.Index(computerInstructions2, "don't()")
			if dontIndex < index-4 && dontIndex != -1 {
				multEnabled = false
			} else if index == 3 {
				break
			} else {
				computerInstructions2 = strings.Replace(computerInstructions2, computerInstructions2[0:index], "", 1)
				index = strings.Index(computerInstructions2, ",")
				index2 := strings.Index(computerInstructions2, ")")
				if index == -1 || index2 == -1 {
					break
				}
				possibleNumber, err := strconv.ParseInt(computerInstructions2[0:index], 0, 0)
				if err != nil {
				} else {
					possibleNumber2, err := strconv.ParseInt(computerInstructions2[index+1:index2], 0, 0)
					if err != nil {
					} else {
						multipliedTotal += possibleNumber * possibleNumber2
					}
				}
			}
		} else {
			doIndex := strings.Index(computerInstructions2, "do()")
			if doIndex != -1 {
				computerInstructions2 = strings.Replace(computerInstructions2, computerInstructions2[0:doIndex+4], "", 1)
				multEnabled = true
			} else {
				break
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", multipliedTotal)
}
