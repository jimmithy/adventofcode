package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	dayThreePartTwo(scanner)
}

func dayThree(scanner *bufio.Scanner) {
	partOneTotal := 0

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		largestNumber := 0
		largestNumberIndex := 0

		// Within each line, first find the largest number
		for i := 9; i > 0; i-- {
			firstNumber := strconv.Itoa(i)

			if strings.Contains(line, firstNumber) {
				largestNumber = i
				largestNumberIndex = strings.Index(line, firstNumber)

				// Can't be the final number in the list
				if largestNumberIndex != length-1 {
					break
				}
			}
		}

		// Once we have the largest number, find the second largest number from this position
		nextLargestNumber := 0

		for i := largestNumberIndex + 1; i < len(line); i++ {
			nextNumber, _ := strconv.Atoi(string(line[i]))

			if nextNumber > nextLargestNumber {
				nextLargestNumber = nextNumber
			}
		}

		// Calculate the total for part one
		partOneTotal += (largestNumber * 10) + nextLargestNumber
	}

	println("Part 1: ", partOneTotal)
}

/**
* Find the largest twelve digit number from each line
 */
func dayThreePartTwo(scanner *bufio.Scanner) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		largestNumber := []string{}
		startingIndex := 0

		// Within each line, find the largest twelve digit number
		for j := 0; j < 12; j++ {
			nextLargestNumber := 0

			for i := startingIndex; i < length; i++ {
				nextNumber, _ := strconv.Atoi(string(line[i]))

				if nextNumber > nextLargestNumber {
					nextLargestNumber = nextNumber
					startingIndex = i + 1
				}

				// If there aren't enough digits left to complete the twelve digit number, break
				if (length - 1 - i) < (12 - j) {
					break
				}
			}

			largestNumber = append(largestNumber, strconv.Itoa(nextLargestNumber))
		}

		// merge string, convert to a number, then add it to the total
		result := strings.Join(largestNumber, "")
		num, _ := strconv.Atoi(result)
		total += num
	}

	println("Part 2: ", total)
}
