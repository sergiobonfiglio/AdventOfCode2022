package utils

import (
	"bufio"
	"os"
)

func ReadLines(inputPath string) []string {
	readFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}
