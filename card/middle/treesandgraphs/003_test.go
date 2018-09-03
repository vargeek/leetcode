package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	pre1 := []int{3, 9, 20, 15, 7}
	in1 := []int{9, 3, 15, 20, 7}
	want1 := newTreeNodeFromInts(3, 9, 20, -1, -1, 15, 7)
	pre2 := []int{}
	in2 := []int{}
	want2 := newTreeNodeFromInts()
	pre3 := []int{1, 2, 3, 4, 5}
	in3 := []int{5, 4, 3, 2, 1}
	want3 := newTreeNodeFromInts(1, 2, -1, 3, -1, 4, -1, 5)

	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{pre1, in1}, want1},
		{"2", args{pre2, in2}, want2},
		{"3", args{pre3, in3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
