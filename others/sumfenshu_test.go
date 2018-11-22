package others

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_sumfenshu(t *testing.T) {
	tests := []struct {
		arg int
	}{
		{arg: 30},
		{arg: 20},
	}

	Convey("problems/sumfenshu_test.go", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				for j := 1; j <= tt.arg; j++ {
					So(sumfenshu(j), ShouldAlmostEqual, sumfenshu1(j))
				}
			})
		}
	})
}
