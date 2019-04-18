package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{edges: [][]int{{1, 2}, {1, 3}, {2, 3}}, want: []int{2, 3}},
		{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, want: []int{1, 4}},
		{edges: [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}, want: []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.edges), func(t *testing.T) {
			if got := findRedundantConnection(tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
