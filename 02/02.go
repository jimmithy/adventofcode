package two

import (
	"bufio"
	"os"
	"strconv"
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
