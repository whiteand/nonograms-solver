package bitsolver

import (
	"fmt"
	"nonogram/boardfromsite"
	"testing"
)

func TestSolve(t *testing.T) {
	rows, cols, err := boardfromsite.ParseFromFile("test_input.html")
	if err != nil {
		t.Errorf("Read file error: %v", err)
	}
	resBoard, err := Solve(rows, cols)
	if err != nil {
		t.Errorf("solve problem: %v", err)
	}
	fmt.Println(resBoard.Width, resBoard.Height)
	resBoard.Println()

}
