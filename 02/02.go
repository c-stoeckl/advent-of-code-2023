package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", partOne(input))
	fmt.Printf("Part 2: %v\n", partTwo(input))
}

func partOne(input string) int {
	var possibleGames []int
	lines := strings.Split(input, "\n")
	expectedRed := 12
	expectedGreen := 13
	expectedBlue := 14

	for _, game := range lines {
		gamePossible := true

		if game == "" {
			continue
		}

		game := strings.Split(game, ":")
		gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		takes := strings.Split(game[1], ";")

		for _, take := range takes {
			colors := strings.Split(take, ",")
			gotRed := 0
			gotGreen := 0
			gotBlue := 0

			for _, color := range colors {
				slice := strings.Split(color, " ")
				cubeCount, _ := strconv.Atoi(slice[1])
				color := slice[2]

				if color == "red" {
					gotRed += cubeCount
				} else if color == "green" {
					gotGreen += cubeCount
				} else if color == "blue" {
					gotBlue += cubeCount
				}
			}

			if gotRed > expectedRed || gotBlue > expectedBlue || gotGreen > expectedGreen {
				gamePossible = false
			}
		}

		if gamePossible {
			possibleGames = append(possibleGames, gameId)
		}

	}

	result := 0
	for _, v := range possibleGames {
		result += v
	}
	return result
}

func partTwo(input string) int {
	var powers []int
	lines := strings.Split(input, "\n")

	for _, game := range lines {
		minRed := 0
		minGreen := 0
		minBlue := 0

		if game == "" {
			continue
		}

		game := strings.Split(game, ":")
		takes := strings.Split(game[1], ";")

		for _, take := range takes {
			colors := strings.Split(take, ",")
			gotRed := 0
			gotGreen := 0
			gotBlue := 0

			for _, color := range colors {
				slice := strings.Split(color, " ")
				cubeCount, _ := strconv.Atoi(slice[1])
				color := slice[2]

				if color == "red" {
					gotRed += cubeCount
				} else if color == "green" {
					gotGreen += cubeCount
				} else if color == "blue" {
					gotBlue += cubeCount
				}
			}

			if gotRed > minRed {
				minRed = gotRed
			}

			if gotGreen > minGreen {
				minGreen = gotGreen
			}

			if gotBlue > minBlue {
				minBlue = gotBlue
			}
		}

		power := minRed * minGreen * minBlue
		powers = append(powers, power)
	}

	result := 0
	for _, v := range powers {
		result += v
	}
	return result
}
