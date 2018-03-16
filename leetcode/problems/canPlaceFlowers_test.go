package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0", args: args{flowerbed: []int{1}, n: 0}, want: true},
		{name: "1", args: args{flowerbed: []int{0}, n: 1}, want: true},
		{name: "2", args: args{flowerbed: []int{0, 0}, n: 1}, want: true},
		{name: "3", args: args{flowerbed: []int{0, 0}, n: 2}, want: true},
		{name: "4", args: args{flowerbed: []int{0, 0, 0}, n: 2}, want: true},
		{name: "5", args: args{flowerbed: []int{0, 0, 1}, n: 1}, want: true},
		{name: "6", args: args{flowerbed: []int{0, 1, 0}, n: 1}, want: true},
		{name: "7", args: args{flowerbed: []int{1, 0, 0}, n: 1}, want: true},
		{name: "8", args: args{flowerbed: []int{0, 1, 0, 1}, n: 1}, want: false},
		{name: "9", args: args{flowerbed: []int{0, 0, 1, 0, 0}, n: 2}, want: true},
	}

	/*
		So(canPlaceFlowers([]int{1}, 0), ShouldEqual, true)
		So(canPlaceFlowers([]int{0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0}, 2), ShouldNotEqual, true)
		So(canPlaceFlowers([]int{0, 0, 0}, 2), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 0, 1}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 1, 0}, 1), ShouldNotEqual, true)
		So(canPlaceFlowers([]int{1, 0, 0}, 1), ShouldEqual, true)
		So(canPlaceFlowers([]int{0, 1, 0, 1}, 1), ShouldEqual, false)
		So(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 2), ShouldEqual, true)
	*/

	Convey("测试花坛放置", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				So(canPlaceFlowers(tt.args.flowerbed, tt.args.n), ShouldEqual, tt.want)
			})
		}
	})
}
