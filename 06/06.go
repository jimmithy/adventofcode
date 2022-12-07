package six

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FindMarkers(input *os.File, markerSize int) string {

	// Read Input
	scanner := bufio.NewScanner(input)
	content := ""

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			content += line
		}
	}

	// Find Markers
	for i := 0; i < len(content)+markerSize; i++ {
		subset := content[i : i+markerSize]

		if isUnique(subset) {
			return strconv.Itoa(i + len(subset))
		}
	}

	return strconv.Itoa(-1)
}

func isUnique(subset string) bool {
	for _, v := range subset {
		if strings.Count(subset, string(v)) > 1 {
			return false
		}
	}

	return true
}
