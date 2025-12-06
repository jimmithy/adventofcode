package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("06.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	println("Part One: ", total)
}
