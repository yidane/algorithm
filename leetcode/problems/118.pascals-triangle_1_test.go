package problems

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			got := generate(tt)

			for _, v := range got {
				fmt.Println(v)
			}
		})
	}
}
