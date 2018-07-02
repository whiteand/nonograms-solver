package board

import (
	"fmt"
	"strings"
)

// Type determining the values of cells on the board
type Cell int

// Constants of Cell type
const (
	White Cell = iota // empty cell
	Black             // filled cell
	Cross             // Cell that can not be black
)

// Stringer interface implementation
func (c Cell) String() string {
	switch c {
	case White:
		return "■"
	case Black:
		return "□"
	case Cross:
		return " "
	}
	return "ERROR"
}

// Repeat returns a new string with the length of the
// cell filled with this value
// Examples:
// If n <= 0: res = Row{}
// Otherwise res = Row{c,c,c,c, ..., c},
// len(res) = n
func (c Cell) Repeat(n int) (res Row) {

	// If n <= 0
	if n <= 0 {
		return Row{}
	}

	for i := 0; i < n; i++ {
		res = append(res, c)
	}
	return res
}

// Type to describe a single row of a board
type Row []Cell

// Stringer interface implementation
func (r Row) String() string {
	if len(r) <= 0 {
		return ""
	}
	var res strings.Builder
	res.WriteString(r[0].String())
	for _, v := range r[1:] {
		res.WriteRune(' ')
		res.WriteString(v.String())
	}
	return res.String()
}

// Series returns a slice of lengths of sequences of black cells
// Examples:
// If r = Row{White, White, White, White...} then res = []int{}
// If r = Row{Black, Black, White, Black} then res = []int{2,1}
func (r Row) Series() (res []int) {
	current := 0
	for _, cell := range r {
		if cell != Black {
			if current != 0 {
				res = append(res, current)
			}
			current = 0
			continue
		}
		current++
	}
	if current != 0 {
		res = append(res, current)
	}
	return res
}

// EqualSeries checks if two int slices are equal
// Examples:
// EqualSeries(nil, nil) = true
// EqualSeries(nil, non_nill) = false
// EqualSeries(non_nill, nil) = false
// EqualSeries(a, b) = false, if len(a) != len(b)
// EqualSeries([]int{1,2,3}, []int{1,2,3}) = true
func EqualSeries(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	for i, aValue := range a {
		if aValue != b[i] {
			return false
		}
	}
	return true
}

// Fill in the line a block of black cells surrounded by
// crosses
// ind - first Black cell position
// count - length of inserted sequence of black cells
// Example:
// Row{White,White,White,White,White}.PlaceBlackBlock(2,1)
// => Row{Cross,Black,Cross,White,White}
func (r Row) PlaceBlackBlock(ind int, count int) {
	for i := 0; i < count; i++ {
		r[ind+i] = Black
	}
	if ind > 0 {
		r[ind-1] = Cross
	}
	if ind+count < len(r) {
		r[ind+count] = Cross
	}
}

// Checks whether it is possible to insert a block of black
// cells - without violating the current position
// ind - first Black cell position
// count - length of inserted sequence of black cells
func (r Row) CanPlaceBlackBlock(ind int, count int) (possiblePlaceBlack bool) {
	if ind+count > len(r) {
		return false
	}
	if ind < 0 {
		return false
	}
	for i := ind; i < ind+count; i++ {
		if r[i] == Cross {
			return false
		}
	}
	return !((ind > 0 && r[ind-1] == Black) || (ind+count < len(r) && r[ind+count] == Black))
}

// Copy return copy of row
func (r Row) Copy() Row {
	res := make([]Cell, len(r))
	copy(res, r)
	return res
}

// Bonds several sequences of cells into one sequence
// Example:
// Concat(Row{White}, Row{Black,Cross}, Row{White})
// returns Row{White, Black, Cross, White}
func Concat(rows ...Row) (res Row) {
	for _, row := range rows {
		res = append(res, row...)
	}
	return res
}

// Looks at all possible solutions for this row of board
// cells that can be generated from this string to match
// the Japanese crossword rules for input numbers.
// If some cell takes a black color for all solutions,
// its color is set to black.
// If a cell takes a white color for all the solutions
// - we denote it by a cross
func (r Row) FillPossible(blackSeries ...int) (res Row) {
	res = make([]Cell, len(r))
	copy(res, r)
	maxPoint := 0
	var counter func(Row, []int, []int, int)
	counter = func(row Row, series []int, blackCount []int, lastPosition int) {
		if len(series) == 0 {
			if !EqualSeries(row.Series(), blackSeries) {
				return
			}
			for i, v := range row {
				if v == Black {
					blackCount[i]++
				}
			}
			maxPoint++
			return
		}
		for i := lastPosition; i < len(row); i++ {
			if !row.CanPlaceBlackBlock(i, series[0]) {
				continue
			}
			tryRow := row.Copy()
			tryRow.PlaceBlackBlock(i, series[0])
			counter(tryRow, series[1:], blackCount, i+series[0])
		}
	}
	blackCountPosition := make([]int, len(r))
	counter(res, blackSeries, blackCountPosition, -1)

	for i := 0; i < len(res); i++ {
		if r[i] == Black {
			blackCountPosition[i]++
		}
	}

	for i := 0; i < len(res); i++ {
		if blackCountPosition[i] == maxPoint {
			res[i] = Black
		}
		if blackCountPosition[i] == 0 {
			res[i] = Cross
		}
	}

	return res
}

// Type representing the board in the program
type Board struct {
	Rows   []Row
	Height int
	Width  int
}

// NewBoard - Construtor of empty board with such size
func NewBoard(height, width int) (r Board) {
	res := make([]Row, height)
	for i := range res {
		res[i] = make([]Cell, width)
		for j := range res[i] {
			res[i][j] = White
		}
	}
	return Board{res, height, width}
}

// Copy returns copy of board
func (b Board) Copy() (res Board) {
	res.Rows = make([]Row, len(b.Rows))
	for i := range res.Rows {
		res.Rows[i] = b.Rows[i].Copy()
	}
	res.Width = b.Width
	res.Height = b.Height
	return res
}

// Println outputs board
func (b Board) Println() {
	for _, row := range b.Rows {
		fmt.Println(row)
	}
}

// Column returns the sequence of cells that are on the
// specified column with a given index
func (b Board) Column(ind int) (res Row) {
	res = make([]Cell, b.Height)
	for i := 0; i < b.Height; i++ {
		res[i] = b.Rows[i][ind]
	}
	return res
}

// InsertColumn inserts a sequence of cells, instead of a
// column with a given index
func (b Board) InsertColumn(ind int, column []Cell) {
	for i := 0; i < b.Height; i++ {
		b.Rows[i][ind] = column[i]
	}
}

// FillPossible tries to find more information about the
// solution of the problem on the basis of already found
// information about the solution and the specified
// requirements for the solution (slices of numbers)
func (b Board) FillPossible(rowBoxes, colBoxes [][]int) (res Board) {
	res = b.Copy()
	// byRow
	for rowInd, row := range res.Rows {
		res.Rows[rowInd] = row.FillPossible(rowBoxes[rowInd]...)
	}

	// byColumns
	for j := 0; j < res.Width; j++ {
		col := res.Column(j)
		col = col.FillPossible(colBoxes[j]...)
		res.InsertColumn(j, col)
	}

	return res
}

// If the field contains white (unfilled cells),
// it returns true
func (b Board) HasWhite() bool {
	for _, row := range b.Rows {
		for _, cell := range row {
			if cell == White {
				return true
			}
		}
	}
	return false
}
