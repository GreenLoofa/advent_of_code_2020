package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day1.in")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	splits := strings.Split(string(data), "\n")

	splits_length := len(splits)

PART1:
	for i, element := range splits {
		num, err := strconv.Atoi(element)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		for a := i; a < splits_length; a++ {
			num2, err := strconv.Atoi(splits[a])
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			if (num + num2) == 2020 {
				fmt.Println("PART 1:", num*num2)
				break PART1
			}
		}
	}

	for i, element := range splits {
		num, err := strconv.Atoi(element)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		for a := i; a < splits_length; a++ {
			num2, err := strconv.Atoi(splits[a])
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			for b := a; b < splits_length; b++ {
				num3, err := strconv.Atoi(splits[b])
				if err != nil {
					// handle error
					fmt.Println(err)
					os.Exit(2)
				}
				if (num + num2 + num3) == 2020 {
					fmt.Println("PART 2:", num*num2*num3)
					return
				}
			}
		}
	}
}
