package one

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func PartOne(elves []int) string {
	// Find largest number
	max := elves[0]

	for _, v := range elves {
		if v > max {
			max = v
		}
	}

	return strconv.Itoa(max)
}

func PartTwo(elves []int) string {
	// Find the top three elves by sorting the array
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	// TODO we should confirm there are at least 3 elves.
	return strconv.Itoa(elves[0] + elves[1] + elves[2])
}

func GetElves(input *os.File) []int {
	scanner := bufio.NewScanner(input)

	elves := make([]int, 0)
	elves = append(elves, 0)

	for scanner.Scan() {
		line, lineErr := strconv.Atoi(scanner.Text())

		if lineErr != nil {
			// Empty line, so start a new elf
			elves = append(elves, 0)
		} else {
			// On the current elf, so increment total
			elves[len(elves)-1] += line
		}
	}

	return elves
}
