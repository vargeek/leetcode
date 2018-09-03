package backtracking

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums1 := []int{1, 2, 3}
	want1 := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{1},
		[]int{2, 3},
		[]int{2},
		[]int{3},
		[]int{},
	}
	nums2 := []int{}
	want2 := [][]int{
		[]int{},
	}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{nums1}, want1},
		{"2", args{nums2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
