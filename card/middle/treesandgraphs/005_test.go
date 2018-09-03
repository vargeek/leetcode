package treesandgraphs

import "testing"

func Test_kthSmallest(t *testing.T) {
	tree1 := newTreeNodeFromInts(3, 1, 4, -1, 2)
	k1 := 1
	want1 := 1

	tree2 := newTreeNodeFromInts(5, 3, 6, 2, 4, -1, -1, 1)
	k2 := 3
	want2 := 3

	tree3 := newTreeNodeFromInts(1)
	k3 := 1
	want3 := 1

	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{tree1, k1}, want1},
		{"2", args{tree2, k2}, want2},
		{"3", args{tree3, k3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
