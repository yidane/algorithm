package problems

import "testing"

func Test_multiply(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{num1: "0", num2: "0", want: "0"},
		{num1: "2", num2: "3", want: "6"},
		{num1: "408", num2: "5", want: "2040"},
		{num1: "123", num2: "456", want: "56088"},
		{num1: "9123", num2: "0", want: "0"},
		{num1: "498828660196", num2: "840477629533", want: "419254329864656431168468"},
	}
	for _, tt := range tests {
		t.Run(tt.num1, func(t *testing.T) {
			if got := multiply(tt.num1, tt.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
