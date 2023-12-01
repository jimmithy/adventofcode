package main

import (
	"fmt"
	"log"
	"os"

	code "jamesob/aoc/2022/07"
)

func main() {

	file, err := os.Open("07/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println(code.PartOne(file))
}
