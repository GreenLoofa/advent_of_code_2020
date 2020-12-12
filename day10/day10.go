package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func parseNumbers(data string) []int {
	var numbers = []int{0}

	for _, numStr := range strings.Split(data, "\n") {
		num, err := strconv.Atoi(numStr)
		checkError(err)

		numbers = append(numbers, num)
	}
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	return numbers
}

func part1(numbers []int) int {
	numDiff1 := 0
	numDiff3 := 0

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff == 1 {
			numDiff1++
		}
		if diff == 3 {
			numDiff3++
		}
	}

	return numDiff1 * numDiff3
}

func part2(numbers []int) float64 {
	// Did not have time to solve
	return 0
}

func main() {
	data, err := ioutil.ReadFile("day10.in")
	checkError(err)

	numbers := parseNumbers(string(data))

	// Part1
	fmt.Println("PART1:", part1(numbers))

	// Part 2
	// fmt.Println("PART2:", part2(numbers))
}
