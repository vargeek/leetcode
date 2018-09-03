package easytree

import "testing"

func Test_isSymmetric(t *testing.T) {
	tree1 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	tree2 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{tree1}, true},
		{"1", args{tree2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
