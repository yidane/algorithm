package problems

import (
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		root func() *TreeNode
		want []float64
	}{
		{root: func() *TreeNode {
			root := NewTreeNode(3).AddLeft(9).AddRight(20)
			root.Right.AddLeft(15).AddRight(7)

			return root
		}, want: []float64{3, 14.5, 11}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := averageOfLevels(tt.root()); !SameFloat64Array(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
