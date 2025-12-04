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

	println("Part One:", dayFour(grid))

}

// 138 - @@.@..@@.@.@@@@@@@..@.@@@@..@@@.@.@@@@@@@@..@@..@.@@@.@@@@...@@.@....@@@@@@@.@..@@@@@@.@@@@@.@@.@@@@@@@@@@.@.@@@@.@@@@@@..@@@.@@.@.@@.@@@..
// 139 - @.@.@@.@@@@@@@@@.@@@@.@..@@@.@..@@@@.@@.@@@@@..@@@@..@...@@@@@@@@......@.@@@@@@@@...@.@@@@.@..@@@.@@@@@@.@@@..@@@@@@@@@@@@.@@.@@@@.@@.@@@@.

func dayFour(grid [][]int) int {
	total := 0
	totalRows := len(grid)

	for rowIndex, row := range grid {
		rowLength := len(row)

		for colIndex := 0; colIndex < rowLength; colIndex++ {
			// Only search around toilet roles
			if row[colIndex] != 1 {
				continue
			}

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
				left := row[colIndex-1]
				if left == 1 {
					count += 1
				}
			}

			if hasRight {
				right := row[colIndex+1]
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

			// finally confirm the count is fewer than four
			if count < 4 {
				total += 1
			}
		}
	}

	return total
}
