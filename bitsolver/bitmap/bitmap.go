package bitmap

import (
	"fmt"
	"strings"
)

type Bitmap uint64

const Capacity int = 64

func NewBitmap() *Bitmap {
	res := Bitmap(0)
	return &res
}
func cap(b Bitmap) int {
	return Capacity
}
func outOfRangeError(ind int) error {
	return fmt.Errorf("Index is out of range: %d", ind)
}
func (b *Bitmap) Set(ind int) *Bitmap {
	if ind < 0 || ind >= Capacity {
		panic(outOfRangeError(ind))
	}
	var prev uint64 = uint64(*b)
	*b = Bitmap(prev | (1 << uint(ind)))
	return b
}
func (b *Bitmap) Remove(ind int) *Bitmap {
	var prev uint64 = uint64(*b)
	*b = Bitmap(prev &^ (1 << uint(ind)))
	return b
}
func (b Bitmap) Has(ind int) bool {
	if ind < 0 || ind >= Capacity {
		panic(outOfRangeError(ind))
	}

	return uint64(b)&uint64(1<<uint(ind)) > 0
}
func (b Bitmap) String() string {
	str := (strings.Repeat("0", Capacity) + fmt.Sprintf("%b", uint64(b)))
	str = str[len(str)-64:]
	var res strings.Builder
	res.WriteString(str[:4])
	str = str[4:]
	sep := " "
	for str != "" {
		res.WriteString(sep)
		if sep == " " {
			sep = " | "
		} else {
			sep = " "
		}
		res.WriteString(str[:4])
		str = str[4:]
	}
	return res.String()
}
func (b Bitmap) Indexes() (res []int) {
	for i := 0; i < Capacity; i++ {
		if b.Has(i) {
			res = append(res, i)
		}
	}
	return res
}

func (b Bitmap) And(other Bitmap) Bitmap {
	return Bitmap(uint64(b) & uint64(other))
}
func (b Bitmap) Or(other Bitmap) Bitmap {
	return Bitmap(uint64(b) | uint64(other))
}

func (b Bitmap) Xor(other Bitmap) Bitmap {
	return Bitmap(uint64(b) ^ uint64(other))
}
func (b Bitmap) AndNot(other Bitmap) Bitmap {
	return Bitmap(uint64(b) &^ uint64(other))
}
func (b Bitmap) Not() Bitmap {
	return Bitmap(^uint64(b))
}
