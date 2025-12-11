package main

import (
	"fmt"
	"log"
	"strconv"
)

type Dial struct {
	Positions       int
	CurrentPosition int
}

type Direction int

const (
	DirectionLeft  Direction = -1
	DirectionRight Direction = 1
)

func (dial *Dial) Move(dir Direction, distance int) int {
	modDist := distance % dial.Positions
	newPos := dial.CurrentPosition + modDist*int(dir)
	if newPos < 0 {
		newPos = dial.Positions + newPos
	}

	dial.CurrentPosition = newPos % dial.Positions
	return dial.CurrentPosition
}

func parseDay01Line(line string) (Direction, int, error) {
	if len(line) == 0 {
		return 0, 0, fmt.Errorf("line is empty")
	}

	dirChar := string(line[0])
	dir := DirectionLeft
	if dirChar == "R" {
		dir = DirectionRight
	} else if dirChar != "L" {
		return 0, 0, fmt.Errorf(`unknown direction "%s"`, dirChar)
	}

	if len(line) == 1 {
		return dir, 0, nil
	}

	dist, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse distance: %w", err)
	}

	return dir, dist, nil
}

func SolveDay1Part1(lines []string, l *log.Logger) (int, error) {
	dial := Dial{100, 50}
	numTimesZero := 0

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir, dist, err := parseDay01Line(line)
		if err != nil {
			return 0, fmt.Errorf(`error parsing line %d ("%s"): %w`, i+1, line, err)
		}

		if dial.Move(dir, dist) == 0 {
			numTimesZero++
		}

	}

	return numTimesZero, nil
}
