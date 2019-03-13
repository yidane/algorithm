package problems

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		args   []int
		target int
		want   []int
	}{
		{args: []int{2, 7, 11, 15, 12, 31, 22, 12, 311, 123, 441, 41223, 3, 43, 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 100000, 999}, target: 100999, want: []int{25, 26}},
		{args: []int{3, 3}, target: 6, want: []int{0, 1}},
		{args: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
	}

	//自定义检测方法
	check := func(actual interface{}, expected ...interface{}) string {
		a := actual.([]int)
		b := expected[0].([]int)
		if a == nil && b == nil {
			return ""
		}
		if (a != nil && b != nil) &&
			((a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])) {
			return ""
		}
		if a == nil || len(b) != 2 {
			return "Result is Nil or len is not equal 2."
		}
		if b == nil || len(b) != 2 {
			return "Expected Result is Nil or len is not equal 2."
		}

		return fmt.Sprintf("Expected Result:%v,Result:%v", a, b)
	}

	Convey("1.two-sum_1_test.go", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(twoSum0(tt.args, tt.target), check, tt.want)
				So(twoSum1(tt.args, tt.target), check, tt.want)
				So(twoSum2(tt.args, tt.target), check, tt.want)
			})
		}
	})
}
