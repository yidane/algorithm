package problems

import "testing"

func Test_findTarget(t *testing.T) {
	tests := []struct {
		arr  []interface{}
		k    int
		want bool
	}{
		{arr: []interface{}{1}, k: 2, want: false},
		{arr: []interface{}{2, 1, 3}, k: 3, want: true},
		{arr: []interface{}{5, 3, 6, 2, 4, nil, 7}, k: 9, want: true},
		{arr: []interface{}{5, 3, 6, 2, 4, nil, 7}, k: 28, want: false},
		{arr: []interface{}{0, -1, 2, -3, nil, nil, 4}, k: -4, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := NewTree(tt.arr)

			if got := findTarget(root, tt.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
