package problems

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_judgePoint24(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{nums: []int{1, 2, 3, 4}}, want: false},
	}

	Convey("计算24点", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(judgePoint24(tt.args.nums), ShouldEqual, tt.want)
			})
		}
	})
}
