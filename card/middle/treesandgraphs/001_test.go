package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	want := []int{1, 3, 2}
	tree2 := &TreeNode{1, nil, nil}
	want2 := []int{1}
	tree3 := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}
	want3 := []int{3, 1, 5}
	var tree4 *TreeNode
	want4 := []int{}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{tree}, want},
		{"1", args{tree2}, want2},
		{"2", args{tree3}, want3},
		{"3", args{tree4}, want4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
