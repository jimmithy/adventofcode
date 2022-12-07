package main

import (
	"fmt"
	"log"
	"os"

	code "jamesob/aoc/2022/06"
)

func main() {

	file, err := os.Open("06/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println(code.FindMarkers(file, 14))
}
