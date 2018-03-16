package problems

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{nums: []int{1, 1, 1, 0}, target: 100}, want: 3},
		{args: args{nums: []int{0, 0, 0, 0}, target: 100}, want: 0},
		{args: args{nums: []int{0, 1, 2}, target: 0}, want: 3},
	}

	Convey("3SumClosest.test", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(threeSumClosest(tt.args.nums, tt.args.target), ShouldEqual, tt.want)
			})
		}
	})
}
