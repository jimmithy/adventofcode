package main

import (
	"fmt"
	"log"
	"os"

	code "jamesob/aoc/2022/03"
)

func main() {

	file, err := os.Open("03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Day Three
	fmt.Println("Day Three Part Two: " + code.PartTwo(file))
}
