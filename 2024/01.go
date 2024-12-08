package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	left := []int{}
	right := []int{}

	// Read the file and split the numbers into two lists
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		leftNumber, leftErr := strconv.Atoi(line[0])

		if leftErr == nil {
			left = append(left, leftNumber)
		}

		rightNumber, rightErr := strconv.Atoi(line[1])

		if rightErr == nil {
			right = append(right, rightNumber)
		}
	}

	// Sort the two lists
	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			total = total + (left[i] - right[i])
		} else {
			total = total + (right[i] - left[i])
		}
	}

	fmt.Println(total)
}

func partTwo(scanner *bufio.Scanner) {
	total := 0
	left := []int{}
	right := []int{}

	// Read the file and split the numbers into two lists
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		leftNumber, leftErr := strconv.Atoi(line[0])

		if leftErr == nil {
			left = append(left, leftNumber)
		}

		rightNumber, rightErr := strconv.Atoi(line[1])

		if rightErr == nil {
			right = append(right, rightNumber)
		}
	}

	counts := make(map[int]int)
	for _, num := range right {
		counts[num]++
	}

	for i := 0; i < len(left); i++ {
		total += left[i] * counts[left[i]]
	}

	fmt.Println(total)
}
