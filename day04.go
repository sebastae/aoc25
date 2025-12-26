package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	Filled bool
}

type Grid struct {
	Height int
	Width  int
	Cells  []Cell

	l *DebugLogger
}

func (g Grid) InBounds(x, y int) bool {
	return !(x < 0 || x >= g.Width || y < 0 || y >= g.Height)
}

func (g Grid) GetCell(x, y int) (*Cell, error) {
	idx := x + y*g.Width
	if !g.InBounds(x, y) {
		return nil, fmt.Errorf("coordinate (%d, %d) out of bounds [%d, %d]", x, y, g.Width, g.Height)
	}

	cell := &g.Cells[idx]
	return cell, nil
}

func (g Grid) GetAdjacent(x, y int) ([]*Cell, error) {
	if !g.InBounds(x, y) {
		return nil, fmt.Errorf("coordinate (%d, %d) out of bounds [%d, %d]", x, y, g.Width, g.Height)
	}

	cells := []*Cell{}
	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			if !(xOff == 0 && yOff == 0) && g.InBounds(x+xOff, y+yOff) {
				cellPtr, err := g.GetCell(x+xOff, y+yOff)
				if err != nil {
					return nil, err
				}

				cells = append(cells, cellPtr)
			}
		}
	}

	return cells, nil
}

func (g Grid) GetAccessibleFilledCells() (int, error) {

	const MAX_FILLED_ADJACENT = 4

	accessibleCells := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell, err := g.GetCell(x, y)
			if err != nil {
				return 0, fmt.Errorf("error getting grid-cell: %w", err)
			}
			if cell.Filled {
				adjacent, err := g.GetAdjacent(x, y)
				if err != nil {
					return 0, fmt.Errorf("error getting adjacent cells: %w", err)
				}

				if len(Filter(adjacent, func(cell *Cell) bool { return cell.Filled })) < MAX_FILLED_ADJACENT {
					accessibleCells++
				}
			}
		}
	}

	return accessibleCells, nil
}

type day04 struct{}

var Day04 day04

func (day04) ParseGrid(lines []string, l *DebugLogger) (*Grid, error) {
	grid := Grid{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		grid.Height++
		for i, ch := range strings.TrimSpace(line) {
			switch ch {
			case '.':
				grid.Cells = append(grid.Cells, Cell{false})
			case '@':
				grid.Cells = append(grid.Cells, Cell{true})
			default:
				return nil, fmt.Errorf(`unknown grid symbol "%v"`, ch)
			}

			if grid.Width < i+1 {
				grid.Width = i + 1
			}
		}
	}

	return &grid, nil
}

func (day04) SolvePart1(lines []string, l *DebugLogger) (int, error) {

	grid, err := Day04.ParseGrid(lines, l)
	if err != nil {
		return 0, fmt.Errorf("error parsing grid: %w", err)
	}

	accessible, err := grid.GetAccessibleFilledCells()
	if err != nil {
		return 0, fmt.Errorf("error getting number of accessible cells: %w", err)
	}

	return accessible, nil
}
