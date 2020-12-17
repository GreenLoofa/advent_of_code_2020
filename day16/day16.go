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

type interval struct {
	min int
	max int
}

func parseTicket(ticketStr string) []int {
	ticket := []int{}
	for _, numStr := range strings.Split(ticketStr, ",") {
		num, err := strconv.Atoi(numStr)
		checkError(err)

		ticket = append(ticket, num)
	}
	return ticket
}

func parseInterval(intervalStr string) interval {
	vals := strings.Split(intervalStr, "-")

	min, err := strconv.Atoi(vals[0])
	checkError(err)

	max, err := strconv.Atoi(vals[1])
	checkError(err)

	return interval{
		min: min,
		max: max,
	}
}

func parseIntervals(intervalsStr string) map[string][]interval {
	retIntervals := map[string][]interval{}

	for _, nameValue := range strings.Split(intervalsStr, "\n") {
		nameIntervalsSplit := strings.Split(nameValue, ": ")
		name := nameIntervalsSplit[0]

		intervalsSplit := strings.Split(nameIntervalsSplit[1], " or ")

		intervals := []interval{}
		intervals = append(intervals, parseInterval(intervalsSplit[0]))
		intervals = append(intervals, parseInterval(intervalsSplit[1]))

		retIntervals[name] = intervals
	}
	return retIntervals
}

func parseInput(data string) (map[string][]interval, []int, [][]int) {
	dataSplit := strings.Split(data, "\n\n")

	intervalsStr := dataSplit[0]
	yourTicketStr := dataSplit[1]
	nearbyTicketsStr := dataSplit[2]

	intervals := parseIntervals(intervalsStr)

	yourTicket := parseTicket(strings.Split(yourTicketStr, "\n")[1])

	nearbyTickets := [][]int{}
	for _, ticketStr := range strings.Split(nearbyTicketsStr, "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(ticketStr))
	}

	return intervals, yourTicket, nearbyTickets
}

func withinAnyIntervals(intervals map[string][]interval, num int) bool {
	for _, interval := range intervals {
		for _, v := range interval {
			if v.min <= num && v.max >= num {
				return true
			}
		}
	}
	return false
}

func part1(intervals map[string][]interval, nearbyTickets [][]int) {
	errorMargin := 0

	for _, ticket := range nearbyTickets {
		for _, num := range ticket {
			if !withinAnyIntervals(intervals, num) {
				errorMargin += num
			}
		}
	}
	fmt.Println("PART1:", errorMargin)
}

func filterTickets(intervals map[string][]interval, nearbyTickets [][]int) [][]int {
	retTickets := [][]int{}

	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, num := range ticket {
			if !withinAnyIntervals(intervals, num) {
				validTicket = false
			}
		}
		if validTicket {
			retTickets = append(retTickets, ticket)
		}
	}
	return retTickets
}

func getPotentials(intervals map[string][]interval, num int) []string {
	potentials := []string{}

	for name, interval := range intervals {
		inside := false
		for _, v := range interval {
			if v.min <= num && v.max >= num {
				inside = true
			}
		}
		if inside {
			potentials = append(potentials, name)
		}
	}

	return potentials
}

func union(set1 []string, set2 []string) []string {
	newSet := []string{}

OUTER_LOOP:
	for _, e1 := range set1 {
		for _, e2 := range set2 {
			if e1 == e2 {
				newSet = append(newSet, e1)
				continue OUTER_LOOP
			}
		}
	}

	return newSet
}

func disjoint(set1 []string, set2 []string) []string {
	newSet := []string{}

OUTER_LOOP:
	for _, e1 := range set1 {
		for _, e2 := range set2 {
			if e1 == e2 {
				continue OUTER_LOOP
			}
		}
		newSet = append(newSet, e1)
	}

	return newSet
}

func removeElement(s []string, elem string) []string {
	newList := []string{}

	for _, e1 := range s {
		if e1 == elem {
			continue
		}
		newList = append(newList, e1)
	}

	return newList
}

func part2(intervals map[string][]interval, yourTicket []int, nearbyTickets [][]int) {
	potentials := map[int][]string{}

	nearbyTickets = append(nearbyTickets, yourTicket)

	for _, ticket := range filterTickets(intervals, nearbyTickets) {
		for i, num := range ticket {
			newPotentials := getPotentials(intervals, num)
			if _, ok := potentials[i]; !ok {
				potentials[i] = newPotentials
			} else {
				potentials[i] = union(potentials[i], newPotentials)
			}
		}
	}

	positions := map[int]string{}
	yourTicketLen := len(yourTicket)

	for {
		if len(positions) == yourTicketLen {
			break
		}

		for k, v := range potentials {
			if len(v) == 1 {
				positions[k] = v[0]
				for k2, v2 := range potentials {
					if k2 == k || len(v2) == 1 {
						continue
					}
					potentials[k2] = removeElement(potentials[k2], v[0])
				}
			}
		}
	}

	departureVals := 1
	for k, v := range positions {
		if strings.Contains(v, "departure") {
			departureVals *= yourTicket[k]
		}
	}

	fmt.Println("PART2:", departureVals)
}

// This is by far the messiest day
func main() {
	data, err := ioutil.ReadFile("day16.in")
	checkError(err)

	intervals, yourTicket, nearbyTickets := parseInput(string(data))

	// PART1
	part1(intervals, nearbyTickets)

	// PART2
	part2(intervals, yourTicket, nearbyTickets)
}
