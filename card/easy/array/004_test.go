package easy

import (
	"reflect"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1,2,3,1", args{[]int{1, 2, 3, 1}}, true},
		{"1,2,3,4", args{[]int{1, 2, 3, 4}}, false},
		{"1,1,1,3,3,4,3,2,4,2", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3, 4}},
		{"2", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"3", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, []int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4}},
		{"4", args{[]int{1, 1, 2, 1, 1}}, []int{1, 1, 1, 1, 2}},
		{"5", args{[]int{1, 1, 2, 2, 2, 1, 1}}, []int{1, 1, 1, 1, 2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
			reflect.DeepEqual(tt.args.nums, tt.want)
		})
	}
}
