package problems

import (
	"fmt"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		args []interface{}
		want [][]int
	}{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, [][]int{{15, 7}, {9, 20}, {3}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			got := levelOrderBottom(root)

			if len(got) != len(tt.want) {
				t.Fatal()
			}

			for i := 0; i < len(got); i++ {
				if !SameArray(got[i], tt.want[i]) {
					t.Fatal()
				}
			}
		})
	}
}
