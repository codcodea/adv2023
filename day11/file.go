package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	input = "input.txt"
	sample = "sample.txt"
)

func ReadFile() (List, ExpandedSpace) {
	s := make(Space, 0) // init empty space

	space := populateSpace(s)
	maps := getExpandedSpace(space)
	list := makeListOfPoints(space)
	return list, maps
}

func populateSpace(s Space) Space {
	file, err := os.Open(input)
	checkErr(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
	}
	return s
}

func getExpandedSpace(s Space) ExpandedSpace {
	e := ExpandedSpace{
		rows: make(map[int]bool),
		cols: make(map[int]bool),
	}

	// for each row
	for i, row := range s {
		hasGallaxy := strings.Contains(row, "#")

		// init rows map 
		if _, ok := e.rows[i]; !ok {
			e.rows[i] = false
		}

		// if the row has no galaxy, set to expand (true)
		if !hasGallaxy {
			e.rows[i] = true
		}

		// for each col
		for j, char := range row {

			// init cols map 
			if _, ok := e.cols[j]; !ok {
				e.cols[j] = true
			}

			// if col has a galaxy, set to not expand (false)
			if char == '#' {
				e.cols[j] = false
			}
		}
	}
	return e
}

// Create a list of galaxy points (x, y) in space
func makeListOfPoints(s Space) List {
	list := make(List, 0)

	for i, line := range s {
		for j, char := range line {
			if char == '#' {
				p := Point{x: i, y: j}
				list = append(list, p)
			}
		}
	}
	return list
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
