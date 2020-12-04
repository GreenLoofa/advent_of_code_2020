package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func countNumCollisions(rows []string, rightPattern int, downPattern int) int {
	rowsLen := len(rows)

	numCollisions := 0

	j := rightPattern
	for i := downPattern; i < rowsLen; i += downPattern {
		if rows[i][j] == '#' {
			numCollisions++
		}

		j = (j + rightPattern) % 31
	}
	return numCollisions
}

func main() {
	data, err := ioutil.ReadFile("day3.in")
	checkError(err)

	rows := strings.Split(string(data), "\n")

	part1Collisions := countNumCollisions(rows, 3, 1)

	fmt.Println("PART1:", part1Collisions)

	part2Answer :=
		part1Collisions *
			countNumCollisions(rows, 1, 1) *
			countNumCollisions(rows, 5, 1) *
			countNumCollisions(rows, 7, 1) *
			countNumCollisions(rows, 1, 2)

	fmt.Println("PART2:", part2Answer)
}
