package problems

import (
	"fmt"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		args int
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := countAndSay(tt.args); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
