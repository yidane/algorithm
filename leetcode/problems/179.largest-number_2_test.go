package problems

import (
	"fmt"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	tests := []struct {
		args []int
		want string
	}{
		{args: []int{0, 0}, want: "0"},
		{args: []int{10, 2}, want: "210"},
		{args: []int{3, 30, 34, 5, 9}, want: "9534330"},
		{args: []int{3, 43, 48, 94, 85, 33, 64, 32, 63, 66}, want: "9485666463484333332"},
		{args: []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}, want: "9609938824824769735703560743981399"},
		{args: []int{121, 12}, want: "12121"},
	}
	//9609938824782469735703560743981399
	//9609938824824769735703560743981399
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := largestNumber1(tt.args); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigger(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want bool
	}{
		{1, 2, false},
		{3, 3, false},
		{3, 31, true},
		{3, 34, false},
		{3, 331, true},
		{3, 30, true},
		{121, 12, false},
		{122, 12, true},
		{123, 12, true},
		{345, 34, true},
		{3435, 34, true},
		{22, 2, true},
		{23, 2, true},
		{21, 2, false},
		{8247, 824, false},
		{824, 8247, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.a, tt.b), func(t *testing.T) {
			if got := bigger(tt.a, tt.b); got != tt.want {
				t.Errorf("bigger() = %v, want %v", got, tt.want)
			}
		})
	}
}
