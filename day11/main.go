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
// Calculate total Manhattan distance between all points x expandFactor (x2)

// Part Two:
// Calculate total Manhattan distance between all points x expandFactor (x1000000)

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

	gallaxyList, expandedSpace := ReadFile()

	// Part One
	expandFactor := 2
	dist := calculateDistance(gallaxyList, expandedSpace, expandFactor)
	fmt.Println("Part One:", dist)

	// Part Two
	expandFactor = 1000000
	dist = calculateDistance(gallaxyList, expandedSpace, expandFactor)
	fmt.Println("Part Two:", dist)
}

// Calculating the total distance between all points
func calculateDistance(list List, e ExpandedSpace, expandFactor int) int {
	dist := 0

	// For each combination of points (x,y) in space
	for i := 0; i < len(list); i++ {
		p1 := list[i]
		for j := i + 1; j < len(list); j++ {
			p2 := list[j]

			// Calculate Manhattan distance between two points
			m := manhattanDistance(p1, p2)

			// Calculate the number of expanded space borders the points are passing through
			b := bordersCount(p1, p2, e)

			// Calculate the point distance (pd) with regard to the expand factor
			pd := m + (b * (expandFactor - 1))

			// Add to total distance
			dist += pd
		}
	}
	return dist
}

// Calculate how many borders of expanded space the points are passing through 
func bordersCount(p1, p2 Point, e ExpandedSpace) int {
	r := 0
	c := 0

	rAbs := int(math.Abs(float64(p1.x - p2.x)))         // absoule row-distance between two points
	rMin := int(math.Min(float64(p1.x), float64(p2.x))) // minimum row index between two points

	for i := rMin; i <= rAbs+rMin; i++ {
		// if row index exists in expanded space, increment row border count
		if e.rows[i] {
			r++
		}
	}

	cAbs := int(math.Abs(float64(p1.y - p2.y)))
	cMin := int(math.Min(float64(p1.y), float64(p2.y)))

	for j := cMin; j <= cAbs+cMin; j++ {
		if e.cols[j] {
			c++
		}
	}
	return r + c
}

func manhattanDistance(p1 Point, p2 Point) int {
	asFloat := math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y))
	return int(asFloat)
}
