package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	numbers := []int{}

	for _, nStr := range strings.Split(data, ",") {
		n, err := strconv.Atoi(nStr)
		checkError(err)

		numbers = append(numbers, n)
	}

	return numbers
}

func getNth(numbers []int, n int) int {
	previousTurns := map[int]int{}

	previousNum := 0

	currTurn := 1

	// Populate previousTurns with the starting numbers
	for _, num := range numbers {
		if _, ok := previousTurns[num]; !ok {
			previousTurns[num] = currTurn + 1
		}
		previousNum = num
		currTurn++
	}

	for currTurn = currTurn; currTurn <= n; currTurn++ {
		if _, ok := previousTurns[previousNum]; !ok {
			previousTurns[previousNum] = currTurn
			previousNum = 0
		} else if num, ok := previousTurns[previousNum]; ok {
			previousTurns[previousNum] = currTurn
			previousNum = currTurn - num
		}
	}

	return previousNum
}

func part1(numbers []int) {
	fmt.Println("PART1:", getNth(numbers, 2020))
}

func part2(numbers []int) {
	fmt.Println("PART2:", getNth(numbers, 30000000))
}

func main() {
	data, err := ioutil.ReadFile("day15.in")
	checkError(err)

	numbers := parseNumbers(string(data))

	// PART 1
	part1(numbers)

	// Part 2
	part2(numbers)
}
