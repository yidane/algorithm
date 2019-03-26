package problems

import (
	"fmt"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	tests := []struct {
		args []int
	}{
		{[]int{1, 3, 2, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			wiggleSort(tt.args)

			//bigger := false
			//
			//fmt.Println(tt.args)
			//for i := 0; i < len(tt.args)-1; i++ {
			//	if tt.args[i] > tt.args[i+1] != bigger {
			//		t.Fatalf("%v should %v", tt.args[i], tt.args[i+1])
			//	}
			//
			//	bigger = !bigger
			//}
		})
	}
}
