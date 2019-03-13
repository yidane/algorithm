package problems

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s    string
		k    int
		want string
	}
	tests := []args{
		{s: "abcdefg", k: 2, want: "bacdfeg"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseStr(tt.s, tt.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
