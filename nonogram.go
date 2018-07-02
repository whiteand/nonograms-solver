package main

import (
	"fmt"
	"log"
	"nonogram/board"
	"nonogram/solver"
	"time"
)

// Start of nonogram solver program
func main() {

	// Horizontal sets of numbers
	rowBoxes := [][]int{
		{4},
		{6},
		{4},
		{3, 1, 2},
		{5, 4},
		{5, 5},
		{1, 9, 1},
		{1, 11},
		{6, 1},
		{5},
		{4},
		{3},
		{1},
		{1},
		{4},
	}

	// Vertical sets of numbers
	colBoxes := [][]int{
		{1, 1},
		{2, 5},
		{6},
		{8},
		{3, 6},
		{3, 6, 1},
		{6, 1},
		{9},
		{6, 1},
		{4},
		{4},
		{5},
		{3, 2},
		{4},
	}

	// Dissable printing of progress
	solver.Logging = false

	// Request solution for this vertical and horizontal lines
	resBoard, err := solver.Solve(rowBoxes, colBoxes)

	// If error occurs, than close program - we cannot find solution
	if err != nil {
		log.Fatal(err)
	}

	// Show solution
	resBoard.Println()
	time.Sleep(time.Second)

	// Output the solution in text form
	for _, row := range resBoard.Rows {
		for j, cell := range row {
			if cell == board.Black && (j == 0 || row[j-1] != board.Black) {
				fmt.Printf("%d->", j+1)
			}
			if cell == board.Black && (j >= len(row)-1 || row[j+1] != board.Black) {
				fmt.Printf("%d\t", j+1)
			}

		}
		fmt.Println()
	}

}
