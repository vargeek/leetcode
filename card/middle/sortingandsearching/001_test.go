package sortingandsearching

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 0, 2, 1, 1, 0}}, []int{0, 0, 1, 1, 2, 2}},
		{"2", args{[]int{}}, []int{}},
		{"3", args{[]int{0}}, []int{0}},
		{"4", args{[]int{1}}, []int{1}},
		{"5", args{[]int{2}}, []int{2}},
		{"6", args{[]int{0, 1}}, []int{0, 1}},
		{"7", args{[]int{1, 0}}, []int{0, 1}},
		{"8", args{[]int{0, 2}}, []int{0, 2}},
		{"9", args{[]int{2, 0}}, []int{0, 2}},
		{"10", args{[]int{1, 2}}, []int{1, 2}},
		{"11", args{[]int{2, 1}}, []int{1, 2}},
		{"12", args{[]int{0, 1, 2}}, []int{0, 1, 2}},
		{"13", args{[]int{2, 1, 0}}, []int{0, 1, 2}},
		{"14", args{[]int{0, 0, 0}}, []int{0, 0, 0}},
		{"15", args{[]int{1, 1, 1}}, []int{1, 1, 1}},
		{"16", args{[]int{2, 2, 2}}, []int{2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
