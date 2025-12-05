package main

import (
	"bufio"
	"log"
	"os"
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
}
