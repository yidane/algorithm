package problems

import (
	"fmt"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		args int
		want string
	}{
		//{1, "A"},
		//{28, "AB"},
		{52, "AZ"},
		//{701, "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := convertToTitle(tt.args); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Println('A')
}
