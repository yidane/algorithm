package problems

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		args uint32
		want int
	}{
		{1, 1},
		{00000000000000000000000000001011, 3},
		{00000000000000000000000010000000, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := hammingWeight(tt.args); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}

			if got := hammingWeight1(tt.args); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
