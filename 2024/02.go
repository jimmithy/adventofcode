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

	partOne(scanner)
	// partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		levels := []int{}

		// Convert strings to numbers
		for _, level := range line {
			number, err := strconv.Atoi(level)

			if err == nil {
				levels = append(levels, number)
			}
		}

		direction := 1
		isSafe := true

		// for each number
		for i := 1; i < len(levels); i++ {
			if isSafe == false {
				break
			}

			// Set the direction we're moving in.
			if i == 1 {
				if levels[i] > levels[i-1] {
					direction = 1
				} else {
					direction = -1
				}
			}

			if i >= 1 {
				// If the direction is not consistent, skip these levels
				if levels[i] == levels[i-1] {
					isSafe = false
				} else if direction == 1 && levels[i-1] > levels[i] {
					isSafe = false
				} else if direction == -1 && levels[i-1] < levels[i] {
					isSafe = false
				}

				// Confirm we've moved at least one and at most three levels
				if direction == 1 {
					distance := levels[i] - levels[i-1]

					if distance < 1 || distance > 3 {
						isSafe = false
					}
				} else if direction == -1 {
					distance := levels[i-1] - levels[i]

					if distance < 1 || distance > 3 {
						isSafe = false
					}
				}
			}
		}

		if isSafe {
			// We've passed all the checks, so increment the total
			total++
		}
	}

	fmt.Println(total)
}
