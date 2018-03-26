package problems

import (
	"fmt"
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

	check := func(actual interface{}, expected ...interface{}) string {
		arr1 := actual.([][]int)
		arr2 := expected[0].([][]int)

		if len(arr1) != len(arr2) {
			return "1"
		}

		arr := make(map[int]int)
		for _, arrv := range arr1 {
			for _, v := range arrv {
				if a, ok := arr[v]; ok {
					arr[v] = a + 1
				} else {
					arr[v] = 1
				}
			}
		}

		for _, arrv := range arr2 {
			for _, v := range arrv {
				if a, ok := arr[v]; ok {
					arr[v] = a - 1
				} else {
					return fmt.Sprintf("Expected Result:%v,Result:%v", arr1, arr2)
				}
			}
		}

		for _, v := range arr {
			if v != 0 {
				return fmt.Sprintf("Expected Result:%v,Result:%v", arr1, arr2)
			}
		}

		return ""
	}

	Convey("4Sum_Test", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(fourSum(tt.args.nums, tt.args.target), check, tt.want)
			})
		}
	})
}
