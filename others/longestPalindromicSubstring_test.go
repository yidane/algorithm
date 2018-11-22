package others

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{arg: "abab", want: "aba"},
		{arg: "aba", want: "aba"},
		{arg: "a", want: "a"},
		{arg: "ccc", want: "ccc"},
	}

	Convey("leetcode/problems/longestPalindromicSubstring_test", t, func() {
		Convey("寻找回环数", func() {

		})
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(longestPalindrome(tt.arg), ShouldEqual, tt.want)
			})
		}
	})
}
