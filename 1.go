package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/1.txt
var input string

var letterDigits = [][]string{{"one", "1"}, {"two", "2"}, {"three", "3"}, {"four", "4"}, {"five", "5"}, {"six", "6"}, {"seven", "7"}, {"eight", "8"}, {"nine", "9"}}

func main() {
	fmt.Printf("Part 1: %v\n", partOne(input))
	fmt.Printf("Part 2: %v\n", partTwo(input))
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

	var calibrationValues []int

	for _, line := range lines {
		var digits []int

		// Skip empty line
		if line == "" {
			continue
		}

		// fmt.Printf("Original Input: %v\n", line)
		// Get digits from string
		for _, char := range line {
			if isNumber(string(char)) {
				n, err := strconv.Atoi(string(char))

				if err != nil {
					panic(err)
				}
				digits = append(digits, n)
			}
		}
		// fmt.Printf("Digits: %v\n", digits)

		if len(digits) == 0 {
			continue
		}

		first := digits[0]
		last := digits[len(digits)-1]

		calibrationValue, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		// fmt.Printf("Calibration Value: %v\n", calibrationValue)
		calibrationValues = append(calibrationValues, calibrationValue)
	}

	// Add up all calibration values
	total := 0
	for _, value := range calibrationValues {
		total += value
	}

	return total
}

func partTwo(input string) int {
	var calibrationValues []int
	lines := strings.Split(input, "\n")

	// Loop through every line
	for _, line := range lines {
		var digits []int

		r := regexp.MustCompile(`(\d|[^\d]+)`)
		slicedLine := r.FindAllString(line, -1)

		fmt.Printf("Line: %v\n", line)

		// Skip empty line
		if line == "" {
			continue
		}

		fmt.Printf("Sliced line: %v\n", slicedLine)

		for _, ls := range slicedLine {
			if isNumber(ls) {
				n, _ := strconv.Atoi(ls)
				digits = append(digits, n)
			} else {

				var tempDigits [][]int
				for _, v := range letterDigits {
					vi := strings.Index(ls, v[0])

					if vi != -1 {
						t, _ := strconv.Atoi(v[1])
						v := []int{vi, t}

						tempDigits = append(tempDigits, v)
					}
				}

				sort.Slice(tempDigits, func(i, j int) bool {
					return tempDigits[i][0] < tempDigits[j][0]
				})

				for _, v := range tempDigits {
					digits = append(digits, v[1])
				}

				// fmt.Println(tempDigits)
			}
		}

		fmt.Printf("Digits: %v\n", digits)

		first := digits[0]
		last := digits[len(digits)-1]

		calibrationValue, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		fmt.Printf("Calibration Value: %v\n", calibrationValue)
		calibrationValues = append(calibrationValues, calibrationValue)

	}

	// Add up all calibration values
	total := 0
	for _, value := range calibrationValues {
		total += value
	}

	return total
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return (err == nil)
}
