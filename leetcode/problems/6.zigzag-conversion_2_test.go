package problems

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		args    string
		numRows int
		want    string
	}{
		{"AB", 1, "AB"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := convert(tt.args, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
