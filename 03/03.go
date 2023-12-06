package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	solve(input)
}

func solve(input string) {
	var schematic [][]string
	var partNumbers []int
	gears := make(map[string][]int)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	// Build schematic map
	for _, line := range lines {
		line := strings.Split(line, "")
		schematic = append(schematic, line)
	}

	for y, row := range schematic {
		for x := 0; x < len(row); x++ {
			element := row[x]

			if isNumber(element) {
				var wholeNumber []string

				// Find whole number
				i := 0
				for isNumber(schematic[y][x+i]) {
					wholeNumber = append(wholeNumber, schematic[y][x+i])
					i = i + 1

					if x+i == len(row) {
						break
					}
				}

				// Scan around the number and see if there is a symbol
				for xs := x - 1; xs <= (x + len(wholeNumber)); xs++ {
					// Check if scanner is out of bounds on x-axis
					if xs < 0 || xs >= len(row) {
						continue
					}
					for ys := y - 1; ys <= (y + 1); ys++ {
						// Don't scan the number itself
						if (xs >= x && xs < x+len(wholeNumber)) && ys == y {
							continue
						}

						// Check if scanner is out of bounds on y-axis
						if ys < 0 || ys >= len(schematic) {
							continue
						}

						scanner := schematic[ys][xs]

						r := regexp.MustCompile(`([^\d\s.]+)`)
						m := r.FindAllString(scanner, -1)
						// Found a symbol
						if m != nil {
							n, _ := strconv.Atoi(strings.Join(wholeNumber, ""))

							if scanner == "*" {
								coords := fmt.Sprintf("%v,%v", xs, ys)
								gears[coords] = append(gears[coords], n)
								partNumbers = append(partNumbers, n)
							} else {
								partNumbers = append(partNumbers, n)
							}
						}
					}
				}

				// Skip scanner to after the whole number
				x = x + len(wholeNumber)
			}
		}
	}

	// Add up all part numbers
	total := 0
	for _, v := range partNumbers {
		total += v
	}
	fmt.Printf("Part 1: %v\n", total)

	// Multiply gear ratios
	sum := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			ratio := gear[0] * gear[1]
			sum += ratio
		}
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return (err == nil)
}
