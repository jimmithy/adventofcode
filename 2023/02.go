package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	total := 0
	gameNumber := 1
	maxColors := map[string]int{"blue": 14, "red": 12, "green": 13}

	for scanner.Scan() {
		line := scanner.Text()
		games := strings.Split(strings.Split(line, ":")[1], ";")
		gameIsValid := true

		// Parse Game
		for _, game := range games {
			cubes := strings.Split(game, ",")

			for i, cube := range cubes {
				// Get count
				num, err := strconv.Atoi(strings.Split(strings.TrimSpace(cube), " ")[0])

				if err == nil {
					// For each colour, check if the number is larger than the max
					for color, max := range maxColors {
						if strings.Contains(cube, color) {
							if num > max {
								gameIsValid = false
							}
						}
					}
				} else {
					fmt.Println("Error parsing cube: " + strconv.Itoa(gameNumber) + " - " + strconv.Itoa(i))
				}
			}
		}

		// Validate Cubes
		if gameIsValid {
			total = total + gameNumber
		}

		// Get Ready for next game
		gameNumber = gameNumber + 1
	}

	fmt.Println("Total: " + strconv.Itoa(total))
}

func partTwo(scanner *bufio.Scanner) {
	total := 0
	gameNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		games := strings.Split(strings.Split(line, ":")[1], ";")
		fewestCubes := map[string]int{"blue": 0, "red": 0, "green": 0}

		// Parse Game
		for _, game := range games {
			cubes := strings.Split(game, ",")

			for i, cube := range cubes {
				// Get count
				num, err := strconv.Atoi(strings.Split(strings.TrimSpace(cube), " ")[0])

				if err == nil {
					for color, fewest := range fewestCubes {
						if strings.Contains(cube, color) {
							if num > fewest {
								fewestCubes[color] = num
							}
						}
					}
				} else {
					fmt.Println("Error parsing cube: " + strconv.Itoa(gameNumber) + " - " + strconv.Itoa(i))
				}
			}
		}

		// Validate Cubes
		power := fewestCubes["red"] * fewestCubes["green"] * fewestCubes["blue"]
		total = total + power

		// Get Ready for next game
		gameNumber = gameNumber + 1
	}

	fmt.Println("Total: " + strconv.Itoa(total))
}
