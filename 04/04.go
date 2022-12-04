package four

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne(input *os.File) string {
	scanner := bufio.NewScanner(input)

	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		first := strings.Split(line[0], "-")
		second := strings.Split(line[1], "-")

		firstStart, err := strconv.Atoi(first[0])
		firstEnd, err := strconv.Atoi(first[1])
		secondStart, err := strconv.Atoi(second[0])
		secondEnd, err := strconv.Atoi(second[1])

		if err != nil {
			continue
		}

		if firstStart >= secondStart && firstEnd <= secondEnd {
			sum += 1
		} else if secondStart >= firstStart && secondEnd <= firstEnd {
			sum += 1
		}

	}

	return strconv.Itoa(sum)
}

func PartTwo(input *os.File) string {
	scanner := bufio.NewScanner(input)

	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		first := strings.Split(line[0], "-")
		second := strings.Split(line[1], "-")

		firstStart, err := strconv.Atoi(first[0])
		firstEnd, err := strconv.Atoi(first[1])
		secondStart, err := strconv.Atoi(second[0])
		secondEnd, err := strconv.Atoi(second[1])

		if err != nil {
			continue
		}

		if secondStart <= firstEnd && secondStart >= firstStart {
			sum += 1
		} else if firstStart <= secondEnd && firstStart >= secondStart {
			sum += 1
		}
	}

	return strconv.Itoa(sum)
}
