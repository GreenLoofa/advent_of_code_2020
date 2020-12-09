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

type instruction struct {
	op  string
	id  int
	val int
}

func parseInstructions(data string) []instruction {
	instructionsStr := strings.Split(data, "\n")

	instructions := make([]instruction, len(instructionsStr))

	for i, instStr := range instructionsStr {
		inst := strings.Split(instStr, " ")

		val, err := strconv.Atoi(inst[1])
		checkError(err)

		instructions[i] = instruction{
			op:  inst[0],
			id:  i,
			val: val,
		}
	}
	return instructions
}

func run(instructions []instruction, iterations int) (int, bool) {
	tracker := map[int]int{}

	acc := 0
	addrPointer := 0

	lenInstructions := len(instructions)

	terminated := false

	for {
		if addrPointer >= lenInstructions {
			terminated = true
			break
		}

		currInstruction := instructions[addrPointer]

		if _, ok := tracker[currInstruction.id]; ok {
			tracker[currInstruction.id]++
			if tracker[currInstruction.id] > iterations {
				break
			}
		} else {
			tracker[currInstruction.id] = 1
		}

		switch currInstruction.op {
		case "nop":
			addrPointer++
		case "acc":
			acc += currInstruction.val
			addrPointer++
		case "jmp":
			addrPointer += currInstruction.val
		}
	}

	return acc, terminated
}

func part1(instructions []instruction) {
	acc, _ := run(instructions, 1)
	fmt.Println("PART1:", acc)
}

func part2(instructions []instruction) {
	for i, inst := range instructions {
		isNop := (inst.op == "nop")
		isJmp := (inst.op == "jmp")

		if !isNop && !isJmp {
			continue
		}

		if isNop {
			instructions[i].op = "jmp"
		}

		if isJmp {
			instructions[i].op = "nop"
		}

		if val, ok := run(instructions, 1); ok {
			fmt.Println("PART2:", val)
			return
		}

		if isNop {
			instructions[i].op = "nop"
		}

		if isJmp {
			instructions[i].op = "jmp"
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("day8.in")
	checkError(err)

	instructions := parseInstructions(string(data))

	// Part1
	part1(instructions)

	// Part 2
	part2(instructions)
}
