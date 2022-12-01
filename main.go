package main

import (
	"fmt"
	"jamesob/aoc/2022/code"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Day One
	elves := code.GetElves(file)
	fmt.Println("Day One Part One: " + code.PartOne(elves))
	fmt.Println("Day One Part Two: " + code.PartTwo(elves))
}
