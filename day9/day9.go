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
	var numbers = []int{}

	for _, numStr := range strings.Split(data, "\n") {
		num, err := strconv.Atoi(numStr)
		checkError(err)

		numbers = append(numbers, num)
	}
	return numbers
}

func isSumOfPrevious(numToCheck int, numbers []int) bool {
	for _, num := range numbers {
		for _, num2 := range numbers {
			if num+num2 == numToCheck {
				return true
			}
		}
	}
	return false
}

func part1(numbers []int) int {
	lookBehind := 25
	i := lookBehind

	for _, num := range numbers[lookBehind:] {
		if !isSumOfPrevious(num, numbers[i-lookBehind:i]) {
			return num
		}
		i++
	}
	return 0
}

func part2(numbers []int, numFromPart1 int) (int, int) {
OUTER_LOOP:
	for i, num := range numbers {
		sumSoFar := num
		for _, num2 := range numbers[i+1:] {
			sumSoFar += num2
			if sumSoFar == numFromPart1 {
				return num, num2
			} else if sumSoFar > numFromPart1 {
				continue OUTER_LOOP
			}
		}
	}
	return 0, 0
}

func main() {
	data, err := ioutil.ReadFile("day9.in")
	checkError(err)

	numbers := parseNumbers(string(data))

	// Part1
	numFromPart1 := part1(numbers)
	fmt.Println("PART1:", numFromPart1)

	// Part 2
	num1, num2 := part2(numbers, numFromPart1)
	fmt.Println("PART2:", num1+num2)
}
