package main

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(file *os.File) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func MustReadLines(file *os.File) []string {
	lines, err := ReadLines(file)
	if err != nil {
		log.Default().Fatalf("could not read file \"%s\": %s\n", file.Name(), err)
	}
	return lines
}

func ReadFile(fp string) ([]string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadLines(f)
}

func ReadInput() ([]string, error) {
	inputFile := "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ReadLines(file)
}

func MustReadInput() []string {
	lines, err := ReadInput()
	if err != nil {
		log.Default().Fatalf("could not read input: %s\n", err)
	}

	return lines
}

type Solution interface {
	Name() string
	Init(f string, l *log.Logger) error
	Solve() ([]string, error)
}
