package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day17() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day17Sample.txt"
	} else {
		fileName = "../inputfiles/Day17.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var registers []int
	var program []int
	for scanner.Scan() {
		if scanner.Text() != "" {
			split := strings.Split(scanner.Text(), ": ")
			registerValue, err := strconv.Atoi(split[1])
			if err != nil {
				programCommands := strings.Split(split[1], ",")
				for _, val := range programCommands {
					programVal, _ := strconv.Atoi(val)
					program = append(program, programVal)
				}
			} else {
				registers = append(registers, registerValue)
			}
		}
	}
	fmt.Print("Part 1 Answer (delete the last ,): ")
	for i := 0; i < len(program); i += 2 {
		switch program[i] {
		case 0:
			if program[i+1] <= 3 {
				registers[0] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
			} else if program[i+1] == 4 {
				registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
			} else if program[i+1] == 5 {
				registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
			} else if program[i+1] == 6 {
				registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
			}
		case 1:
			//bitwise XOR
			bitB := convertToBase(registers[1], 2)
			operand := convertToBase(program[i+1], 2)
			for len(bitB) < len(operand) {
				bitB = "0" + bitB
			}
			for len(operand) < len(bitB) {
				operand = "0" + operand
			}
			var bitWise string
			for i := 0; i < len(bitB); i++ {
				if bitB[i] == operand[i] {
					bitWise += "0"
				} else {
					bitWise += "1"
				}
			}
			storeValue, _ := strconv.ParseInt(bitWise, 2, 0)
			registers[1] = int(storeValue)
		case 2:
			if program[i+1] <= 3 {
				registers[1] = program[i+1] % 8
			} else if program[i+1] == 4 {
				registers[1] = registers[0] % 8
			} else if program[i+1] == 5 {
				registers[1] = registers[1] % 8
			} else if program[i+1] == 6 {
				registers[1] = registers[2] % 8
			}
		case 3:
			if !(registers[0] == 0) {
				i = program[i+1] - 2
			}
		case 4:
			//bitwise XOR
			bitB := convertToBase(registers[1], 2)
			bitC := convertToBase(registers[2], 2)
			for len(bitB) < len(bitC) {
				bitB = "0" + bitB
			}
			for len(bitC) < len(bitB) {
				bitC = "0" + bitC
			}
			var bitWise string
			for i := 0; i < len(bitB); i++ {
				if bitB[i] == bitC[i] {
					bitWise += "0"
				} else {
					bitWise += "1"
				}
			}
			storeValue, _ := strconv.ParseInt(bitWise, 2, 0)
			registers[1] = int(storeValue)
		case 5:
			if program[i+1] <= 3 {
				fmt.Printf("%d,", program[i+1]%8)
			} else if program[i+1] == 4 {
				fmt.Printf("%d,", registers[0]%8)
			} else if program[i+1] == 5 {
				fmt.Printf("%d,", registers[1]%8)
			} else if program[i+1] == 6 {
				fmt.Printf("%d,", registers[2]%8)
			}
		case 6:
			if program[i+1] <= 3 {
				registers[1] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
			} else if program[i+1] == 4 {
				registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
			} else if program[i+1] == 5 {
				registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
			} else if program[i+1] == 6 {
				registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
			}
		case 7:
			if program[i+1] <= 3 {
				registers[2] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
			} else if program[i+1] == 4 {
				registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
			} else if program[i+1] == 5 {
				registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
			} else if program[i+1] == 6 {
				registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
			}
		}
	}
	fmt.Println()
	var initialNumber int
	var exactCopy bool
	var programString string
	var outputString string
	for i := 0; i < len(program); i++ {
		programString += strconv.Itoa(program[i]) + ","
	}
	for !exactCopy {
		registers[0] = initialNumber
		for i := 0; i < len(program); i += 2 {
			switch program[i] {
			case 0:
				if program[i+1] <= 3 {
					registers[0] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
				} else if program[i+1] == 4 {
					registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
				} else if program[i+1] == 5 {
					registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
				} else if program[i+1] == 6 {
					registers[0] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
				}
			case 1:
				//bitwise XOR
				bitB := convertToBase(registers[1], 2)
				operand := convertToBase(program[i+1], 2)
				for len(bitB) < len(operand) {
					bitB = "0" + bitB
				}
				for len(operand) < len(bitB) {
					operand = "0" + operand
				}
				var bitWise string
				for i := 0; i < len(bitB); i++ {
					if bitB[i] == operand[i] {
						bitWise += "0"
					} else {
						bitWise += "1"
					}
				}
				storeValue, _ := strconv.ParseInt(bitWise, 2, 0)
				registers[1] = int(storeValue)
			case 2:
				if program[i+1] <= 3 {
					registers[1] = program[i+1] % 8
				} else if program[i+1] == 4 {
					registers[1] = registers[0] % 8
				} else if program[i+1] == 5 {
					registers[1] = registers[1] % 8
				} else if program[i+1] == 6 {
					registers[1] = registers[2] % 8
				}
			case 3:
				if !(registers[0] == 0) {
					i = program[i+1] - 2
				}
			case 4:
				//bitwise XOR
				bitB := convertToBase(registers[1], 2)
				bitC := convertToBase(registers[2], 2)
				for len(bitB) < len(bitC) {
					bitB = "0" + bitB
				}
				for len(bitC) < len(bitB) {
					bitC = "0" + bitC
				}
				var bitWise string
				for i := 0; i < len(bitB); i++ {
					if bitB[i] == bitC[i] {
						bitWise += "0"
					} else {
						bitWise += "1"
					}
				}
				storeValue, _ := strconv.ParseInt(bitWise, 2, 0)
				registers[1] = int(storeValue)
			case 5:
				if program[i+1] <= 3 {
					outputString += strconv.Itoa(program[i+1]%8) + ","
				} else if program[i+1] == 4 {
					outputString += strconv.Itoa(registers[0]%8) + ","
				} else if program[i+1] == 5 {
					outputString += strconv.Itoa(registers[1]) + ","
				} else if program[i+1] == 6 {
					outputString += strconv.Itoa(registers[2]%8) + ","
				}
				if outputString != programString[0:len(outputString)] {
					outputString = ""
					initialNumber++
					i = len(program)
				} else if outputString == programString {
					exactCopy = true
				}
			case 6:
				if program[i+1] <= 3 {
					registers[1] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
				} else if program[i+1] == 4 {
					registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
				} else if program[i+1] == 5 {
					registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
				} else if program[i+1] == 6 {
					registers[1] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
				}
			case 7:
				if program[i+1] <= 3 {
					registers[2] = int(float64(registers[0]) / math.Pow(2, float64(program[i+1])))
				} else if program[i+1] == 4 {
					registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[0])))
				} else if program[i+1] == 5 {
					registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[1])))
				} else if program[i+1] == 6 {
					registers[2] = int(float64(registers[0]) / math.Pow(2, float64(registers[2])))
				}
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", initialNumber)
}
