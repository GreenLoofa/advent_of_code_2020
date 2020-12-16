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

func applyMask(mask string, num uint64) uint64 {
	maxBits := 35

	retNum := num
	for i, bit := range mask {
		if bit == 'X' {
			continue
		}
		if bit == '1' {
			retNum |= (1 << (maxBits - i))
		} else if bit == '0' {
			retNum &= ^(1 << (maxBits - i))
		}
	}
	return retNum
}

func parseMemory(data string) map[uint64]uint64 {
	currMask := ""

	memory := map[uint64]uint64{}

	for _, line := range strings.Split(data, "\n") {
		lineSplit := strings.Split(line, " = ")

		if lineSplit[0] == "mask" {
			currMask = lineSplit[1]
		} else {
			memAddrStr := strings.Split(lineSplit[0], "[")
			memAddrStr = strings.Split(memAddrStr[1], "]")

			memAddr, err := strconv.Atoi(memAddrStr[0])
			checkError(err)

			memValueInt, err := strconv.Atoi(lineSplit[1])
			checkError(err)

			memValue := applyMask(currMask, uint64(memValueInt))

			memory[uint64(memAddr)] = uint64(memValue)
		}
	}

	return memory
}

func sumMemory(memory map[uint64]uint64) uint64 {
	var sum uint64 = 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func part1(memory map[uint64]uint64) {
	fmt.Println("PART1:", sumMemory(memory))
}

func main() {
	data, err := ioutil.ReadFile("day14.in")
	checkError(err)

	memory := parseMemory(string(data))

	// PART 1
	part1(memory)
}
