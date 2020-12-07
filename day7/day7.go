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

func containsBag(bagsGraph map[string][]string, currKey string, keyToFind string) bool {
	if currKey == keyToFind {
		return true
	}

	if bagsGraph[currKey][0] == "no other bags" {
		return false
	}

	doesContainBag := false

	for _, bag := range bagsGraph[currKey] {
		if containsBag(bagsGraph, bag, keyToFind) {
			doesContainBag = true
			break
		}
	}

	return doesContainBag
}

func part1(bagsStr []string) {
	// Part 1
	bagsGraph := map[string][]string{}

	for _, bagStr := range bagsStr {
		bagInfo := strings.Split(bagStr, " bags contain ")
		for _, containedBagsStr := range strings.Split(bagInfo[1], ", ") {
			if strings.Contains(containedBagsStr, "no other bags") {
				bagsGraph[bagInfo[0]] = append(bagsGraph[bagInfo[0]], "no other bags")
				continue
			}
			containedBagSplit := strings.Split(containedBagsStr, " ")
			bagsGraph[bagInfo[0]] = append(bagsGraph[bagInfo[0]], strings.Join(containedBagSplit[1:3], " "))
		}
	}

	numBagsPart1 := 0

	for k := range bagsGraph {
		if containsBag(bagsGraph, k, "shiny gold") {
			numBagsPart1++
		}
	}

	fmt.Println("PART1:", numBagsPart1-1)
}

func countBags(bagsGraph map[string]map[string]int, currKey string) int {
	if _, ok := bagsGraph[currKey]["no other bags"]; ok {
		return 0
	}

	sum := 0

	for k, v := range bagsGraph[currKey] {
		sum += (v * countBags(bagsGraph, k))
		sum += v
	}

	return sum
}

func part2(bagsStr []string) {
	bagsGraph := map[string]map[string]int{}

	for _, bagStr := range bagsStr {
		bagInfo := strings.Split(bagStr, " bags contain ")
		outerBagName := bagInfo[0]
		for _, containedBagsStr := range strings.Split(bagInfo[1], ", ") {
			if bagsGraph[outerBagName] == nil {
				bagsGraph[outerBagName] = map[string]int{}
			}

			if strings.Contains(containedBagsStr, "no other bags") {
				bagsGraph[outerBagName]["no other bags"] = 0
				continue
			}
			containedBagSplit := strings.Split(containedBagsStr, " ")

			numBags, err := strconv.Atoi(containedBagSplit[0])
			checkError(err)

			bagName := strings.Join(containedBagSplit[1:3], " ")

			bagsGraph[outerBagName][bagName] = numBags
		}
	}

	fmt.Println("PART2:", countBags(bagsGraph, "shiny gold"))
}

func main() {
	data, err := ioutil.ReadFile("day7.in")
	checkError(err)

	bagsStr := strings.Split(string(data), "\n")

	// Part1
	part1(bagsStr)

	// Part 2
	part2(bagsStr)
}
