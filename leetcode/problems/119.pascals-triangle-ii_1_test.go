package problems

import (
	"fmt"
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, tt := range tests {
		pascals := generate(30)

		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			got := getRow(tt)
			if !SameArray(got, pascals[tt]) {
				t.Fatal()
			}
		})
	}
}
