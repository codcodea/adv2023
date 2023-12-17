package main

import (
	"fmt"
	"math"
)

// Common:
// - Read file to create NxM matrix of "space" ( where "." = empty space, "#" = galaxy)
// - Record expanded space (empty rows or columns)
// - Create a list of galaxy points (x, y) in space

// Part One:
// Calculate total Manhattan distance between all points x expandBy (x2)

// Part Two:
// Calculate total Manhattan distance between all points x expandBy (x1000000)

type Space []string

type ExpandedSpace struct {
	rows map[int]bool // map[rowIndex] = true if row is empty
	cols map[int]bool // map[colIndex] = true if col is empty
}

type Point struct {
	x int
	y int
}

type List []Point

func main() {
	gList, eSpace := ReadFile()

	// Part One
	expandBy := 2
	dist := calculateDistance(gList, eSpace, expandBy)
	fmt.Println("Part One:", dist)

	// Part Two
	expandBy = 1000000
	dist = calculateDistance(gList, eSpace, expandBy)
	fmt.Println("Part Two:", dist)
}

func calculateDistance(list List, e ExpandedSpace, expandFactor int) int {
	dist := 0

	// For each combination of points (x,y) 
	for i := 0; i < len(list); i++ {
		p1 := list[i]

		for j := i + 1; j < len(list); j++ {
			p2 := list[j]

			// Calculate Manhattan distance 
			m := manhattanDistance(p1, p2)

			// Calculate how many borders of expanded space the points are passing through
			b := bordersCount(p1, p2, e)

			// Calculate the point distance (pd) 
			pd := m + (b * (expandFactor - 1))

			// Add to total distance
			dist += pd
		}
	}
	return dist
}

func bordersCount(p1, p2 Point, e ExpandedSpace) int {
	count := 0

	// Traverse X-axis
	xAbs := int(math.Abs(float64(p1.x - p2.x)))        
	xMin := int(math.Min(float64(p1.x), float64(p2.x))) 

	for i := xMin; i <= xAbs+xMin; i++ {
		// if index exists in expanded space map, increment count
		if e.rows[i] {
			count++
		}
	}

	// Traverse Y-axis
	yAbs := int(math.Abs(float64(p1.y - p2.y)))
	yMin := int(math.Min(float64(p1.y), float64(p2.y)))

	for j := yMin; j <= yAbs+yMin; j++ {
		if e.cols[j] {
			count++
		}
	}
	return count
}

func manhattanDistance(p1 Point, p2 Point) int {
	asFloat := math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y))
	return int(asFloat)
}
