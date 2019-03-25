package problems

import (
	"fmt"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		args []int
	}{
		{args: []int{4, 2, 5, 7}},
		{args: []int{3, 1, 4, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			var sort = func(f func([]int) []int, name string) {
				arr := make([]int, len(tt.args))
				copy(arr, tt.args)

				got := f(tt.args)

				for i := 0; i < len(got); i++ {
					if i%2 != got[i]%2 {
						t.Errorf("%s() = %v, want %v", name, got, i)
					}
				}
			}

			sort(sortArrayByParityIIChan, "sortArrayByParityIIChan")
			sort(sortArrayByParityII, "sortArrayByParityII")
			sort(sortArrayByParityII1, "sortArrayByParityII1")
		})
	}
}
