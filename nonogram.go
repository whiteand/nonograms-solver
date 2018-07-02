package main

import (
	"fmt"
	"log"
	"nonogram/board"
	"nonogram/boardfromsite"
	"nonogram/solver"
	"time"
)

func main() {

	rowBoxes, colBoxes, err := boardfromsite.ParseFromFile(`test_input.html`)
	if err != nil {
		log.Fatal(err)
	}

	solver.Logging = false
	resBoard, err := solver.Solve(rowBoxes, colBoxes)
	if err != nil {
		log.Fatal(err)
	}
	resBoard.Println()
	time.Sleep(time.Second)

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
