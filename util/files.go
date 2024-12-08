package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileAsGrid(fileName string) [][]string {
	var matrix [][]string
	lines := ReadFile(fileName)
	for _, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}
	return matrix
}

func ReadFile(fileName string) []string {
	inputFile, err := os.Getwd()
	//inputFile = inputFile + "/day1/" + fileName
	inputFile = inputFile + fileName
	CheckError(err)

	readFile, err := os.Open(inputFile)
	CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	return fileLines
}

func ReadFileAsString(fileName string) string {
	inputFile, err := os.Getwd()
	inputFile = inputFile + fileName
	CheckError(err)
	f, err := os.Open(inputFile)
	CheckError(err)
	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	return b.String()
}
