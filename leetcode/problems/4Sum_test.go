package problems

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{args: args{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{args: args{nums: []int{-3, -1, 0, 2, 4, 5}, target: 0}, want: [][]int{{-3, -1, 0, 4}}},
	}
	Convey("4Sum_Test", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(fourSum(tt.args.nums, tt.args.target), ShouldEqual, tt.want)
			})
		}
	})
}
