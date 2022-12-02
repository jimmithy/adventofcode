package main

import (
	"fmt"
	"log"
	"os"

	dayOne "jamesob/aoc/2022/01"
)

func main() {

	file, err := os.Open("01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Day One
	elves := dayOne.GetElves(file)
	fmt.Println("Day One Part One: " + dayOne.PartOne(elves))
	fmt.Println("Day One Part Two: " + dayOne.PartTwo(elves))
}
