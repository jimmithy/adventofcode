package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("07.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	println("Part One: ", daySevenPartOne(lines))
}

func daySevenPartOne(lines []string) int {

	// This is the total number of times the lines are split
	total := 0

	// First line, find the start index
	firstLine := lines[0]
	startIndex := 0

	for i := 0; i < len(firstLine); i++ {
		if firstLine[i] == 'S' {
			// Found our start index
			startIndex = i
			break
		}
	}

	println("Start index is", startIndex)

	// Using the start index, work our way through each line until we hit a "^" character.
	beams := make([]int, 0)
	beams = append(beams, startIndex)

	for _, line := range lines {
		foundBeams := make([]int, 0)
		for _, beam := range beams {
			if line[beam] == '^' {
				total += 1
				foundBeams = append(foundBeams, beam)
			}
		}

		if len(foundBeams) > 0 {
			for _, beam := range foundBeams {
				beams = slices.DeleteFunc(beams, func(value int) bool {
					return value == beam
				})
				beams = addUnique(beams, beam-1)
				beams = addUnique(beams, beam+1)
			}

			foundBeams = make([]int, 0)
		}
	}

	return total
}

func addUnique(beams []int, value int) []int {
	for _, beam := range beams {
		if beam == value {
			return beams
		}
	}
	return append(beams, value)
}
