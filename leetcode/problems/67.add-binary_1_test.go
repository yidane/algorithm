package problems

import (
	"fmt"
	"testing"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		//{"1", "11", "100"},
		//{"1010", "1011", "10101"},
		{"1111", "1111", "11110"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
