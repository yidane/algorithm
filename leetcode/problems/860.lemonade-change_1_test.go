package problems

import (
	"fmt"
	"testing"
)

func Test_lemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{bills: []int{5, 5, 5, 10, 20}, want: true},
		{bills: []int{5, 5, 10}, want: true},
		{bills: []int{10, 10}, want: false},
		{bills: []int{5, 5, 10, 10, 20}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.bills), func(t *testing.T) {
			if got := lemonadeChange(tt.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
