package problems

import "testing"

func Test_tree2str(t *testing.T) {
	tests := []struct {
		args []interface{}
		want string
	}{
		//{args: []interface{}{}, want: ""},
		{args: []interface{}{1, 2, 3, 4}, want: "1(2(4))(3)"},
		{args: []interface{}{1, 2, 3, nil, 4}, want: "1(2()(4))(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			tree := NewTree(tt.args)

			if got := tree2str(tree); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
