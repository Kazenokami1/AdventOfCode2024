package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * I realize this code isn't the cleanest lol, but it works and works quickly :)
 * Part 1 is pretty much brute force and because I like having part 1 and part 2 in the same file I didn't want to change it
 * Part 2 does the solving of the stones instead of brute force
 * -- Known Errors:
 * --- Ideally the solving for Part 2 should solve Part 1 if you change 75 to 25.
 * --- The sample data works as expected returning the correct value for Part 1 and the same value for Part 2 with or without Part 1
 * --- The Real Data without Part 1 returns the wrong answer for both Part 1 and Part 2
 */

func Day11() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day11Sample.txt"
	} else {
		fileName = "../inputfiles/Day11.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var stones []stone
	for scanner.Scan() {
		stoneNumbers := strings.Split(scanner.Text(), " ")
		for _, val := range stoneNumbers {
			number, _ := strconv.Atoi(val)
			stones = append(stones, stone{number: number, blinks: 0})
		}
	}
	for i := 0; i < len(stones); i++ {
		for blinks := stones[i].blinks; blinks < 25; blinks++ {
			stones[i].blinks++
			if stones[i].number == 0 {
				stones[i].number = 1
			} else if len(strconv.Itoa(stones[i].number))%2 == 0 {
				stringNumber := strconv.Itoa(stones[i].number)
				stoneOneNumber, _ := strconv.Atoi(stringNumber[0 : len(stringNumber)/2])
				stoneTwoNumber, _ := strconv.Atoi(stringNumber[len(stringNumber)/2:])
				stones[i].number = stoneOneNumber
				stones = append(stones, stone{number: stoneTwoNumber, blinks: stones[i].blinks})
			} else {
				stones[i].number *= 2024
			}
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", len(stones))
	solvedStones := make(map[stone]int)
	var emptyStone stone
	var totalStones int
	for len(stones) > 0 {
		stoneToIncrement := stones[len(stones)-1]
		stones = stones[0 : len(stones)-1]
		var stoneUnsolved bool
		for !stoneUnsolved {
			if val, ok := solvedStones[stoneToIncrement]; !ok {
				stoneUnsolved = true
			} else {
				totalStones += val
				stoneToIncrement = emptyStone
			}
		}
		if stoneToIncrement != emptyStone {
			if stoneToIncrement.blinks == 75 {
				solvedStones[stoneToIncrement] = 1
				totalStones++
			} else {
				stonesReturned := incrementStone(stoneToIncrement)
				var solvedLength int
				for _, checkStone := range stonesReturned {
					if val, ok := solvedStones[checkStone]; !ok {
						solvedLength = 0
						break
					} else {
						solvedLength += val
					}
				}
				if solvedLength > 0 {
					solvedStones[stoneToIncrement] = solvedLength
					totalStones += solvedLength
				} else {
					stones = append(stones, stonesReturned...)
				}
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", totalStones)
}

func incrementStone(currentStone stone) []stone {
	currentStone.blinks++
	if currentStone.number == 0 {
		currentStone.number = 1
		return []stone{currentStone}
	} else if len(strconv.Itoa(currentStone.number))%2 == 0 {
		stringNumber := strconv.Itoa(currentStone.number)
		stoneOneNumber, _ := strconv.Atoi(stringNumber[0 : len(stringNumber)/2])
		stoneTwoNumber, _ := strconv.Atoi(stringNumber[len(stringNumber)/2:])
		currentStone.number = stoneOneNumber
		return []stone{currentStone, {number: stoneTwoNumber, blinks: currentStone.blinks}}
	} else {
		currentStone.number *= 2024
		return []stone{currentStone}
	}
}
