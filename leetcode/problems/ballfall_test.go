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
		for i, tt := range tests {
			fmt.Println(i)
			Convey(strconv.Itoa(i), func() {
				h, tot := ballfall(tt.h, tt.t)
				So(h, ShouldBeGreaterThan, 0)
				So(tot, ShouldBeGreaterThan, 0)
			})

			//fmt.Println(i, tt)
		}
	})
}
