package problems

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_dayInYear(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{year: 2017, month: 1, day: 1}, want: 1},
		{args: args{year: 2017, month: 2, day: 20}, want: 51},
		{args: args{year: 2017, month: 3, day: 9}, want: 68},
		{args: args{year: 2016, month: 3, day: 9}, want: 69},
		{args: args{year: 2016, month: 2, day: 28}, want: 59},
		{args: args{year: 2016, month: 12, day: 31}, want: 366},
	}

	Convey("problems/dayInYear_test.go", t, func() {
		for _, tt := range tests {
			Convey(fmt.Sprintf("%v-%v-%v", tt.args.year, tt.args.month, tt.args.day), func() {
				So(dayInYear(tt.args.year, tt.args.month, tt.args.day), ShouldEqual, tt.want)
			})
		}

		Convey("不合法的日期格式", func() {
			So(func() { dayInYear(2016, 6, 31) }, ShouldPanic)
			So(func() { dayInYear(2016, 2, 30) }, ShouldPanic)
			So(func() { dayInYear(2016, 1, 40) }, ShouldPanic)
			So(func() { dayInYear(2016, -2, 20) }, ShouldPanic)
			So(func() { dayInYear(2016, 14, 40) }, ShouldPanic)
		})
	})
}
