package boardfromsite

import (
	"fmt"
	"testing"
)

func EqualInts(a, b []int) error {
	if a == nil && b == nil {
		return nil
	}
	if a == nil || b == nil {
		return fmt.Errorf("%v != %v", a, b)
	}
	if l1, l2 := len(a), len(b); l1 != l2 {
		return fmt.Errorf("Different length of (%d)%v and (%d)%v", a, l1, b, l2)
	}
	for i, v := range a {
		if b[i] != v {
			return fmt.Errorf("%v != %v", a, b)
		}
	}
	return nil
}

func EqualBlocks(a, b [][]int) (err error) {
	if a == nil && b == nil {
		return nil
	}
	if a == nil || b == nil {
		return fmt.Errorf("One of blocks is nil")
	}
	if len(a) != len(b) {
		return fmt.Errorf("Different length of blocks: %d != %d", len(a), len(b))
	}
	for i, v := range a {
		err := EqualInts(b[i], v)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestParse(t *testing.T) {
	outRowBlocks := [][]int{
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
	outColBlocks := [][]int{
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
	fileName := "test_input.html"
	rowBlocks, colBlocks, err := ParseFromFile(fileName)
	if err != nil {
		t.Errorf("Error getted: %v", err)
	}

	if err := EqualBlocks(rowBlocks, outRowBlocks); err != nil {
		t.Error(err)
	}
	if err := EqualBlocks(colBlocks, outColBlocks); err != nil {
		t.Error(err)
	}
}
