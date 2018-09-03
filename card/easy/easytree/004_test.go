package easytree

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	// [3,9,20,null,null,15,7]
	tree1 := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	want1 := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{tree1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
