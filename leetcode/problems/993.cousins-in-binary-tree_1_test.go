package problems

import "testing"

func Test_isCousins(t *testing.T) {
	tests := []struct {
		arr  []interface{}
		x    int
		y    int
		want bool
	}{
		{arr: []interface{}{1, 2, 3, nil, 4, nil, 5}, x: 5, y: 4, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := NewTree(tt.arr)
			if got := isCousins(root, tt.x, tt.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
