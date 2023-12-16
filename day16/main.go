package main

import (
	"fmt"
)

func main() {
	// Part One
	pacMan := PacMan{
		Position:      Point{Row: 0, Col: -1},
		NextDirection: "right",
	}

	res := worker(&pacMan)
	fmt.Println("Part One:", res)

	// Part Two
	pacMans := partTwoPacMans()
	max := 0

	for _, p := range pacMans {
		if r := worker(p); r > max {
			max = r
		}
	}
	fmt.Println("Part Two:", max)
}

func worker(p *PacMan) int {
	grid := CreateGrid()
	grid.Workers = append(grid.Workers, p)

	for {
		valid := grid.Next()
		if !valid {
			break
		}
	}
	return collectEnergized(grid)
}

func collectEnergized(g *Grid) int {
	count := 0
	for _, p := range g.Grid {
		if p.Energized {
			count++
		}
	}
	return count
}

// Generate a list of PacMans for each starting point on the grid
func partTwoPacMans() []*PacMan {
	grid := CreateGrid()
	rows := grid.Boundary.Row
	cols := grid.Boundary.Col

	pacMans := make([]*PacMan, 0)

	for r := 0; r < rows; r++ {
		pacManRight := PacMan{
			Position:      Point{Row: r, Col: -1},
			NextDirection: "right",
		}
		pacManLeft := PacMan{
			Position:      Point{Row: r, Col: cols + 1},
			NextDirection: "left",
		}
		pacMans = append(pacMans, &pacManRight, &pacManLeft)
	}

	for c := 0; c < cols; c++ {
		pacManUp := PacMan{
			Position:      Point{Row: -1, Col: c},
			NextDirection: "down",
		}
		pacManDown := PacMan{
			Position:      Point{Row: rows + 1, Col: c},
			NextDirection: "up",
		}
		pacMans = append(pacMans, &pacManUp, &pacManDown)
	}

	return pacMans
}
