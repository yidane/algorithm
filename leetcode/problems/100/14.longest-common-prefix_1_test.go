package problems

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		//{args: []string{}, want: ""},
		//{args: []string{"", "", ""}, want: ""},
		//{args: []string{"c", "c"}, want: "c"},
		{args: []string{"c", "c", "c"}, want: "c"},
		//{args: []string{"yidane", "yinsiwen", "yichen"}, want: "yi"},
		//{args: []string{"flower", "flow", "flight"}, want: "fl"},
		//{args: []string{"dog", "racecar", "car"}, want: ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := longestCommonPrefix(tt.args); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
