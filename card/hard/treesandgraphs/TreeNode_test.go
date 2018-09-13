package treesandgraphs

import (
	"reflect"
	"testing"
)

func Test_newTreeNodeFromInts(t *testing.T) {
	nums1 := []int{3, 9, 20, null, null, 15, 7}
	want1 := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	nums2 := []int{3}
	want2 := &TreeNode{3, nil, nil}
	nums3 := []int{3, null, 5}
	want3 := &TreeNode{3, nil, &TreeNode{5, nil, nil}}
	nums4 := []int{}
	var want4 *TreeNode

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{nums1}, want1},
		{"2", args{nums2}, want2},
		{"3", args{nums3}, want3},
		{"4", args{nums4}, want4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTreeNodeFromInts(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTreeNodeFromInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
