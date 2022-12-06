package main

import (
	"fmt"
	"log"
	"os"

	code "jamesob/aoc/2022/05"
)

func main() {

	file, err := os.Open("05/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println(code.PartTwo(file))
}
