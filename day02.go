package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ID int

func (id ID) Valid() bool {
	str := strconv.Itoa(int(id))
	// Only ids with an even number of digits can be repeated
	if len(str)%2 == 0 {
		// Is first half equal to last half
		first := str[0 : len(str)/2]
		last := str[len(str)/2:]

		if first == last {
			return false
		}
	}
	return true
}

func (id ID) HasRepeatingPattern() bool {
	str := strconv.Itoa(int(id))
	if len(str) < 2 {
		return false
	}

	for pLen := 1; pLen <= len(str)/2; pLen++ {
		if len(str)%pLen != 0 {
			// The string length must be divisible by the pattern length for it to be able to be repeating
			continue
		}

		repeatedPattern := strings.Repeat(str[0:pLen], len(str)/pLen)
		if str == repeatedPattern {
			return true
		}
	}

	return false
}

type Range struct {
	Start ID
	End   ID
	l     *DebugLogger
}

func (r *Range) Expand() []ID {
	ids := make([]ID, 0, r.End-r.Start+1)
	for i := r.Start; i <= r.End; i++ {
		ids = append(ids, ID(i))
	}
	return ids
}

func (r *Range) GetInvalidIDs() []ID {
	invalidIDs := []ID{}
	for _, id := range r.Expand() {
		if !id.Valid() {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs
}

func (r *Range) GetRepeatingIDs() []ID {
	ids := []ID{}
	for _, id := range r.Expand() {
		if id.HasRepeatingPattern() {
			ids = append(ids, id)
		}
	}

	return ids
}

func ParseRanges(lines []string, l *DebugLogger) ([]Range, error) {
	ranges := []Range{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for rangeString := range strings.SplitSeq(line, ",") {
			parts := strings.Split(strings.TrimSpace(rangeString), "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf(`invalid range format "%s"`, rangeString)
			}

			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf(`error parsing range "%s", part "%s": %w`, rangeString, parts[0], err)
			}

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf(`error parsing range "%s", part "%s": %w`, rangeString, parts[1], err)
			}

			ranges = append(ranges, Range{Start: ID(start), End: ID(end), l: l})
		}
	}
	return ranges, nil
}

func SolveDay2Part1(lines []string, l *DebugLogger) (int, error) {

	ranges, err := ParseRanges(lines, l)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ranges: %w", err)
	}

	l.Debugf("n_ranges = %d", len(ranges))

	invalidIdsSum := 0
	for _, r := range ranges {
		ids := r.GetInvalidIDs()
		l.Debugf(`range=%d-%d, n_invalid=%d`, r.Start, r.End, len(ids))
		for _, id := range ids {
			invalidIdsSum += int(id)
		}
	}

	return invalidIdsSum, nil
}

func SolveDay2Part2(lines []string, l *DebugLogger) (int, error) {
	ranges, err := ParseRanges(lines, l)
	if err != nil {
		return 0, fmt.Errorf("error parsing ranges: %w", err)
	}

	invalidSum := 0
	for _, r := range ranges {
		ids := r.GetRepeatingIDs()
		for _, id := range ids {
			invalidSum += int(id)
		}
	}

	return invalidSum, nil
}
