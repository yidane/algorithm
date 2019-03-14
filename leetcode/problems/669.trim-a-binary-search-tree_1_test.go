package problems

import (
	"testing"
)

func Test_trimBST(t *testing.T) {
	tests := []struct {
		root func() *TreeNode
		L    int
		R    int
		want func() *TreeNode
	}{
		{root: func() *TreeNode {
			return NewTreeNode(1).AddLeft(0).AddRight(2)
		}, L: 1, R: 2, want: func() *TreeNode {
			return NewTreeNode(1).AddRight(2)
		}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := trimBST(tt.root(), tt.L, tt.R); !got.Equal(tt.want()) {
				t.Fail()
			}
		})

		t.Run("", func(t *testing.T) {
			if got := trimBST1(tt.root(), tt.L, tt.R); !got.Equal(tt.want()) {
				t.Fail()
			}
		})
	}
}
