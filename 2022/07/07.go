package seven

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(input *os.File) string {
	// Read Input
	scanner := bufio.NewScanner(input)

	// depth := make([]string, 0)
	files := make(map[string]int)
	currentFolder := "/"

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			// Was a command
			cmd := strings.Split(line, " ")

			if cmd[1] == "cd" {
				if cmd[2] == ".." {
					// Move back
					lastFolder := strings.LastIndex(currentFolder, "/")
					if lastFolder != -1 {
						fmt.Println("Moving " + currentFolder + " to " + currentFolder[:lastFolder])
						currentFolder = currentFolder[:lastFolder]
					}
				} else {
					if currentFolder != "/" {
						currentFolder = currentFolder + "/" + cmd[2]
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") == false && len(line) > 0 {
			// Find files, ignore directories
			file := strings.Split(line, " ")
			size, _ := strconv.Atoi(file[0])

			files[currentFolder] = files[currentFolder] + size
		}
	}

	// Find all folders that
	sum := 0

	for _, v := range files {
		if v > 100000 {
			sum += v
		}
	}

	return strconv.Itoa(sum)
}
