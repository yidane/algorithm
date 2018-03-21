package problems

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ballfall(t *testing.T) {

	tests := []struct {
		h float64
		t uint
	}{
		{h: 100, t: 10},
	}

	Convey("计算球落地弹起的高度", t, func() {
		Convey("leetcode/problems/ballfall_test.go", func() {})
		for i, tt := range tests {
			fmt.Println(i)
			Convey(strconv.Itoa(i), func() {
				h, tot := ballfall(tt.h, tt.t)
				h1, tot1 := ballfall1(tt.h, tt.t)
				So(h, ShouldEqual, h1)
				So(tot, ShouldEqual, tot1)
			})
		}
	})
}
