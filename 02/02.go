package two

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne(input *os.File) string {
	scanner := bufio.NewScanner(input)

	// Define a key value pair for the scoring
	// I don't know how to make this a constant
	scores := make(map[string]int)
	scores["A X"] = 1 + 3
	scores["A Y"] = 2 + 6
	scores["A Z"] = 3 + 0
	scores["B X"] = 1 + 0
	scores["B Y"] = 2 + 3
	scores["B Z"] = 3 + 6
	scores["C X"] = 1 + 6
	scores["C Y"] = 2 + 0
	scores["C Z"] = 3 + 3

	total := 0

	for scanner.Scan() {
		total += scores[scanner.Text()]
	}

	return strconv.Itoa(total)
}

func PartTwo(input *os.File) string {
	scanner := bufio.NewScanner(input)

	// Define a key value pair for the scoring
	scores := make(map[string]int)
	scores["A"] = 1
	scores["B"] = 2
	scores["C"] = 3

	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[1] == "X" {
			// Round must end in a loss
			// total += 0

			// Determine what should be played, and add to total
			if line[0] == "A" {
				total += scores["C"]
			} else if line[0] == "B" {
				total += scores["A"]
			} else if line[0] == "C" {
				total += scores["B"]
			}
		} else if line[1] == "Y" {
			// Round must end in a draw
			// Return total plus the value of whatever they played
			total += 3 + scores[line[0]]
		} else if line[1] == "Z" {
			// Round must end with a win
			total += 6

			// Determine what should be played, and add to total
			if line[0] == "A" {
				total += scores["B"]
			} else if line[0] == "B" {
				total += scores["C"]
			} else if line[0] == "C" {
				total += scores["A"]
			}
		}
	}

	return strconv.Itoa(total)
}
