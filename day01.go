package main

import (
	"fmt"
	"log"
)

type Day01 struct {
	input []string
	log   *log.Logger
}

func (Day01) Name() string { return "Day01" }

func (d Day01) Init(f string, l *log.Logger) error {
	d.log = l
	d.input = make([]string, 0)
	lines, err := ReadFile(f)
	if err != nil {
		return err
	}
	d.input = lines

	return nil
}

func (d Day01) Solve() ([]string, error) {
	lineSum := 0
	d.log.Printf("Input length: %d", len(d.input))
	for _, line := range d.input {
		prefix := line[0]
		num := line[0:]
		d.log.Printf("Line: %s, Prefix: %s, Num: %s", line, prefix, num)
	}

	return []string{fmt.Sprintf("%d", lineSum)}, nil
}
