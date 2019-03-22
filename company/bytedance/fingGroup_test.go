package bytedance

import (
	"fmt"
	"testing"
)

func Test_findGroup(t *testing.T) {
	tests := []struct {
		args         [][]int
		wantCount    int
		wantMaxGroup int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			gotCount, gotMaxGroup := findGroup(tt.args)
			if gotCount != tt.wantCount {
				t.Errorf("findGroup() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if gotMaxGroup != tt.wantMaxGroup {
				t.Errorf("findGroup() gotMaxGroup = %v, want %v", gotMaxGroup, tt.wantMaxGroup)
			}
		})
	}
}
