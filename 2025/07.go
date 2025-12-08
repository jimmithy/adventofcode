package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
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

	// This is the totalSplits number of times the lines are split
	totalSplits := 0

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
				totalSplits += 1
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

	println("Part One", totalSplits)

	trimmedLines := []string{}

	for _, line := range lines {
		if strings.Contains(line, "^") {
			fmt.Printf("%v\n", line)
			trimmedLines = append(trimmedLines, line)
		}
	}

	// Part Two: Count unique paths to the bottom
	totalLength := len(trimmedLines)
	visited := make(map[int]int)
	possiblePaths := search(trimmedLines, 0, totalLength, startIndex, visited)

	println("Part Two", possiblePaths)
}

func addUnique(beams []int, value int) []int {
	for _, beam := range beams {
		if beam == value {
			return beams
		}
	}
	return append(beams, value)
}

func search(lines []string, lineIndex int, totalLength int, beamIndex int, visited map[int]int) int {
	if lineIndex >= totalLength {
		return 1
	}

	// Have we been here before?
	key := lineIndex*1000 + beamIndex
	if visited[key] != 0 {
		return visited[key]
	}

	currentLine := lines[lineIndex]
	result := 0

	if currentLine[beamIndex] == '^' {
		// left path
		left := search(lines, lineIndex+1, totalLength, beamIndex-1, visited)
		// right path
		right := search(lines, lineIndex+1, totalLength, beamIndex+1, visited)
		result = left + right
	} else if currentLine[beamIndex] == '.' {
		// Next Level down
		result = search(lines, lineIndex+1, totalLength, beamIndex, visited)
	}

	visited[key] = result
	return result
}
