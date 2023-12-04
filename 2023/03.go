package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("03.txt")

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
	previousLine := ""

	for scanner.Scan() {
		line := scanner.Text()

		// Iterate each character
		for i, c := range line {

			// Find a Symbol
			if c != '.' && !unicode.IsDigit(c) {
				foundNumber := ""

				if i > 0 && unicode.IsDigit(rune(line[i-1])) {
					// Found a digit to the left of this symbol
					foundNumber = resolveNumber(line, i-1)
					fmt.Println("Found number before symbol: " + foundNumber + " " + string(c))
					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				} else if i < len(line)-1 && unicode.IsDigit(rune(line[i+1])) {
					// found a digit to the right of this symbol
					foundNumber = resolveNumber(line, i+1)

					fmt.Println("Found number after symbol: " + foundNumber + " " + string(c))
					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				} else if previousLine != "" {
					// There could be more than one number above this symbol
					foundNumbers := []string{}

					if i > 0 && unicode.IsDigit(rune(previousLine[i-1])) {
						// there is a number top left of this symbol
						foundNumber := resolveNumber(previousLine, i-1)
						foundNumbers = append(foundNumbers, foundNumber)
						fmt.Println("Found number top left of this symbol : " + foundNumber + " " + string(c))
					}

					if unicode.IsDigit(rune(previousLine[i])) {
						// There is a number above this symbol
						foundNumber := resolveNumber(previousLine, i)

						if !slices.Contains(foundNumbers, foundNumber) {
							fmt.Println("Found number above this symbol: " + foundNumber + " " + string(c))
							foundNumbers = append(foundNumbers, foundNumber)
						}
					}

					if (i < len(previousLine)-1) && unicode.IsDigit(rune(previousLine[i+1])) {
						// there is a number top right of this symbol
						foundNumber := resolveNumber(previousLine, i+1)

						if !slices.Contains(foundNumbers, foundNumber) {
							fmt.Println("Found number top right of this symbol: " + foundNumber + " " + string(c))
							foundNumbers = append(foundNumbers, foundNumber)
						}
					}
				}

				if foundNumber != "" {
					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				}
			} else if unicode.IsDigit(c) && previousLine != "" {
				// We've found a digit, check previous line for a symbol

				if previousLine[i] != '.' && !unicode.IsDigit(rune(previousLine[i])) {
					// We found a symbol above this number
					foundNumber := resolveNumber(line, i)
					fmt.Println("Found symbole above this number: " + foundNumber + " " + string(previousLine[i]))

					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				} else if i > 0 && line[i-1] == '.' && previousLine[i-1] != '.' && !unicode.IsDigit(rune(previousLine[i-1])) {
					// We found a symbol above and to the left of this number
					foundNumber := resolveNumber(line, i)
					fmt.Println("Found symbol top left of this number : " + foundNumber + " " + string(previousLine[i-1]))

					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				} else if i < len(line)-1 && line[i+1] == '.' && previousLine[i+1] != '.' && !unicode.IsDigit(rune(previousLine[i+1])) {
					// We found a symbol above and to the right of this number
					foundNumber := resolveNumber(line, i)
					fmt.Println("Found symbol top right of this number: " + foundNumber + " " + string(previousLine[i+1]))

					num, err := strconv.Atoi(foundNumber)
					if err == nil {
						total += num
					}
				}
			}
		}

		previousLine = line
	}

	fmt.Println(strconv.Itoa(total))
}

func partOneDisaster(scanner *bufio.Scanner) {
	total := 0
	previousLine := ""

	for scanner.Scan() {
		line := scanner.Text()
		currentIndex := -1
		currentNumber := ""

		for i, c := range line {
			if unicode.IsDigit(c) {
				// We found another digit, so add it to the current number or start a new one
				if currentIndex == -1 {
					currentIndex = i
				}
				currentNumber += string(c)

				// check if this is the last character in the line
				if i == len(line)-1 {
					// Check if there is a symbol before this number
					if currentIndex > 0 && line[currentIndex-1] != '.' {
						num, err := strconv.Atoi(currentNumber)
						if err == nil {
							total += num
						}
					} else if previousLine != "" {
						// Now check the previous line for a symbol
						foundSymbol := false

						// Look for a symbol on the previous line from currentIndex - 1 to currentIndex + len(currentNumber) + 1
						// If we find one, add the number to the total

						startIndex := currentIndex
						if currentIndex > 0 {
							startIndex--
						}

						fmt.Println(currentNumber + " = " + strconv.Itoa(startIndex) + ":" + strconv.Itoa(i) + " length: " + strconv.Itoa(len(previousLine)))

						for _, c := range previousLine[startIndex:i] {
							if c != '.' && !unicode.IsDigit(c) {
								foundSymbol = true
								break
							}
						}

						if foundSymbol {
							num, err := strconv.Atoi(currentNumber)
							if err == nil {
								total += num
							}
						}
					}

					// Reset
					currentIndex = -1
					currentNumber = ""
				}
			} else if c == '.' {
				// This is an ignored space, but...

				// If we have a pending number, save it
				if currentNumber != "" {
					// Check before the number for a symbol
					if currentIndex > 0 && line[currentIndex-1] != '.' {
						// We know this number is next to a symbol, so add to total
						num, err := strconv.Atoi(currentNumber)
						if err == nil {
							total += num
						}
					} else if previousLine != "" {
						// Now check the previous line for a symbol
						foundSymbol := false

						// Look for a symbol on the previous line from currentIndex - 1 to currentIndex + len(currentNumber) + 1
						// If we find one, add the number to the total

						startIndex := currentIndex
						if currentIndex > 0 {
							startIndex--
						}

						endIndex := currentIndex + len(currentNumber) - 1
						if endIndex < len(previousLine)-1 {
							endIndex++
						}

						for _, c := range previousLine[startIndex:endIndex] {
							if c != '.' && !unicode.IsDigit(c) {
								foundSymbol = true
								break
							}
						}

						if foundSymbol {
							num, err := strconv.Atoi(currentNumber)
							if err == nil {
								total += num
							}
						}
					}

					// Reset
					currentIndex = -1
					currentNumber = ""
				}
			} else {
				// We know it's a symbol

				// If we have a pending number, we know it's a part
				if currentNumber != "" {
					// We know this number is next to a symbol, so add to total
					num, err := strconv.Atoi(currentNumber)
					if err == nil {
						total += num
					}

					// Reset
					currentIndex = -1
					currentNumber = ""
				} else if previousLine != "" {
					// Check the previous line for a number
					foundNumbers := []string{}

					if unicode.IsDigit(rune(previousLine[i])) {
						// Character below a number
						foundNumber := resolveNumber(previousLine, i)
						foundNumbers = append(foundNumbers, foundNumber)
					}

					if i > 0 && unicode.IsDigit(rune(previousLine[i-1])) {
						// Charater is diagonal to a number
						foundNumber := resolveNumber(previousLine, i-1)
						// Make sure we don't already have this number
						if !slices.Contains(foundNumbers, foundNumber) {
							foundNumbers = append(foundNumbers, foundNumber)
						}
					}

					if (i < len(previousLine)-1) && unicode.IsDigit(rune(previousLine[i+1])) {
						// Charater is diagonal to a number
						foundNumber := resolveNumber(previousLine, i+1)
						// Make sure we don't already have this number
						if !slices.Contains(foundNumbers, foundNumber) {
							foundNumbers = append(foundNumbers, foundNumber)
						}
					}

					for _, foundNumber := range foundNumbers {
						num, err := strconv.Atoi(foundNumber)
						if err == nil {
							total += num
						}
					}
				}
			}
		}

		previousLine = line
	}

	fmt.Println(strconv.Itoa(total))
}

func resolveNumber(line string, index int) string {
	// find start of number
	startIndex := index
	for {
		if startIndex == 0 || !unicode.IsDigit(rune(line[startIndex-1])) {
			break
		} else {
			startIndex = startIndex - 1
		}
	}

	// Move forward until we find something that is not a digit or we get to the end of the line
	number := ""
	for _, c := range line[startIndex:] {
		if unicode.IsDigit(c) {
			number += string(c)
		} else {
			break
		}
	}

	return number
}
