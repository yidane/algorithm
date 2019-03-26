package problems

import (
	"fmt"
	"testing"
)

func Test_convertBST(t *testing.T) {
	tests := []struct {
		args []interface{}
		want []interface{}
	}{
		{args: []interface{}{5, 2, 13}, want: []interface{}{18, 20, 13}},
		{args: []interface{}{2, 0, 3, -4, 1}, want: []interface{}{5, 6, 3, 2, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			tree := NewTree(tt.args)
			got := convertBST(tree)
			want := NewTree(tt.want)

			if !got.Equal(want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
