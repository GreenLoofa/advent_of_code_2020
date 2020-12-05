package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func getSeatID(boardingPass string) int {
	upperRow := 127
	lowerRow := 0

	for _, direction := range boardingPass[0:7] {
		diff := upperRow - lowerRow
		rem := diff % 2
		if direction == 'B' {
			lowerRow += (diff / 2) + rem
		}
		if direction == 'F' {
			upperRow -= (diff / 2) + rem
		}
	}

	upperColumn := 7
	lowerColumn := 0

	for _, direction := range boardingPass[7:] {
		diff := upperColumn - lowerColumn
		rem := diff % 2
		if direction == 'R' {
			lowerColumn += (diff / 2) + rem
		}
		if direction == 'L' {
			upperColumn -= (diff / 2) + rem
		}
	}

	return upperRow*8 + upperColumn
}

func isValidID(seatID int) bool {
	columns := [8]int{0, 1, 2, 3, 4, 5, 6, 7}

	for _, c := range columns {
		if (seatID-c)%8 == 0 {
			return true
		}
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile("day5.in")
	checkError(err)

	boardingPasses := strings.Split(string(data), "\n")

	// Part 1
	maxSeatID := 0

	for _, boardingPass := range boardingPasses {
		checkSeatID := getSeatID(boardingPass)
		if checkSeatID > maxSeatID {
			maxSeatID = checkSeatID
		}
	}

	fmt.Println("PART1:", maxSeatID)

	// Part 2
	var seatIDs = make([]int, len(boardingPasses))

	for i, boardingPass := range boardingPasses {
		seatIDs[i] = getSeatID(boardingPass)
	}
	sort.Ints(seatIDs)

	numIDs := len(seatIDs)

	for i := 0; i < numIDs-1; i++ {
		if seatIDs[i+1]-seatIDs[i] == 2 && isValidID(seatIDs[i]+1) {
			fmt.Println("PART2:", seatIDs[i]+1)
			break
		}
	}
}
