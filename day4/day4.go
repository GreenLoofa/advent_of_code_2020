package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func getKeys(amap map[string]string) []string {
	keys := make([]string, len(amap))

	i := 0
	for k := range amap {
		keys[i] = k
		i++
	}
	return keys
}

var validPassportFields = [7]string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid",
}

func validatePassportKeys(passport map[string]string) bool {
	keys := getKeys(passport)

	for _, field := range validPassportFields {
		fieldExists := false
		for _, key := range keys {
			if field == key {
				fieldExists = true
			}
		}
		if !fieldExists {
			return false
		}
	}
	return true
}

func parseToInt(s string) int {
	i, err := strconv.Atoi(s)
	checkError(err)
	return i
}

/**
 *	This would have been better to do with some validator technique like the following:
 *	https://www.golangprograms.com/go-struct-and-field-validation-examples.html
 */
func validatePassport(p map[string]string) bool {
	if !validatePassportKeys(p) {
		return false
	}

	valid := true

	if birthYearStr, ok := p["byr"]; ok {
		birthYear := parseToInt(birthYearStr)
		if birthYear < 1920 || birthYear > 2002 {
			valid = false
		}
	}

	if issueYearStr, ok := p["iyr"]; ok {
		issueYear := parseToInt(issueYearStr)
		if issueYear < 2010 || issueYear > 2020 {
			valid = false
		}
	}

	if expirationYearStr, ok := p["eyr"]; ok {
		expirationYear := parseToInt(expirationYearStr)
		if expirationYear < 2020 || expirationYear > 2030 {
			valid = false
		}
	}

	if heightStr, ok := p["hgt"]; ok {
		if strings.Contains(heightStr, "cm") {
			heightStr = strings.ReplaceAll(heightStr, "cm", "")
			height := parseToInt(heightStr)
			if height < 150 || height > 193 {
				valid = false
			}
		} else if strings.Contains(heightStr, "in") {
			heightStr = strings.ReplaceAll(heightStr, "in", "")
			height := parseToInt(heightStr)
			if height < 59 || height > 76 {
				valid = false
			}
		} else {
			valid = false
		}
	}

	if hairColourStr, ok := p["hcl"]; ok {
		re := regexp.MustCompile("^[a-f0-9]{6}$")
		if hairColourStr[0] != '#' {
			valid = false
		} else if !re.Match([]byte(hairColourStr[1:])) {
			valid = false
		}
	}

	if eyeColourStr, ok := p["ecl"]; ok {
		re := regexp.MustCompile("amb|blu|brn|gry|grn|hzl|oth")
		if !re.Match([]byte(eyeColourStr)) {
			valid = false
		}
	}

	if passportIDStr, ok := p["pid"]; ok {
		re := regexp.MustCompile("^[0-9]{9}$")
		if !re.Match([]byte(passportIDStr)) {
			valid = false
		}
	}

	return valid
}

func main() {
	data, err := ioutil.ReadFile("day4.in")
	checkError(err)

	passportsToParse := strings.Split(string(data), "\n\n")

	passportsToCheck := make([]map[string]string, len(passportsToParse))

	for _, p := range passportsToParse {
		trimmedPassport := strings.ReplaceAll(p, "\n", " ")
		passport := strings.Split(trimmedPassport, " ")

		var passportMap = make(map[string]string)

		for _, fieldValue := range passport {
			fieldValuePair := strings.Split(fieldValue, ":")

			passportMap[fieldValuePair[0]] = fieldValuePair[1]
		}
		passportsToCheck = append(passportsToCheck, passportMap)
	}

	// Part 1
	numValidPassportsPart1 := 0

	for _, passportToCheck := range passportsToCheck {
		if validatePassportKeys(passportToCheck) {
			numValidPassportsPart1++
		}
	}
	fmt.Println("PART1:", numValidPassportsPart1)

	// Part 2
	numValidPassportsPart2 := 0
	for _, passportToCheck := range passportsToCheck {
		if validatePassport(passportToCheck) {
			numValidPassportsPart2++
		}
	}
	fmt.Println("PART2:", numValidPassportsPart2)
}
