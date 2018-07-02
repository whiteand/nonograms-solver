package board

import (
	"fmt"
	"strings"
)

type Cell int

const (
	White Cell = iota
	Black
	Cross
)

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
func (c Cell) Times(n int) (res Row) {
	if n <= 0 {
		return Row{}
	}

	for i := 0; i < n; i++ {
		res = append(res, c)
	}
	return res
}

type Row []Cell

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

func (r Row) PlaceBlackFrom(ind int, count int) (previousSubRow Row) {
	previousSubRow = r.GetSubRow(ind-1, ind+count+1)

	for i := 0; i < count; i++ {
		r[ind+i] = Black
	}
	if ind > 0 {
		r[ind-1] = Cross
	}
	if ind+count < len(r) {
		r[ind+count] = Cross
	}
	return previousSubRow
}
func (r Row) GetSubRow(start, end int) (res Row) {
	res = make([]Cell, 0, end-start+1)
	for i := start; i < end; i++ {
		if i < 0 || i >= len(r) {
			res = append(res, White)
		} else {
			res = append(res, r[i])
		}
	}
	return res
}
func (r Row) InsertSubRow(subrow Row, ind int) {
	if ind < 0 {
		ind = 0
		subrow = subrow[-ind:]
	}
	for i := 0; i < len(subrow); i++ {
		if ind+i >= len(r) {
			return
		}
		r[ind+i] = subrow[i]
	}
}
func (r Row) CanPlaceBlackFrom(ind int, count int) (possiblePlaceBlack bool) {
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
func (r Row) Copy() Row {
	res := make([]Cell, len(r))
	copy(res, r)
	return res
}
func Concat(rows ...Row) (res Row) {
	for _, row := range rows {
		res = append(res, row...)
	}
	return res
}

func (r Row) FillPossible(blackSeries ...int) (res Row) {
	res = make([]Cell, len(r))
	copy(res, r)
	const maxPointLimit = 100000
	maxPoint := 0
	overflow := false
	var counter func(Row, []int, []int, int)
	counter = func(row Row, series []int, blackCount []int, lastPosition int) {
		if len(series) == 0 {
			if !EqualSeries(row.Series(), blackSeries) {
				return
			}
			//fmt.Println(maxPoint, "\t", row)
			for i, v := range row {
				if v == Black {
					blackCount[i]++
				}
			}
			maxPoint++
			return
		}
		// Find all positions for one
		for i := lastPosition; i < len(row); i++ {
			if !row.CanPlaceBlackFrom(i, series[0]) {
				continue
			}
			tryRow := row.Copy()
			tryRow.PlaceBlackFrom(i, series[0])
			counter(tryRow, series[1:], blackCount, i+series[0])
			// if maxPoint > maxPointLimit {
			// 	overflow = true

			// 	break
			// }
		}
	}
	blackCountPosition := make([]int, len(r))
	counter(res, blackSeries, blackCountPosition, -1)
	if overflow {
		fmt.Println("Overflowed!")
		return r
	}
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

type Board struct {
	Rows   []Row
	Height int
	Width  int
}

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
func (b Board) Copy() (res Board) {
	for i := range b.Rows {
		res.Rows = append(res.Rows, b.Rows[i].Copy())
	}
	res.Width = b.Width
	res.Height = b.Height
	return res
}
func (b Board) Println() {
	for _, row := range b.Rows {
		fmt.Println(row)
	}
}
func (b Board) Column(ind int) (res Row) {
	res = make([]Cell, b.Height)
	for i := 0; i < b.Height; i++ {
		res[i] = b.Rows[i][ind]
	}
	return res
}
func (b Board) InsertColumn(ind int, column []Cell) {
	for i := 0; i < b.Height; i++ {
		b.Rows[i][ind] = column[i]
	}
}
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
