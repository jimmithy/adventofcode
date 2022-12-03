package three

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PartOne(input *os.File) string {
	scanner := bufio.NewScanner(input)

	priorities := "abcdefghijklmnopqrstuvwxyz"
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		left := line[:half]
		right := line[half:]

		for _, character := range left {
			// Does the character exist in the right compartment
			if strings.IndexRune(right, character) != -1 {
				// use the position to get the priority
				position := strings.IndexRune(priorities, unicode.ToLower(character)) + 1

				// if upper case, add the total number of characters
				if unicode.IsUpper(character) {
					position += len(priorities)
				}

				sum += position
				break
			}
		}
	}

	return strconv.Itoa(sum)
}

func PartTwo(input *os.File) string {
	scanner := bufio.NewScanner(input)

	priorities := "abcdefghijklmnopqrstuvwxyz"
	sum := 0

	first := ""
	second := ""

	for scanner.Scan() {
		line := scanner.Text()

		if first == "" {
			first = line
		} else if second == "" {
			second = line
		} else {
			for _, character := range line {
				// Does the character exist in the right compartment
				if strings.IndexRune(first, character) != -1 && strings.IndexRune(second, character) != -1 {
					// use the position to get the priority
					position := strings.IndexRune(priorities, unicode.ToLower(character)) + 1

					// if upper case, add the total number of characters
					if unicode.IsUpper(character) {
						position += len(priorities)
					}

					sum += position

					// Reset
					first = ""
					second = ""
					break
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
