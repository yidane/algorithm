package others

import "math"

type bigData struct {
	ints []int64
}

func NewBigData(size int) *bigData {
	n := size / 64
	return &bigData{
		ints: make([]int64, n+1),
	}
}

func (bd *bigData) set(v int) {
	n := v / 64
	m := v % 64

	i := bd.ints[n]

	if m == 63 {
		i = math.MinInt64
	} else {
		i = i | (1 << uint64(m))
	}

	bd.ints[n] = i
}

func (bd *bigData) find() int {
	for i := 0; i < len(bd.ints); i++ {
		bi := bd.ints[i]
		if bi == math.MinInt64 {
			continue
		}

		var j uint32
		for ; j < 64; j++ {
			if bi>>j == 0 {
				return i*64 + int(j)
			}
		}
	}

	return -1
}

func findOneInBigData(size int, c chan int) int {
	bd := NewBigData(size)

	for v := range c {
		bd.set(v)
	}

	return bd.find()
}
