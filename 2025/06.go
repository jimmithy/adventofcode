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
	file, err := os.Open("06.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// println("Part One: ", daySixPartOne(scanner))
	println("Part Two:", daySixPartTwo(scanner))

}

func daySixPartTwo(scanner *bufio.Scanner) int {
	total := 0
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	mathProblems := map[int][]int{}
	operators := lines[len(lines)-1]
	currentOperator := ""
	columnCount := 0

	for i := 0; i <= len(operators); i++ {
		fullNumber := ""
		if i < len(operators) {
			firstNum := string(lines[0][i])
			secondNum := string(lines[1][i])
			thirdNum := string(lines[2][i])
			fourthNum := string(lines[3][i])
			operator := string(operators[i])

			if operator != " " {
				currentOperator = operator
			}

			fullNumber = strings.TrimSpace(firstNum + secondNum + thirdNum + fourthNum)
		}

		if fullNumber == "" {
			// Finished collection.
			// Calculate the result based on the operator.
			digits := len(mathProblems[columnCount])
			sum := 0
			for j := 0; j < digits; j++ {
				if currentOperator == "+" {
					sum += mathProblems[columnCount][j]
				} else {
					if sum <= 0 {
						sum = 1
					}
					sum *= mathProblems[columnCount][j]
				}
			}

			fmt.Printf("%v %v %v\n", mathProblems[columnCount], currentOperator, sum)

			total += sum

			// Reset for next calculation
			columnCount += 1
		} else {
			// Collect numbers
			num, err := strconv.Atoi(fullNumber)

			if err == nil {
				mathProblems[columnCount] = append(mathProblems[columnCount], num)
			}
		}
	}

	return total
}

func daySixPartOne(scanner *bufio.Scanner) int {
	total := 0
	mathProblems := map[int][]int{}
	columnCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		number := ""

		for i := 0; i <= len(line); i++ {
			if i == len(line) || line[i] == ' ' {
				// Numbers are separated by spaces
				if number != "" {
					if (number == "+") || (number == "*") {
						digits := len(mathProblems[columnCount])
						sum := 0
						for j := 0; j < digits; j++ {
							if number == "+" {
								sum += mathProblems[columnCount][j]
							} else {
								if sum <= 0 {
									sum = 1
								}
								sum *= mathProblems[columnCount][j]
							}
						}

						total += sum
					} else {
						num, err := strconv.Atoi(number)

						if err == nil {
							mathProblems[columnCount] = append(mathProblems[columnCount], num)
						}
					}

					// We've now processed the number, so move to the next column and reset the current number
					number = ""
					columnCount += 1
				}
			} else {
				// Build the current number
				number += string(line[i])
			}
		}

		columnCount = 0
	}

	return total
}
