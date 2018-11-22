package others

import "testing"

func Test_convertToBase7(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{name: "100", arg: 100, want: "202"},
		{name: "-7", arg: -7, want: "-10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.arg); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
