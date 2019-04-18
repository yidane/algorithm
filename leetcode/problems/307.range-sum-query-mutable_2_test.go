package problems

import (
	"testing"
)

func TestNumArray(t *testing.T) {
	numArr := NumArrayConstructor([]int{1, 3, 5})
	t1 := numArr.SumRange(0, 2)
	if t1 != 9 {
		t.Fatal()
	}

	numArr.Update(1, 2)
	t2 := numArr.SumRange(0, 2)
	if t2 != 8 {
		t.Fatal()
	}
}
