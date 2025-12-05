package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IngredientRange struct {
	min int
	max int
}

func main() {
	file, err := os.Open("05.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	freshIngredients := []IngredientRange{}
	hasFreshIngredients := false

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// We've hit the break, so mark that we're now reading ingredients
			hasFreshIngredients = true
			continue
		} else if hasFreshIngredients {
			// Check if this ingredient is fresh
			ingredientID, ingredientErr := strconv.Atoi(line)

			if ingredientErr == nil {
				for i := 0; i < len(freshIngredients); i++ {
					if (ingredientID >= freshIngredients[i].min) && (ingredientID <= freshIngredients[i].max) {
						total += 1
						break
					}
				}
			}

		} else {
			// Add range of fresh ingredients to map
			ingredientRange := strings.Split(line, "-")
			min, minErr := strconv.Atoi(ingredientRange[0])
			max, maxErr := strconv.Atoi(ingredientRange[1])

			if (minErr == nil) && (maxErr == nil) {
				freshIngredients = append(freshIngredients, IngredientRange{min, max})
			}
		}

	}

	println("Part One:", total)

	// Count the number of ingredientIds that are fresh
	count := 0

	// Sort the ranges by min value
	slices.SortFunc(freshIngredients, func(a, b IngredientRange) int {
		return a.min - b.min
	})

	// Merge overlapping ranges
	mergedRanges := []IngredientRange{}

	for i := 0; i < len(freshIngredients); i++ {
		currentRange := freshIngredients[i]

		if len(mergedRanges) == 0 {
			// add the first range
			mergedRanges = append(mergedRanges, currentRange)
			continue
		}

		for j := 0; j < len(mergedRanges); j++ {
			if currentRange.max >= mergedRanges[j].min && currentRange.min <= mergedRanges[j].max {
				// Ranges overlap, so merge them
				mergedRanges[j].min = min(mergedRanges[j].min, currentRange.min)
				mergedRanges[j].max = max(mergedRanges[j].max, currentRange.max)
				break
			} else if j == len(mergedRanges)-1 {
				// No overlap found, so add the current range to mergedRanges
				mergedRanges = append(mergedRanges, currentRange)
			}
		}
	}

	// Count total number of fresh ingredient IDs
	for i := 0; i < len(mergedRanges); i++ {
		count += (mergedRanges[i].max - mergedRanges[i].min + 1)
	}

	println("Part Two:", count)
}
