package treesandgraphs

import "testing"

func Test_maxPathSum(t *testing.T) {
	input1 := newTreeNodeFromInts(1, 2, 3)
	want1 := 6

	input2 := newTreeNodeFromInts(-10, 9, 20, null, null, 15, 7)
	want2 := 42

	input3 := newTreeNodeFromInts(5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1)
	want3 := 48

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{input1}, want1},
		{"2", args{input2}, want2},
		{"3", args{input3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
