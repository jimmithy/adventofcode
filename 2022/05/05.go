package five

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input *os.File) string {
	scanner := bufio.NewScanner(input)

	stacks := make(map[int][]string)
	numberRegex := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "move") {
			// Handle movement
			numbers := numberRegex.FindAllString(line, -1)
			movements, _ := strconv.Atoi(numbers[0])
			from, _ := strconv.Atoi(numbers[1])
			to, _ := strconv.Atoi(numbers[2])

			for i := 0; i < movements && len(stacks[from]) > 0; i++ {
				crate := stacks[from][0]
				stacks[from] = stacks[from][1:]
				stacks[to] = append([]string{crate}, stacks[to]...)
			}
		} else if len(line) > 0 {
			// Build stacks. FML.
			max := len(line)
			count := (max + 1) / 9

			for i := 0; i <= max; i += count {
				entry := strings.Trim(line[i:i+count-1], " ")

				// Add entries if they are not a number
				if entry != "" && numberRegex.Match([]byte(entry)) == false {
					stackIndex := i/count + 1
					stacks[stackIndex] = append(stacks[stackIndex], entry)
				}
			}
		}
	}

	output := ""

	for i := 0; i < len(stacks); i++ {
		output = output + stacks[i+1][0]
	}

	return output
}

func PartTwo(input *os.File) string {
	scanner := bufio.NewScanner(input)

	stacks := make(map[int][]string)
	numberRegex := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "move") {
			// Handle movement
			numbers := numberRegex.FindAllString(line, -1)
			movements, _ := strconv.Atoi(numbers[0])
			from, _ := strconv.Atoi(numbers[1])
			to, _ := strconv.Atoi(numbers[2])

			// TIL I have to specifically make a copy of the slice
			crates := make([]string, movements)
			copy(crates, stacks[from][:movements])

			stacks[from] = stacks[from][movements:]
			stacks[to] = append(crates, stacks[to]...)

		} else if len(line) > 0 {
			// Build stacks. FML.
			max := len(line)
			count := (max + 1) / 9

			for i := 0; i <= max; i += count {
				entry := strings.Trim(line[i:i+count-1], " ")

				// Add entries if they are not a number
				if entry != "" && numberRegex.Match([]byte(entry)) == false {
					stackIndex := i/count + 1
					stacks[stackIndex] = append(stacks[stackIndex], entry)
				}
			}
		}
	}

	output := ""

	for i := 1; i <= len(stacks); i++ {
		output = output + stacks[i][0]
	}

	return output
}
