package problems

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		args int
		want string
	}{
		//{1, "I"},
		//{3, "III"},
		//{4, "IV"},
		//{9, "IX"},
		{1000, "M"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := intToRoman(tt.args); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
