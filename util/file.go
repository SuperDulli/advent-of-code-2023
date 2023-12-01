package util

import (
	"bufio"
	"log"
	"os"
	"path"
)

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

func GetLinesFromKnownFile(name string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return ReadLines(path.Join(cwd,name))
}

func GetExample() []string {
	return GetLinesFromKnownFile("example.txt")
}

func GetInput() []string {
	return GetLinesFromKnownFile("input.txt")
}
