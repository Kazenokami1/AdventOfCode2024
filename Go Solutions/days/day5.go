package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "../inputfiles/Day5Sample.txt"
	} else {
		fileName = "../inputfiles/Day5.txt"
	}
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	pageOrder := make(map[int64][]int64)
	var manuals [][]int64
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			pages := strings.Split(scanner.Text(), "|")
			pageOne, _ := strconv.ParseInt(pages[0], 0, 0)
			pageTwo, _ := strconv.ParseInt(pages[1], 0, 0)
			pageOrder[pageOne] = append(pageOrder[pageOne], pageTwo)
		} else if scanner.Text() != "" {
			pages := strings.Split(scanner.Text(), ",")
			var manual []int64
			for i := 0; i < len(pages); i++ {
				page, _ := strconv.ParseInt(pages[i], 0, 0)
				manual = append(manual, page)
			}
			manuals = append(manuals, manual)
		}
	}
	var correctManuals [][]int64
	var incorrectManuals [][]int64
	for _, manual := range manuals {
		correctOrder := true
		for i := 0; i < len(manual)-1; i++ {
			if !slices.Contains(pageOrder[manual[i]], manual[i+1]) {
				correctOrder = false
				break
			}
		}
		if correctOrder {
			correctManuals = append(correctManuals, manual)
		} else {
			incorrectManuals = append(incorrectManuals, manual)
		}
	}
	var middleAdditions int64
	for _, manual := range correctManuals {
		middleAdditions += manual[len(manual)/2]
	}
	fmt.Printf("Part 1 Answer: %d\n", middleAdditions)
	var correctedManuals [][]int64
	for _, manual := range incorrectManuals {
		var ordered bool
		for !ordered {
			ordered = true
			for i := 0; i < len(manual)-1; i++ {
				if slices.Contains(pageOrder[manual[i+1]], manual[i]) {
					ordered = false
					page := manual[i]
					manual = slices.Delete(manual, i, i+1)
					manual = append(manual, page)
				}
			}
			if ordered {
				correctedManuals = append(correctedManuals, manual)
			}
		}
	}
	middleAdditions = 0
	for _, manual := range correctedManuals {
		middleAdditions += manual[len(manual)/2]
	}
	fmt.Printf("Part 2 Answer: %d\n", middleAdditions)
}
