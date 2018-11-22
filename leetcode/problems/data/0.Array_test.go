package data

import "testing"

func TestSameArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{arr1: nil, arr2: []int{1, 2}}, want: false},
		{name: "2", args: args{arr1: nil, arr2: nil}, want: true},
		{name: "3", args: args{arr1: []int{1, 2}, arr2: []int{1, 2}}, want: true},
		{name: "4", args: args{arr1: []int{1, 2}, arr2: []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameArray(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("testCase %v: SameArray() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
