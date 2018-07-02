package bitmap

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	if cap(Bitmap(32)) != 64 {
		t.Error("Capacity must be 64")
	}
}
func TestHas(t *testing.T) {
	tests := []struct {
		bitmap  Bitmap
		ind     int
		out     bool
		isPanic bool
	}{

		{Bitmap(0), -1, false, true},
		{Bitmap(0), 0, false, false},
		{Bitmap(0), 1, false, false},
		{Bitmap(0), 63, false, false},
		{Bitmap(0), 64, false, true},

		{Bitmap(32), -1, false, true},
		{Bitmap(32), 0, false, false},
		{Bitmap(32), 4, false, false},
		{Bitmap(32), 5, true, false},
		{Bitmap(32), 6, false, false},
		{Bitmap(32), 63, false, false},
		{Bitmap(32), 64, false, true},

		{Bitmap(uint64(1 << 63)), -1, false, true},
		{Bitmap(uint64(1 << 63)), 0, false, false},
		{Bitmap(uint64(1 << 63)), 62, false, false},
		{Bitmap(uint64(1 << 63)), 63, true, false},

		{Bitmap(uint64(1 << 63)), 64, false, true},
	}

	for _, test := range tests {
		func() {
			defer func() {
				x := recover()
				if x != nil && !test.isPanic {
					t.Errorf("Panic, but expected (%v).Has(%d) => %v", test.bitmap, test.ind, test.out)
					return
				}
				if x == nil && test.isPanic {
					t.Errorf("Has no panic, for (%v).Has(%d)", test.bitmap, test.ind)
				}

			}()
			if res := test.bitmap.Has(test.ind); res != test.out {
				t.Errorf("Has() error: (%v).Has(%d) => %v, but %v expected", test.bitmap.Indexes(), test.ind, res, test.out)
			}
		}()
	}
}

func TestSet(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 61, 62, 63},
	}
	for _, test := range tests {
		a := NewBitmap()
		for _, ind := range test {
			a.Set(ind)
			if !a.Has(ind) {
				t.Errorf("%v must have %d, but it didn't", a, ind)
			}
		}

	}

}
