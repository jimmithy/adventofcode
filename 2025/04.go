package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("04.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []int{})

		for i := 0; i < len(line); i++ {
			hit := 0
			if line[i] == '@' {
				hit = 1
			}

			grid[row] = append(grid[row], hit)
		}

		row += 1
	}

	println("Part One:", dayFourPartOne(grid))
	println("Part Two:", dayFourPartTwo(grid))
}

func dayFourPartOne(grid [][]int) int {
	total := 0
	totalRows := len(grid)

	for rowIndex, row := range grid {
		rowLength := len(row)

		for colIndex := 0; colIndex < rowLength; colIndex++ {
			// Only search around toilet roles
			if row[colIndex] != 1 {
				continue
			}

			// finally confirm the count is fewer than four
			if hasFourAdjacent(grid, rowIndex, colIndex, totalRows, rowLength) {
				total += 1
			}
		}
	}

	return total
}

// 138 - @@.@..@@.@.@@@@@@@..@.@@@@..@@@.@.@@@@@@@@..@@..@.@@@.@@@@...@@.@....@@@@@@@.@..@@@@@@.@@@@@.@@.@@@@@@@@@@.@.@@@@.@@@@@@..@@@.@@.@.@@.@@@..
// 139 - @.@.@@.@@@@@@@@@.@@@@.@..@@@.@..@@@@.@@.@@@@@..@@@@..@...@@@@@@@@......@.@@@@@@@@...@.@@@@.@..@@@.@@@@@@.@@@..@@@@@@@@@@@@.@@.@@@@.@@.@@@@.

func dayFourPartTwo(grid [][]int) int {
	total := 0
	totalRows := len(grid)
	continueLooping := true
	mutableGrid := grid

	// Loop until we run out of toilet rolls to remove
	for continueLooping {
		hasChanges := false

		for rowIndex, _ := range grid {
			rowLength := len(grid[rowIndex])

			for colIndex := 0; colIndex < rowLength; colIndex++ {
				// Only search around toilet roles
				if grid[rowIndex][colIndex] != 1 {
					continue
				}

				// finally confirm the count is fewer than four
				if hasFourAdjacent(grid, rowIndex, colIndex, totalRows, rowLength) {
					total += 1
					mutableGrid[rowIndex][colIndex] = 0
					hasChanges = true
				}
			}
		}

		continueLooping = hasChanges
		grid = mutableGrid
	}

	return total
}

func hasFourAdjacent(grid [][]int, rowIndex int, colIndex int, totalRows int, rowLength int) bool {
	count := 0
	hasPreviousRow := rowIndex > 0
	hasLeft := colIndex > 0
	hasRight := colIndex < rowLength-1
	hasNextRow := rowIndex < totalRows-1

	// Search previous row
	if hasPreviousRow {
		previousRow := grid[rowIndex-1]
		if hasLeft {
			topLeft := previousRow[colIndex-1]
			if topLeft == 1 {
				count += 1
			}
		}

		above := previousRow[colIndex]
		if above == 1 {
			count += 1
		}

		if hasRight {
			topRight := previousRow[colIndex+1]

			if topRight == 1 {
				count += 1
			}
		}
	}

	// Search left and right positions
	if hasLeft {
		left := grid[rowIndex][colIndex-1]
		if left == 1 {
			count += 1
		}
	}

	if hasRight {
		right := grid[rowIndex][colIndex+1]
		if right == 1 {
			count += 1
		}
	}

	// search next row
	if hasNextRow {
		nextRow := grid[rowIndex+1]

		if hasLeft {
			bottomLeft := nextRow[colIndex-1]
			if bottomLeft == 1 {
				count += 1
			}
		}

		below := nextRow[colIndex]
		if below == 1 {
			count += 1
		}

		if hasRight {
			bottomRight := nextRow[colIndex+1]
			if bottomRight == 1 {
				count += 1
			}
		}
	}

	return count < 4
}
