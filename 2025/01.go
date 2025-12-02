package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	dial := 50
	total := 0
	passesZero := 0

	for scanner.Scan() {
		line := scanner.Text()

		clicks, clicksErr := strconv.Atoi(line[1:])

		if clicksErr == nil {
			// If it's over 100, reduce it because the dial only has 100 positions
			if clicks >= 100 {
				passesZero = passesZero + (clicks / 100)
				clicks = clicks % 100
			}

			// Rotate the dial
			switch line[0:1] {
			case "R":
				dial = dial + clicks
			case "L":
				dial = dial - clicks
			}

			// Validate the dial position
			if dial > 100 {
				dial = dial - 100
				passesZero = passesZero + 1
			} else if dial < 0 {
				dial = dial + 100
				passesZero = passesZero + 1
			} else if dial == 100 {
				dial = 0
			}

			// If we finish on 0, increase total
			if dial == 0 {
				passesZero = passesZero + 1
				total = total + 1
			}
		}
	}

	println("Part 1: ", total)
	println("Part 2: ", passesZero)
}

/**
* For some reason my approach to get the first answer isn't working for the second part.
* So I'll just brute force it by iterating each click one by one.
 */
func partTwo(scanner *bufio.Scanner) {
	dial := 50
	total := 0
	passesZero := 0
	direction := 1

	for scanner.Scan() {
		line := scanner.Text()

		clicks, clicksErr := strconv.Atoi(line[1:])

		if clicksErr == nil {
			// Rotate the dial
			switch line[0:1] {
			case "R":
				direction = 1
			case "L":
				direction = -1
			}

			// Loop through clicks one by one
			for i := 0; i < clicks; i++ {
				dial = dial + direction

				if dial > 99 {
					dial = 0
				} else if dial < 0 {
					dial = 99
				}

				// dial is at zero, so count it
				if dial == 0 {
					passesZero = passesZero + 1
				}
			}

			// If we finish on 0, increase total
			if dial == 0 {
				total = total + 1
			}
		}
	}

	println("Part 1: ", total)
	println("Part 2:", passesZero)
}
