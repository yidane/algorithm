package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{args: []int{5, 2, 6, 1}, want: []int{2, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := countSmaller(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
