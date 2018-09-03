package easytree

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	want1 := &TreeNode{0, &TreeNode{-3, &TreeNode{-10, nil, nil}, nil}, &TreeNode{9, &TreeNode{5, nil, nil}, nil}}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"0", args{[]int{-10, -3, 0, 5, 9}}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
