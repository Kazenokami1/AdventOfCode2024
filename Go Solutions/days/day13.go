package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day13() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day13Sample.txt"
	} else {
		fileName = "../inputfiles/Day13.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var machineStrings []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			machineStrings = append(machineStrings, scanner.Text())
		}
	}
	var machines []*clawMachine
	for i := 0; i < len(machineStrings); i += 3 {
		abuttonStrings := strings.Split(machineStrings[i], "+")
		bbuttonStrings := strings.Split(machineStrings[i+1], "+")
		prizeStrings := strings.Split(machineStrings[i+2], "=")
		aComma := strings.Index(abuttonStrings[1], ",")
		aButtonX, _ := strconv.Atoi(abuttonStrings[1][0:aComma])
		aButtonY, _ := strconv.Atoi(abuttonStrings[2])
		bComma := strings.Index(bbuttonStrings[1], ",")
		bButtonX, _ := strconv.Atoi(bbuttonStrings[1][0:bComma])
		bButtonY, _ := strconv.Atoi(bbuttonStrings[2])
		prizeComma := strings.Index(prizeStrings[1], ",")
		prizeX, _ := strconv.Atoi(prizeStrings[1][0:prizeComma])
		prizeY, _ := strconv.Atoi(prizeStrings[2])
		machines = append(machines, &clawMachine{aButton: [2]float64{float64(aButtonX), float64(aButtonY)}, bButton: [2]float64{float64(bButtonX), float64(bButtonY)}, prize: [2]float64{float64(prizeX), float64(prizeY)}})
	}
	var tokensSpent int
	for _, machine := range machines {
		machine.calcButtonPresses()
		tokensSpent += machine.tokensForPrize
	}
	fmt.Printf("Part 1 Answer: %d\n", tokensSpent)
	for _, machine := range machines {
		machine.prize = [2]float64{machine.prize[0] + 10000000000000, machine.prize[1] + 10000000000000}
		machine.tokensForPrize = 0
	}
	tokensSpent = 0
	for _, machine := range machines {
		machine.calcButtonPresses()
		tokensSpent += machine.tokensForPrize
	}
	fmt.Printf("Part 2 Answer: %d\n", tokensSpent)
}
