package problems

import (
	"fmt"
	"testing"
)

func Test_canVisitAllRooms(t *testing.T) {
	tests := []struct {
		rooms [][]int
		want  bool
	}{
		//{rooms: [][]int{{1}, {2}, {3}, {}}, want: true},
		//{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, want: false},
		{rooms: [][]int{{4}, {3}, {}, {2, 5, 7}, {1}, {}, {8, 9}, {}, {}, {6}}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.rooms), func(t *testing.T) {
			if got := canVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
