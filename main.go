package main

import (
	"fmt"
	"log"
	"os"

	dayTwo "jamesob/aoc/2022/02"
)

func main() {

	file, err := os.Open("02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Day One
	// elves := dayOne.GetElves(file)
	// fmt.Println("Day One Part One: " + dayOne.PartOne(elves))
	// fmt.Println("Day One Part Two: " + dayOne.PartTwo(elves))

	// Day Two
	// fmt.Println("Day Two Part One: " + dayTwo.PartOne(file))
	fmt.Println("Day Two Part Two: " + dayTwo.PartTwo(file))
}
