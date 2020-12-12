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

func countAdjacentOccupiedSeats(seats [][]byte, x int, y int) int {
	rowLen := len(seats)
	colLen := len(seats[0])
	adjacentSeats := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || i >= rowLen {
				continue
			}
			if j < 0 || j >= colLen {
				continue
			}
			seat := seats[i][j]
			if seat == '#' {
				adjacentSeats++
			}
		}
	}
	return adjacentSeats
}

func toggleSeat(seatsToCheck [][]byte, seatsToChange [][]byte, x int, y int) {
	// If seat is empty (L) and no adjacent occupied seats, seat becomes occupied (#)
	// If seat is occupied (#) and 4 or more adjacent seats also occupied, seat becomes empty (L)
	// Otherwise, seat's state does not change

	seat := seatsToCheck[x][y]
	adjacentSeats := countAdjacentOccupiedSeats(seatsToCheck, x, y)

	if seat == 'L' && adjacentSeats == 0 {
		seatsToChange[x][y] = '#'
	} else if seat == '#' && adjacentSeats >= 4 {
		seatsToChange[x][y] = 'L'
	}
}

func fillSeats(seatsToCheck [][]byte, seatsToChange [][]byte) {
	for x, row := range seatsToCheck {
		for y := range row {
			toggleSeat(seatsToCheck, seatsToChange, x, y)
		}
	}
}

func countOccupiedSeats(seats [][]byte) int {
	numOccupied := 0
	for _, row := range seats {
		for _, col := range row {
			if col == '#' {
				numOccupied++
			}
		}
	}
	return numOccupied
}

func isEqual(seats1 [][]byte, seats2 [][]byte) bool {
	for x, row := range seats1 {
		for y := range row {
			if seats1[x][y] != seats2[x][y] {
				return false
			}
		}
	}
	return true
}

func copySeats(seats [][]byte) [][]byte {
	newSeats := [][]byte{}

	for _, row := range seats {
		newRow := []byte{}
		for _, col := range row {
			newRow = append(newRow, col)
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats
}

func part1(seats [][]byte) int {
	seats2 := copySeats(seats)

	i := 0
	for {
		i++
		fmt.Println(i)

		seats1 := copySeats(seats2)
		fillSeats(seats1, seats2)

		if isEqual(seats1, seats2) {
			return countOccupiedSeats(seats1)
		}
	}
	return 0
}

func printSeats(seats [][]byte) {
	for _, row := range seats {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func parseSeats(data string) [][]byte {
	seats := strings.Split(data, "\n")

	byteSeats := [][]byte{}

	for _, row := range seats {
		byteSeats = append(byteSeats, []byte(row))
	}
	return byteSeats
}

func main() {
	data, err := ioutil.ReadFile("day11.in")
	checkError(err)

	seats := parseSeats(string(data))

	fmt.Println("PART1:", part1(seats))
}
