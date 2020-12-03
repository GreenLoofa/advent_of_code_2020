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

func main() {
	data, err := ioutil.ReadFile("day2.in")
	checkError(err)

	passwords := strings.Split(string(data), "\n")

	validPasswordsPart1 := 0

	for _, password := range passwords {
		s := strings.Fields(password)

		minMax := strings.Split(s[0], "-")

		charToCheck := strings.Split(s[1], ":")[0]
		passwordToCheck := strings.TrimLeft(s[2], " ")

		min, err := strconv.Atoi(minMax[0])
		checkError(err)
		max, err := strconv.Atoi(minMax[1])
		checkError(err)

		charCount := strings.Count(passwordToCheck, charToCheck)

		if charCount >= min && charCount <= max {
			validPasswordsPart1++
		}
	}
	fmt.Println("PART 1:", validPasswordsPart1)

	validPasswordsPart2 := 0

	for _, password := range passwords {
		s := strings.Fields(password)
		indices := strings.Split(s[0], "-")

		charToCheck := strings.Split(s[1], ":")[0]
		passwordToCheck := strings.TrimLeft(s[2], " ")

		index1, err := strconv.Atoi(indices[0])
		checkError(err)
		index2, err := strconv.Atoi(indices[1])
		checkError(err)

		atFirstIndex := (passwordToCheck[index1-1:index1] == charToCheck)
		atSecondIndex := (passwordToCheck[index2-1:index2] == charToCheck)

		if atFirstIndex && atSecondIndex {
			continue
		}
		if atFirstIndex || atSecondIndex {
			validPasswordsPart2++
		}
	}
	fmt.Println("PART 2:", validPasswordsPart2)
}
