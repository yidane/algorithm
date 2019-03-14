package problems

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{}, target: 5}, want: 0},
		{args: args{nums: []int{5}, target: 5}, want: 0},
		{args: args{nums: []int{4}, target: 5}, want: 1},
		{args: args{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{args: args{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{args: args{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
		{args: args{nums: []int{1, 3, 5, 6}, target: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert(%v,%v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			}
		})
	}
}