package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Definitely some optimization to be had, takes about 10s for both parts but it works so...
*/

func Day9() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day9Sample.txt"
	} else {
		fileName = "../inputfiles/Day9.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var fileSystem string
	for scanner.Scan() {
		fileSystem = scanner.Text()
	}
	separatedFiles := make(map[int]interface{})
	separatedFiles2 := make(map[int]interface{})
	fileIdSizes := make(map[int]int64)
	var orderedFileIndex int
	for i, size := range fileSystem {
		sizeInt, _ := strconv.ParseInt(string(size), 0, 0)
		if i%2 == 0 {
			for j := 0; j < int(sizeInt); j++ {
				separatedFiles[orderedFileIndex] = i / 2
				separatedFiles2[orderedFileIndex] = i / 2
				orderedFileIndex++
			}
			fileIdSizes[i/2] = sizeInt
		} else {
			for j := 0; j < int(sizeInt); j++ {
				separatedFiles[orderedFileIndex] = "."
				separatedFiles2[orderedFileIndex] = "."
				orderedFileIndex++
			}
		}
	}
	for i := 0; i < len(separatedFiles)-1; i++ {
		for separatedFiles[len(separatedFiles)-1] == "." {
			delete(separatedFiles, len(separatedFiles)-1)
		}
		if separatedFiles[i] == "." {
			separatedFiles[i] = separatedFiles[len(separatedFiles)-1]
			delete(separatedFiles, len(separatedFiles)-1)
		}
	}
	var checkSum int
	for i := 0; i < len(separatedFiles); i++ {
		if separatedFiles[i] != "." {
			checkSum += i * separatedFiles[i].(int)
		}
	}
	fmt.Printf("Part 1 Answer: %d\n", checkSum)
	fileIdToMoveIndex := len(separatedFiles2) - 1
	for fileIdToMoveIndex > 0 {
		for separatedFiles2[fileIdToMoveIndex] == "." {
			fileIdToMoveIndex--
		}
		fileIdToMove := separatedFiles2[fileIdToMoveIndex]
		fileSize := int(fileIdSizes[fileIdToMove.(int)])
		for i := 0; i < fileIdToMoveIndex; i++ {
			var freeSpace int
			for separatedFiles2[i+freeSpace] == "." {
				freeSpace++
			}
			if fileSize <= freeSpace {
				for j := 0; j < fileSize; j++ {
					separatedFiles2[i+j] = fileIdToMove
					separatedFiles2[fileIdToMoveIndex-j] = "."
				}
				break
			}
		}
		fileIdToMoveIndex -= fileSize
	}
	checkSum = 0
	for i := 0; i < len(separatedFiles2); i++ {
		if separatedFiles2[i] != "." {
			checkSum += i * separatedFiles2[i].(int)
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", checkSum)
}
