package sortingandsearching

import (
	"reflect"
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"2", args{[]int{1}, 1}, []int{1}},
		{"3", args{[]int{3, 2, 2, 1, 1, 1}, 2}, []int{1, 2}},
		{"4", args{[]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3}, []int{-1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}

}
