package others

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_judgePoint24(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 1, 1, 4}, false},
		{[]int{1, 2, 3, 4}, true},
		{[]int{4, 1, 8, 7}, true},
		//{[]int{1, 5, 5, 5}, true}, //（5-1/5）×5=24/5*8=24
	}

	Convey("计算24点", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(judgePoint24(tt.nums), ShouldEqual, tt.want)
			})
		}
	})
}

func Test_judgePoint2(t *testing.T) {
	tests := []struct {
		r      int
		s      int
		target int
		want   bool
	}{
		{1, 2, 3, true},
		{1, 2, 1, true},
		{2, 2, 4, true},
		{1, 2, 4, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.r, tt.s, tt.target), func(t *testing.T) {
			exps := judgePoint2(newFraction(tt.r), newFraction(tt.s), newFraction(tt.target))
			if len(exps) > 0 != tt.want {
				t.Errorf("judgePoint2() = %v, want %v", !tt.want, tt.want)
			}

			if len(exps) > 0 {
				for _, exp := range exps {
					fmt.Println(fmt.Sprintf("%s=%d", exp, tt.target))
				}
			}
		})
	}
}

func Test_judgePoint3(t *testing.T) {
	tests := []struct {
		r      int
		s      int
		t      int
		target int
		want   bool
	}{
		{1, 2, 3, 4, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.r, tt.s, tt.t, tt.target), func(t *testing.T) {
			exps := judgePoint3(newFraction(tt.r), newFraction(tt.s), newFraction(tt.t), newFraction(tt.target))
			if len(exps) > 0 != tt.want {
				t.Errorf("judgePoint3() = %v, want %v", !tt.want, tt.want)
			}

			if len(exps) > 0 {
				for _, exp := range exps {
					fmt.Println(fmt.Sprintf("%s=%d", exp, tt.target))
				}
			}
		})
	}
}
