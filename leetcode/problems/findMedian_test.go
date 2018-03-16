package problems

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		want float64
	}{
		{arr1: []int{1}, arr2: []int{2}, want: 1.5},
		{arr1: []int{1, 3, 4, 5, 6}, arr2: []int{4}, want: 4},
		{arr1: []int{1, 2, 3}, arr2: []int{1, 2, 3}, want: 2},
	}

	Convey("寻找中值", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(findMedianSortedArrays(tt.arr1, tt.arr2), ShouldEqual, tt.want)
			})
		}
	})
}
