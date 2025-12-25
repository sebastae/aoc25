package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Battery int
type BatteryBank struct {
	Batteries []Battery
	l         *DebugLogger
}

type BatteryPosition struct {
	Index   int
	Battery Battery
}

func (bank BatteryBank) GetMaxUsingN(n int) int {
	out := 0

	maxPos := BatteryPosition{-1, -1}
	for p := n - 1; p >= 0; p-- {
		for i, b := range bank.Batteries[maxPos.Index+1 : len(bank.Batteries)-p] {
			if maxPos.Index < 0 || b > maxPos.Battery {
				maxPos = BatteryPosition{i, b}
			}
		}

		out += int(maxPos.Battery) * int(math.Pow10(p))
		maxPos.Battery = -1
	}

	return out
}

type day03 struct{}

var Day03 day03

func (day03) ParseBatteryBanks(lines []string, l *DebugLogger) ([]BatteryBank, error) {
	banks := []BatteryBank{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		bank := BatteryBank{Batteries: []Battery{}, l: l}
		for _, ch := range strings.TrimSpace(line) {
			digit, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, fmt.Errorf(`error parsing battery "%v": %w`, ch, err)
			}

			bank.Batteries = append(bank.Batteries, Battery(digit))
		}

		banks = append(banks, bank)
	}

	return banks, nil
}

func (day03) SolvePart1(lines []string, l *DebugLogger) (int, error) {

	banks, err := Day03.ParseBatteryBanks(lines, l)
	if err != nil {
		return 0, fmt.Errorf("could not parse banks: %w", err)
	}

	const NumBatteriesActivePerBank = 2

	battSum := 0
	for _, bank := range banks {
		battSum += bank.GetMaxUsingN(NumBatteriesActivePerBank)
	}

	return battSum, nil
}
