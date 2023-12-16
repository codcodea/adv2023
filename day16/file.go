package main

import (
	"bufio"
	"os"
	"sync"
)

const (
	input  = "input.txt"
	sample = "sample.txt"
)

// Store the original grid in a global variable to avoid re-reading the file
var (
	originalGrid Grid
	once         sync.Once
)

// A Point is a coordinate (x,y) in the grid
type Point struct {
	Row int
	Col int
}

// PointState stores the current state of a point (x,y) in the grid
type PointState struct {
	Operation   string // (e.g. ".", "/", "\", "|", "-")
	Energized   bool
	ViditedFrom []Point
}

// A Grid is a map of Points with their current states
type Grid struct {
	Grid     map[Point]PointState
	Boundary Point     // Last point in the grid defines the boundary
	Workers  []*PacMan // Stores PacMan workers
}

// Grid.Next() moves all workers one step ahead
func (g *Grid) Next() bool {
	hasNext := false

	for _, worker := range g.Workers {
		if worker.Move(g, &g.Boundary) {
			hasNext = true
		}
	}
	return hasNext
}

// PacMan is a worker that moves through the grid
type PacMan struct {
	Position      Point
	NextDirection string
	PointState    PointState
	IsDead        bool
}

// PacMan.Move() moved the PacMan one step ahead
func (p *PacMan) Move(g *Grid, boundary *Point) bool {

	// return if PacMan is dead
	if p.IsDead {
		return false
	}

	// store old position
	oldPoint := p.Position

	// calculate next position
	switch p.NextDirection {
	case "up":
		p.Position.Row--
	case "down":
		p.Position.Row++
	case "left":
		p.Position.Col--
	case "right":
		p.Position.Col++
	}

	// return and kill PacMan if out of bounds
	if p.Position.Row < 0 || p.Position.Col < 0 || p.Position.Row > boundary.Row || p.Position.Col > boundary.Col {
		p.IsDead = true
		return false
	}

	// load next PointState from the grid
	next := g.Grid[p.Position]

	// if a point is visited from the same direction twice, PacMan will die
	for _, v := range next.ViditedFrom {
		if v == oldPoint {
			p.IsDead = true
			return false
		}
	}

	// update state 
	next.Energized = true
	next.ViditedFrom = append(next.ViditedFrom, oldPoint)
	p.PointState = next
	g.Grid[p.Position] = next

	// Execute the PointState operation
	switch p.PointState.Operation {
	case ".":
		// Do nothing
	case "/":
		switch p.NextDirection {
		case "up":
			p.NextDirection = "right"
		case "down":
			p.NextDirection = "left"
		case "left":
			p.NextDirection = "down"
		case "right":
			p.NextDirection = "up"
		}
	case "\\":
		switch p.NextDirection {
		case "up":
			p.NextDirection = "left"
		case "down":
			p.NextDirection = "right"
		case "left":
			p.NextDirection = "up"
		case "right":
			p.NextDirection = "down"
		}
	case "|":
		switch p.NextDirection {
		case "up", "down":
			// Do nothing
		case "left", "right":
			p.NextDirection = "up"

			// Spawn a new PacMan
			pacMan := PacMan{
				Position:      Point{Row: p.Position.Row, Col: p.Position.Col},
				NextDirection: "down",
				PointState:    next,
			}
			g.Workers = append(g.Workers, &pacMan)
		}
	case "-":
		switch p.NextDirection {
		case "up", "down":
			p.NextDirection = "left"

			// Spawn a new PacMan
			pacMan := PacMan{
				Position:      Point{Row: p.Position.Row, Col: p.Position.Col},
				NextDirection: "right",
				PointState:    next,
			}

			g.Workers = append(g.Workers, &pacMan)
		case "left", "right":
			// Do nothing
		}
	}
	return true
}

// CreateGrid reads the input file and creates a grid
func CreateGrid() *Grid {
	// Read the input file only once
	once.Do(func() {
		originalGrid = Grid{
			Grid: make(map[Point]PointState),
		}

		// Read the input file
		file, err := os.Open(input)
		checkErr(err)

		scanner := bufio.NewScanner(file)

		n := -1
		m := -1

		// Set values from file
		for scanner.Scan() {
			line := scanner.Text()
			n++
			for j, char := range line {

				if j > m {
					m = j
				}

				p := Point{Row: n, Col: j}
				s := PointState{Operation: string(char), Energized: false, ViditedFrom: make([]Point, 0)}
				originalGrid.Grid[p] = s
			}
		}

		originalGrid.Boundary = Point{Row: n, Col: m}
	})

	// Create a new grid with its original state
	g := &Grid{
		Grid: make(map[Point]PointState),
	}

	// Copy the values from the originalGrid to the new grid
	for k, v := range originalGrid.Grid {
		g.Grid[k] = v
	}

	g.Boundary = originalGrid.Boundary
	return g
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
