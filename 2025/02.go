package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := ""

	for scanner.Scan() {
		line = scanner.Text()
	}

	dayTwo(line)
}

func dayTwo(line string) {
	partOneTotal := 0
	partTwoTotal := 0

	productIDRange := strings.Split(line, ",")

	for _, id := range productIDRange {
		ids := strings.SplitN(id, "-", 2)

		first, firstErr := strconv.Atoi(ids[0])
		second, secondErr := strconv.Atoi(ids[1])

		if (firstErr != nil) || (secondErr != nil) {
			continue
		}

		for i := first; i <= second; i++ {
			partOneTotal += isMirror(i)
			partTwoTotal += hasRepeatedDigits(i)
		}
	}

	println("Part 1: ", partOneTotal)
	println("Part 2: ", partTwoTotal)
}

func isMirror(productId int) int {
	runes := []rune(strconv.Itoa(productId))
	length := len(runes)

	// Must be an even number of digits to be a mirror
	if length%2 != 0 {
		return 0
	}

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[(length/2)+i] {
			// Not a mirror, return zero
			return 0
		}
	}

	return productId
}

func hasRepeatedDigits(productId int) int {
	strId := strconv.Itoa(productId)
	length := len(strId)

	for i := 0; i < length/2; i++ {
		matches, err := regexp.MatchString("^("+strId[0:i+1]+")+$", strId)
		if matches == false || err != nil {
			continue
		}

		if matches == true {
			return productId
		}
	}

	return 0
}
