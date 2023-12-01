package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("01.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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
