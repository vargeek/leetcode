package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tree1 := newTreeNodeFromInts(3, 9, 20, -1, -1, 15, 7)
	want1 := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	}
	var tree2 *TreeNode
	want2 := [][]int{}
	tree3 := newTreeNodeFromInts(1, 2, -1, 3, -1, 4, -1)
	want3 := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{4},
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
		{"1", args{tree2}, want2},
		{"2", args{tree3}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}

	newTreeNodeFromInts(11, 2, 3)
}
