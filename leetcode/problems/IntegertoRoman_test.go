package problems

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{num: 2}, want: "II"},
		{args: args{num: 12}, want: "XII"},
		{args: args{num: 4}, want: "IV"},
		{args: args{num: 29}, want: "XXIX"},
		{args: args{num: 49}, want: "XLIX"},
		{args: args{num: 6}, want: "VI"},
		{args: args{num: 101}, want: "CI"},
		{args: args{num: 110}, want: "CX"},
		{args: args{num: 1001}, want: "MI"},
	}

	Convey("数字转换为罗马数字", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(intToRoman(tt.args.num), ShouldEqual, tt.want)
			})
		}
	})
}
