package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("01.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile("[0-9]")
		matches := re.FindAllString(line, -1)

		first := matches[0]
		last := matches[len(matches)-1]

		fmt.Println(first, last)

		matchesNumeric, err := strconv.Atoi(first + "" + last)

		if err == nil {
			total = total + matchesNumeric
		}
	}

	fmt.Println(total)
}

func partTwo(scanner *bufio.Scanner) {
	total := 0

	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	backwardsReplacer := strings.NewReplacer(
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	)

	for scanner.Scan() {
		raw := scanner.Text()
		line := replacer.Replace(raw)

		re := regexp.MustCompile("[0-9]")
		matches := re.FindAllString(line, -1)

		first := matches[0]

		backwards := Reverse(raw)
		backwardsLine := backwardsReplacer.Replace(backwards)
		backwardsMatches := re.FindAllString(backwardsLine, -1)

		last := backwardsMatches[0]

		matchesNumeric, err := strconv.Atoi(first + "" + last)

		if err == nil {
			total = total + matchesNumeric
		} else {
			fmt.Println("Error: " + err.Error())
		}
	}

	fmt.Println(total)
}

func Reverse(s string) string {
	var reverse string

	for i := len(s) - 1; i >= 0; i-- {
		reverse = reverse + string(s[i])
	}

	return reverse
}
