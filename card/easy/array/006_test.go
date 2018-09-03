package easy

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{"2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
		{"3", args{[]int{61, 24, 20, 58, 95, 53, 17, 32, 45, 85, 70, 20, 83, 62, 35, 89, 5, 95, 12, 86, 58, 77, 30, 64, 46, 13, 5, 92, 67, 40, 20, 38, 31, 18, 89, 85, 7, 30, 67, 34, 62, 35, 47, 98, 3, 41, 53, 26, 66, 40, 54, 44, 57, 46, 70, 60, 4, 63, 82, 42, 65, 59, 17, 98, 29, 72, 1, 96, 82, 66, 98, 6, 92, 31, 43, 81, 88, 60, 10, 55, 66, 82, 0, 79, 11, 81}, []int{5, 25, 4, 39, 57, 49, 93, 79, 7, 8, 49, 89, 2, 7, 73, 88, 45, 15, 34, 92, 84, 38, 85, 34, 16, 6, 99, 0, 2, 36, 68, 52, 73, 50, 77, 44, 61, 48}}, []int{5, 4, 57, 79, 7, 89, 88, 45, 34, 92, 38, 85, 6, 0, 77, 44, 61}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newBSTTreeFromInts(t *testing.T) {
	tree := &bstTree{node: &bstNode{val: 20, count: 1, left: &bstNode{val: 10, count: 2, left: &bstNode{val: 5, count: 1, left: nil, right: nil}, right: nil}, right: nil}}
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *bstTree
	}{
		{"1", args{[]int{20, 10, 5, 10}}, tree},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBSTTreeFromInts(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBSTTreeFromInts() = %v, want %v", got, tt.want)
			}
		})
	}

	if !tree.contains(20) {
		t.Error("tree should contain 20")
	}

	if !tree.contains(10) {
		t.Error("tree should contain 10")
	}
	if !tree.contains(5) {
		t.Error("tree should contain 5")
	}

	if tree.contains(30) {
		t.Error("tree should not contain 30")
	}

}
