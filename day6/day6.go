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

type familyForm struct {
	numPeople int
	form      map[rune]int
}

func getKeys(amap map[rune]int) []rune {
	keys := make([]rune, len(amap))

	i := 0
	for k := range amap {
		keys[i] = k
		i++
	}
	return keys
}

func getForms(formStr string) familyForm {
	form := map[rune]int{}

	answers := strings.Split(formStr, "\n")
	for _, a := range answers {
		for _, c := range a {
			if _, ok := form[c]; ok {
				form[c]++
			} else {
				form[c] = 1
			}
		}
	}

	return familyForm{
		form:      form,
		numPeople: len(answers),
	}
}

func main() {
	data, err := ioutil.ReadFile("day6.in")
	checkError(err)

	declarationForms := strings.Split(string(data), "\n\n")

	familyForms := make([]familyForm, len(declarationForms))

	for i, declarationForm := range declarationForms {
		familyForms[i] = getForms(declarationForm)
	}

	// Part 1
	sumPart1 := 0

	for _, form := range familyForms {
		sumPart1 += len(getKeys(form.form))
	}

	fmt.Println("PART1:", sumPart1)

	// Part 2
	sumPart2 := 0

	for _, form := range familyForms {
		for _, v := range form.form {
			if v == form.numPeople {
				sumPart2++
			}
		}
	}

	fmt.Println("PART2:", sumPart2)
}
