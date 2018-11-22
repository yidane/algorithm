package others

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'g'}, want: 'j'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'j'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'}, want: 'c'},
	}

	Convey("probelms/nextGreatestLetter_test.go", t, func() {
		for i, tt := range tests {
			Convey(strconv.Itoa(i), func() {
				So(nextGreatestLetter(tt.args.letters, tt.args.target), ShouldEqual, tt.want)
			})
		}
	})
}
