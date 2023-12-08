package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// WriteLines writes the lines to the given file.
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func GetCharMatrix(path string) [][]string {
	var matrix [][]string
	lines := ReadLines(path)
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func ConvertToNumbers(arr []string) []int {
	var numbers []int
	for _, elem := range arr {
		if elem == " " || elem == "" {
			continue
		}
		n, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func ConvertToNumber(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
