package easytree

import "testing"

func Test_isValidBST(t *testing.T) {
	tree1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	tree2 := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	tree3 := &TreeNode{1, &TreeNode{1, nil, nil}, nil}
	tree4 := &TreeNode{10, &TreeNode{5, nil, nil}, &TreeNode{15, &TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}}}
	tree5 := &TreeNode{34, &TreeNode{-6, &TreeNode{-21, nil, nil}, nil}, nil}

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
		{"2", args{tree3}, false},
		{"3", args{tree4}, false},
		{"4", args{tree5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
