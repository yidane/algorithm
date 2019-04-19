package problems

import (
	"fmt"
	"testing"
)

func Test_compareVersion(t *testing.T) {
	tests := []struct {
		version1 string
		version2 string
		want     int
	}{
		{"0.1", "1.1", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"1.1", "1.10", -1},
		{"1.2", "1.10", -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.version1, tt.version2), func(t *testing.T) {
			if got := compareVersion(tt.version1, tt.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
