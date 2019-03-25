package problems

import (
	"fmt"
	"testing"
)

func TestNewRecentCounter(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3001, 3002}, []int{1, 2, 3, 3}},
	}
	for _, tt := range tests {
		counter := NewRecentCounter()
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			if len(tt.input) != len(tt.output) {
				t.Fatal("input dot not match output")
			}

			for i := 0; i < len(tt.input); i++ {
				if got := counter.Ping(tt.input[i]); got != tt.output[i] {
					t.Fatalf("%d != %d", got, tt.output[i])
				}
			}
		})
	}
}
