package solver

import (
	"fmt"
	"nonogram/board"
)

var Logging = true

func Solve(rowBlocks, colBlocks [][]int) (result *board.Board, err error) {
	if len(rowBlocks) == 0 || len(colBlocks) == 0 {
		return nil, fmt.Errorf("Wrong input data (%v, %v)", rowBlocks, colBlocks)
	}
	b := board.NewBoard(len(rowBlocks), len(colBlocks))
	for i := 0; i < 100; i++ {
		b = b.FillPossible(rowBlocks, colBlocks)
		if Logging {
			fmt.Println()
			b.Println()
		}
		hasWhite := b.HasWhite()
		if !hasWhite {
			break
		}
	}
	return &b, nil

}
