package main

import (
	"fmt"
	"log"
	"os"

	code "jamesob/aoc/2022/04"
)

func main() {

	file, err := os.Open("04/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Day Four
	fmt.Println(code.PartTwo(file))
}
