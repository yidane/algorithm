package problems

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{args: "())", want: 1},
		{args: "(((", want: 3},
		{args: "()", want: 0},
		{args: "()))((", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
