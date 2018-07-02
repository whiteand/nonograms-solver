package solver

import (
	"fmt"
	"nonogram/board"
)

var Logging = true

// A function that takes input horizontal and vertical sets
// of numbers, which are the initial data of the problem
//
// Returns the solution of the problem, or at least an
// intermediate solution
func Solve(rowBlocks, colBlocks [][]int) (result *board.Board, err error) {
	// If there are no numbers, then there is no solution
	if len(rowBlocks) == 0 || len(colBlocks) == 0 {
		return nil, fmt.Errorf("Wrong input data (%v, %v)", rowBlocks, colBlocks)
	}

	// Initializing an empty field
	b := board.NewBoard(len(rowBlocks), len(colBlocks))

	// Solution Iterations
	for i := 0; i < 100; i++ {
		// We are trying to find a more complete solution on
		// the basis of current data about the field and the
		// necessary conditions by certain input data of the
		// problem - numbers
		b = b.FillPossible(rowBlocks, colBlocks)
		if Logging {
			fmt.Println()
			b.Println()
		}

		// If there are no empty (white) cells left on the
		// field, then - a solution is found.
		hasWhite := b.HasWhite()
		if !hasWhite {
			break
		}
	}

	// return the solution
	return &b, nil

}
