package problems

import "testing"

func Test_maxProduct152(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct152(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct152() = %v, want %v", got, tt.want)
			}
		})
	}
}
