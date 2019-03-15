package problems

import (
	"fmt"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	root := NewTreeNode(379).AddLeft(826)

	result := increasingBST(root)

	fmt.Println(result)
}
