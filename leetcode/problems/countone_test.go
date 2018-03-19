package problems

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_CountOne(t *testing.T) {

	tests := []struct {
		arg  int
		want int
	}{
		{arg: 1, want: 1},
		{arg: 2, want: 1},
		{arg: 11, want: 3},
		{arg: 19, want: 11},
		{arg: 100, want: 20},
	}

	Convey("", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(CountOne1(tt.arg), ShouldEqual, tt.want)
			})
		}
	})
}

func Test_Fn(t *testing.T) {
	for i := 0; i < 200000; i++ {
		if i == Fn(i) {
			fmt.Println(i)
		}
	}
}
