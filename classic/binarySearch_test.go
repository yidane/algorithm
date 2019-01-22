package classic

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 21}
	e := 3

	i := BinarySearch(arr, e)
	if i == -1 {
		t.Fatal(i)
	}

	t.Log(i)
}
