package problems

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s     []byte
		wants []byte
	}
	tests := []args{
		{s: []byte{}, wants: []byte{}},
		{s: []byte{'1'}, wants: []byte{'1'}},
		{s: []byte{'1', '2'}, wants: []byte{'2', '1'}},
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, wants: []byte{'o', 'l', 'l', 'e', 'h'}},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, wants: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			reverseString(tt.s)

			for i := 0; i < len(tt.s); i++ {
				if tt.s[i] != tt.wants[i] {
					t.Fatalf("%v should be equal %v", tt.s[i], tt.wants[i])
				}
			}
		})
	}
}
